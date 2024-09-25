package jobsuc

import (
	"context"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
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

func (uc *createJobUseCase) CreateJob(ctx context.Context, request CreateJobRequest) (*CreateJobResponse, error) {

	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	var job = &entities.Job{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		Title:       request.Title,
		Description: request.Description,
		Budget:      request.Budget,
		Deadline:    request.Deadline,
	}

	// TODO: create job or return err

	var response = &CreateJobResponse{Job: job}
	return response, nil
}
