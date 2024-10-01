package usermgmtuc

import (
	"context"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sayeed1999/freelance-bangladesh/shared/enums"
)

type RegisterRequest struct {
	Password     string `valdate:"required"`
	FirstName    string `validate:"min=1,max=30"`
	LastName     string `validate:"min=1,max=30"`
	Email        string `validate:"required,email"`
	MobileNumber string
	Role         string
}

type RegisterResponse struct {
	User *gocloak.User
}

type registerUseCase struct {
	identityManager identityManager
}

func NewRegisterUseCase(im identityManager) *registerUseCase {
	return &registerUseCase{
		identityManager: im,
	}
}

func (uc *registerUseCase) Register(ctx context.Context, request RegisterRequest) (*RegisterResponse, error) {
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}

	// Force role = 'talent' if user doesn't specify role!
	if request.Role == "" {
		request.Role = string(enums.ROLE_TALENT)
	}

	var user = gocloak.User{
		Email:         gocloak.StringP(request.Email),
		Username:      gocloak.StringP(request.Email),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		EmailVerified: gocloak.BoolP(false),
		Enabled:       gocloak.BoolP(true),
		RealmRoles:    &[]string{request.Role},
	}

	var roleNameLowerCase = strings.ToLower(request.Role)
	switch roleNameLowerCase {
	case string(enums.ROLE_CLIENT):
	case string(enums.ROLE_TALENT):
		break
	default:
		return nil, errors.Wrap(err, "unable to signup user other than client or talent")
	}

	userResponse, err := uc.identityManager.CreateUser(
		ctx, user, request.Password, roleNameLowerCase, request.MobileNumber)
	if err != nil {
		return nil, err
	}

	var response = &RegisterResponse{User: userResponse}
	return response, nil
}
