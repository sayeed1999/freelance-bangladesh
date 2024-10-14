package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	getclients "github.com/sayeed1999/freelance-bangladesh/features/admin/getClients"
	gettalents "github.com/sayeed1999/freelance-bangladesh/features/admin/getTalents"
	updateclient "github.com/sayeed1999/freelance-bangladesh/features/admin/updateClient"
	updatetalent "github.com/sayeed1999/freelance-bangladesh/features/admin/updateTalent"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

func RegisterAdminRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	getClientsUseCase := getclients.NewGetClientsUseCase()
	getTalentsUseCase := gettalents.NewGetTalentsUseCase()
	updateClientUseCase := updateclient.NewUpdateClientUseCase()
	updateTalentUseCase := updatetalent.NewUpdateTalentUseCase()

	adminRoutes := rg.Group("/admin-dashboard")
	{
		adminRoutes.Use(middlewares.Authorize(string((enums.ROLE_ADMIN))))

		adminRoutes.GET("/clients",
			getclients.GetClientsHandler(getClientsUseCase))

		adminRoutes.GET("/talents",
			gettalents.GetTalentsHandler(getTalentsUseCase))

		adminRoutes.POST("/clients",
			updateclient.UpdateClientHandler(updateClientUseCase))

		adminRoutes.POST("/talents",
			updatetalent.UpdateTalentHandler(updateTalentUseCase))

	}

	return adminRoutes
}
