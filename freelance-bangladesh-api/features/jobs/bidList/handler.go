package bidlist

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BidListUseCase interface {
	BidList(ctx context.Context, jobID string) ([]BidResponse, error)
}

func BidListHandler(useCase BidListUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		jobID := c.Param("jobid")
		if jobID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "job id is required"})
			return
		}

		// Call the use case to list bids for the specified job
		bids, err := useCase.BidList(c.Request.Context(), jobID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, bids)
	}
}
