package jobsuc

import (
	"context"
	"fmt"

	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
)

type getActiveJobsUseCase struct{}

func NewGetActiveJobsUseCase() *getActiveJobsUseCase {
	return &getActiveJobsUseCase{}
}

func (uc *getActiveJobsUseCase) GetActiveJobs(ctx context.Context) ([]entities.Job, error) {
	db := database.DB.Db

	jobs := []entities.Job{}

	if err := db.Where("status = ?", entities.ACTIVE).Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err.Error())
	}

	return jobs, nil
}
