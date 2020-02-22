package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"goscrum/server/constants"
	"os"
	"strings"

	"goscrum/server/models"
	"goscrum/server/services"
	"goscrum/server/util"

	"github.com/aws/aws-lambda-go/events"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
)

type AuthController struct {
	workspaceService services.WorkspaceService
}

func NewAuthController(workspaceService services.WorkspaceService) AuthController {
	return AuthController{workspaceService: workspaceService}
}

func (a *AuthController) getMatterMostOAuthClient(workspace models.Workspace) *oauth2.Config {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
		fmt.Println(pair[1])
	}
	// TODO -- remove extra slash at the end of workspace URL when creating workspace
	fmt.Println(fmt.Sprintf("%s/oauth/mattermost/callback", os.Getenv(constants.ApiUrl)))
	conf := &oauth2.Config{
		ClientID:     workspace.ClientID,
		ClientSecret: workspace.ClientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/oauth/access_token", workspace.URL),
			AuthURL:  fmt.Sprintf("%s/oauth/authorize", workspace.URL),
		},
		RedirectURL: fmt.Sprintf("%s/oauth/mattermost/callback", os.Getenv(constants.ApiUrl)),
	}
	return conf
}

func (a *AuthController) MattermostLogin(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	workspaceId, err := util.GetStringKey(req.PathParameters, "workspaceId")

	if err != nil {
		return util.ServerError(err)
	}
	workspace, err := a.workspaceService.GetWorkspace(workspaceId)

	if err != nil {
		return util.ServerError(err)
	}

	// TODO -- remove extra slash at the end of workspace URL when creating workspace
	conf := a.getMatterMostOAuthClient(workspace)

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL(workspace.ID, oauth2.AccessTypeOffline)

	fmt.Println(url)
	return util.Redirect(url)
}

func (a *AuthController) MattermostOauth(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Use the authorization code that is pushed to the redirect
	//URL. Exchange will do the handshake to retrieve the
	//initial access token. The HTTP Client returned by
	//conf.Client will refresh the token as necessary.
	code, err := util.GetStringKey(request.QueryStringParameters, "code")

	if err != nil {
		return util.ServerError(err)
	}

	workspaceId, err := util.GetStringKey(request.QueryStringParameters, "state")
	if err != nil {
		return util.ServerError(err)
	}
	workspace, err := a.workspaceService.GetWorkspace(workspaceId)
	if err != nil {
		return util.ServerError(err)
	}

	conf := a.getMatterMostOAuthClient(workspace)
	ctx := context.Background()

	tok, err := conf.Exchange(ctx, code, oauth2.AccessTypeOffline)
	if err != nil {
		return util.ServerError(err)
	}

	workspace.AccessToken = tok.AccessToken
	workspace.RefreshToken = tok.RefreshToken
	workspace.Expiry = &tok.Expiry
	if workspace.PersonalToken == "" {
		workspace.PersonalToken = base64.StdEncoding.EncodeToString(uuid.NewV4().Bytes())
	}

	err = a.workspaceService.Update(workspaceId, workspace)
	// TODO for now, show error message, later redirect to beautiful page.
	if err != nil {
		return util.ServerError(err)
	}

	// TODO redirect to application workspace page with hash access token
	return util.Redirect("https://google.com")
}
