package jobsuc

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/database"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

type getJobsUseCase struct{}

func NewGetJobsUseCase() *getJobsUseCase {
	return &getJobsUseCase{}
}

func (uc *getJobsUseCase) GetJobs(ctx context.Context, userClaims middlewares.Claims) ([]entities.Job, error) {
	db := database.DB.Db

	jobs := []entities.Job{}

	switch {
	case slices.Contains(userClaims.RealmAccess.Roles, string(enums.ROLE_ADMIN)):
		// No filtering needed.
		break
	case slices.Contains(userClaims.RealmAccess.Roles, string(enums.ROLE_TALENT)):
		// Talent: Return all active jobs
		db = db.Where("status = ?", entities.ACTIVE)
	case slices.Contains(userClaims.RealmAccess.Roles, string(enums.ROLE_CLIENT)):
		// Client: Return jobs for this client
		db = db.Where("client_keycloak_id = ?", userClaims.Email)
	default:
		// If role doesn't match any expected roles, return an error
		return nil, errors.New("unauthorized: invalid role")
	}

	if err := db.Find(&jobs).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err.Error())
	}

	return jobs, nil
}
