package assignments

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments/assign"
	assignmentlist "github.com/sayeed1999/freelance-bangladesh/features/assignments/assignment-list"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments/review"
	reviewlist "github.com/sayeed1999/freelance-bangladesh/features/assignments/review-list"
	"github.com/sayeed1999/freelance-bangladesh/features/assignments/submit"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

func RegisterAssignmentRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	assignTalentUseCase := assign.NewAssignTalentUseCase()
	assignmentListUseCase := assignmentlist.NewAssignmentListUseCase()
	submitWorkUseCase := submit.NewSubmitWorkUseCase()
	reviewWorkUseCase := review.NewReviewWorkUseCase()
	reviewListUseCase := reviewlist.NewReviewListUseCase()

	assignments := rg.Group("/assignments")
	{
		assignments.POST(
			"",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			assign.AssignTalentHandler(assignTalentUseCase),
		)

		assignments.GET(
			"",
			middlewares.Authorize(string(enums.ROLE_TALENT)),
			assignmentlist.AssignmentListHandler(assignmentListUseCase),
		)

		assignments.PATCH(
			"/:assignmentid",
			middlewares.Authorize(
				string(enums.ROLE_TALENT)),
			submit.SubmitWorkHandler(submitWorkUseCase),
		)

		assignments.POST(
			"/:assignmentid/reviews",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			review.ReviewWorkHandler(reviewWorkUseCase),
		)

		assignments.GET(
			"/:assignmentid/reviews",
			middlewares.Authorize(string(enums.ROLE_CLIENT), string(enums.ROLE_TALENT)),
			reviewlist.ReviewListHandler(reviewListUseCase),
		)
	}

	return assignments
}
