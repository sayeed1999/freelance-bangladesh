package jobsuc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type CreateJobRequest struct {
	Title       string  `validate:"required,min=3,max=50"`
	Description string  `validate:"min=0,max=1000"`
	Budget      float32 `validate:"required"`
	Deadline    *time.Time
}

type CreateJobResponse struct {
	Job *entities.Job
}

// interface for the usecase
type createJobUseCase struct{}

func NewCreateJobUseCase() *createJobUseCase {
	return &createJobUseCase{}
}

func (uc *createJobUseCase) CreateJob(ctx context.Context, claims middlewares.Claims, request CreateJobRequest) (*CreateJobResponse, error) {
	db := database.DB.Db

	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var client entities.Client

	if err := db.First(&client, "Email = ?", claims.Email).Error; err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err.Error())
	}

	if !client.IsVerified {
		return nil, fmt.Errorf("failed to create job: client account is not verified")
	}

	var job = &entities.Job{
		ClientID:    client.ID,
		Title:       request.Title,
		Description: request.Description,
		Budget:      request.Budget,
		Deadline:    request.Deadline,
	}

	if err := db.Create(&job).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %s", err.Error())
	}

	var response = &CreateJobResponse{Job: job}
	return response, nil
}
