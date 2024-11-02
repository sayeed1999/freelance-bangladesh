package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sayeed1999/freelance-bangladesh/api/middlewares"
	auth "github.com/sayeed1999/freelance-bangladesh/features/auth/syncUser"
	"github.com/sayeed1999/freelance-bangladesh/infrastructure/identity"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

func RegisterUserManagementRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	syncUserUseCase := auth.NewSyncUserUseCase()
	identityManager := identity.NewIdentityManager()
	registerUseCase := NewRegisterUseCase(identityManager)

	users := rg.Group("/users")
	{
		users.Use(middlewares.NoStoreCache())

		users.POST("/sync-user",
			middlewares.Authorize(),
			auth.SyncUserHandler(syncUserUseCase))

		// N.B: client sigup considered admin route!
		users.POST("/client-signup",
			middlewares.Authorize(string(enums.ROLE_ADMIN)),
			RegisterClientHandler(registerUseCase))

		users.POST("/talent-signup",
			RegisterTalentHandler(registerUseCase))
	}

	return users
}
