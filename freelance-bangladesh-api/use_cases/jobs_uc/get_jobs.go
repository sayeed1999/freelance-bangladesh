package jobsuc

import (
	"context"
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type getJobsUseCase struct{}

func NewGetJobsUseCase() *getJobsUseCase {
	return &getJobsUseCase{}
}

func (uc *getJobsUseCase) GetJobs(ctx context.Context) ([]entities.Job, error) {
	db := database.DB.Db

	jobs := []entities.Job{}

	if err := db.Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err.Error())
	}

	return jobs, nil
}
