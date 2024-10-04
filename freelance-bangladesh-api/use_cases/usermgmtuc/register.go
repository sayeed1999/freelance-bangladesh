package usermgmtuc

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator/v10"
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

	roleName, err := uc.isRoleAllowedToSignup(request.Role)
	if err != nil {
		return nil, err
	}

	var user = gocloak.User{
		Email:         gocloak.StringP(request.Email),
		Username:      gocloak.StringP(request.Email),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		EmailVerified: gocloak.BoolP(false),
		Enabled:       gocloak.BoolP(true),
		RealmRoles:    &[]string{roleName},
	}

	// checking should the email be verified for this user
	if roleName == "client" {
		user.EmailVerified = gocloak.BoolP(true)
	} else if roleName == "talent" {
		user.EmailVerified = gocloak.BoolP(false)
	}

	userResponse, err := uc.identityManager.CreateUser(
		ctx, user, request.Password, roleName, request.MobileNumber)
	if err != nil {
		return nil, err
	}

	var response = &RegisterResponse{User: userResponse}
	return response, nil
}

func (uc *registerUseCase) isRoleAllowedToSignup(role string) (string, error) {
	roleNameLowerCase := strings.ToLower(role)
	err := ""

	switch roleNameLowerCase {
	case "": // force role = 'talent' if user doesn't specify role
		roleNameLowerCase = string(enums.ROLE_TALENT)
	case string(enums.ROLE_CLIENT):
	case string(enums.ROLE_TALENT):
		break
	default:
		err = "unable to signup user other than client or talent"
	}

	if len(err) > 0 {
		return "", fmt.Errorf(err)
	}
	return roleNameLowerCase, nil
}
