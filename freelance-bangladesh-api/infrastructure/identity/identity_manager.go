package identity

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type identityManager struct {
	baseUrl             string
	realm               string
	restApiClientId     string
	restApiClientSecret string
}

func NewIdentityManager() *identityManager {
	return &identityManager{
		baseUrl:             viper.GetString("Keycloak.BaseUrl"),
		realm:               viper.GetString("Keycloak.Realm"),
		restApiClientId:     viper.GetString("Keycloak.RestApi.ClientId"),
		restApiClientSecret: viper.GetString("Keycloak.RestApi.ClientSecret"),
	}
}

func (im *identityManager) loginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(im.baseUrl, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms"))

	token, err := client.LoginClient(ctx, im.restApiClientId, im.restApiClientSecret, im.realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login the rest client")
	}
	return token, nil
}

func (im *identityManager) CreateUser(ctx context.Context, user gocloak.User, password string, role string) (*gocloak.User, error) {

	token, err := im.loginRestApiClient(ctx)
	if err != nil {
		return nil, err
	}

	client := gocloak.NewClient(im.baseUrl)

	userId, err := client.CreateUser(ctx, token.AccessToken, im.realm, user)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create the user")
	}

	err = client.SetPassword(ctx, token.AccessToken, userId, im.realm, password, false)
	if err != nil {
		return nil, errors.Wrap(err, "unable to set the password for the user")
	}

	roleKeycloak, err := client.GetRealmRole(ctx, token.AccessToken, im.realm, role)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to get role by name: '%v'", role))
	}
	err = client.AddRealmRoleToUser(ctx, token.AccessToken, im.realm, userId, []gocloak.Role{
		*roleKeycloak,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to add a realm role to user")
	}

	userKeycloak, err := client.GetUserByID(ctx, token.AccessToken, im.realm, userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get recently created user")
	}

	return userKeycloak, nil
}
