package jobs

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	bidjob "github.com/sayeed1999/freelance-bangladesh/features/jobs/bidJob"
	bidlist "github.com/sayeed1999/freelance-bangladesh/features/jobs/bidList"
	createjob "github.com/sayeed1999/freelance-bangladesh/features/jobs/createJob"
	getjobs "github.com/sayeed1999/freelance-bangladesh/features/jobs/getJobs"
	pendingreviewlist "github.com/sayeed1999/freelance-bangladesh/features/jobs/pending-review-list"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

func RegisterJobRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	createJobUseCase := createjob.NewCreateJobUseCase()
	getJobsUseCase := getjobs.NewGetJobsUseCase()
	bidJobUseCase := bidjob.NewBidOnJobUseCase()
	bidListUseCase := bidlist.NewBidListUseCase()
	pendingreviewlistUseCase := pendingreviewlist.NewPendingReviewListUseCase()

	jobs := rg.Group("/jobs")
	{
		jobs.POST(
			"",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			createjob.CreateJobHandler(createJobUseCase),
		)

		jobs.GET(
			"",
			middlewares.PrivateCache(),
			middlewares.Authorize(
				string(enums.ROLE_ADMIN),
				string(enums.ROLE_CLIENT),
				string(enums.ROLE_TALENT)),
			getjobs.GetJobsHandler(getJobsUseCase),
		)

		jobs.POST(
			"/:jobid/bids",
			middlewares.Authorize(string(enums.ROLE_TALENT)),
			bidjob.BidOnJobHandler(bidJobUseCase),
		)

		jobs.GET(
			"/:jobid/bids",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			bidlist.BidListHandler(bidListUseCase),
		)

		jobs.GET(
			"/:jobid/pending-reviews",
			middlewares.Authorize(string(enums.ROLE_CLIENT)),
			pendingreviewlist.PendingReviewListHandler(pendingreviewlistUseCase),
		)
	}

	return jobs
}
