package jobsuc

import (
	"context"

	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type getJobsUseCase struct{}

func NewGetJobsUseCase() *getJobsUseCase {
	return &getJobsUseCase{}
}

func (uc *getJobsUseCase) GetJobs(ctx context.Context) []entities.Job {
	all := []entities.Job{}
	return all
}
