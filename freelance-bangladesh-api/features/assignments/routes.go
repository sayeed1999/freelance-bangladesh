package assignments

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments/assign"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments/review"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments/submit"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

func RegisterAssignmentRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	assignTalentUseCase := assign.NewAssignTalentUseCase()
	submitWorkUseCase := submit.NewSubmitWorkUseCase()
	reviewWorkUseCase := review.NewReviewWorkUseCase()

	jobs := rg.Group("/assignments")
	{
		jobs.POST(
			"",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			assign.AssignTalentHandler(assignTalentUseCase),
		)

		jobs.PATCH(
			"/:assignmentid",
			middlewares.Authorize(
				string(enums.ROLE_TALENT)),
			submit.SubmitWorkHandler(submitWorkUseCase),
		)

		jobs.POST(
			"/:assignmentid/reviews",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			review.ReviewWorkHandler(reviewWorkUseCase),
		)
	}

	return jobs
}
