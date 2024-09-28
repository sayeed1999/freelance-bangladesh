package jobsuc

import (
	"context"

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

	db.Find(&jobs)

	return jobs, nil
}
