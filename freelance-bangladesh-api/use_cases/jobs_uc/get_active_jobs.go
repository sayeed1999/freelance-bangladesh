package jobsuc

import (
	"context"
	"fmt"
	"slices"

	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

type getActiveJobsUseCase struct{}

func NewGetActiveJobsUseCase() *getActiveJobsUseCase {
	return &getActiveJobsUseCase{}
}

func (uc *getActiveJobsUseCase) GetActiveJobs(ctx context.Context, userClaims middlewares.Claims) ([]entities.Job, error) {
	db := database.DB.Db

	jobs := []entities.Job{}

	if slices.Contains(userClaims.RealmAccess.Roles, string(enums.ROLE_TALENT)) {
		db = db.Where("status = ?", entities.ACTIVE)
	} else if slices.Contains(userClaims.RealmAccess.Roles, string(enums.ROLE_CLIENT)) {
		db = db.Where("client_keycloak_id = ?", userClaims.Email)
	}

	if err := db.Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err.Error())
	}

	return jobs, nil
}
