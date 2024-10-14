package auth

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

type identityManager interface {
	CreateUser(
		ctx context.Context,
		user gocloak.User,
		password string,
		role string,
		phone string) (*gocloak.User, error)
}
