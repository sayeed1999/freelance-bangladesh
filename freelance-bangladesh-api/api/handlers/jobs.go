package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/domain/entities"
	jobsuc "github.com/sayeed1999/freelance-bangladesh/use_cases/jobs_uc"
)

type CreateJobUseCase interface {
	CreateJob(ctx context.Context, request jobsuc.CreateJobRequest) (*jobsuc.CreateJobResponse, error)
}

type GetJobsUseCase interface {
	GetJobs(ctx context.Context) []entities.Job
}

func CreateJobHandler(useCase CreateJobUseCase) fiber.Handler {
	return func(c fiber.Ctx) error {
		var ctx = c.UserContext()

		var request = jobsuc.CreateJobRequest{}

		err := c.Bind().Body(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.CreateJob(ctx, request)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func GetJobsHandler(useCase GetJobsUseCase) fiber.Handler {
	return func(c fiber.Ctx) error {

		var ctx = c.UserContext()

		Jobs := useCase.GetJobs(ctx)
		return c.Status(fiber.StatusOK).JSON(Jobs)
	}
}
