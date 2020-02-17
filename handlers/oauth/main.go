//package main
//
//import (
//	"fmt"
//	"net/http"
//	"time"
//
//	"github.com/davecgh/go-spew/spew"
//	"golang.org/x/oauth2"
//)
//
//func main() {
//	//ctx := context.Background()
//	//conf := &clientcredentials.Config{
//	//	ClientID:     "mizxg8mxfibi9cy35xc1z8m9ge",
//	//	ClientSecret: "9rixankt4ff5fqdu74dcbowimy",
//	//	TokenURL:     "http://localhost:8065/oauth/access_token",
//	//	EndpointParams: map[string][]string{
//	//		"authorize_url": []string{"http://localhost:8065/oauth/authorize"},
//	//		//"authorize_url": []string{"/oauth/authorize"},
//	//	},
//	//}
//	//
//	//token, err := conf.Token(ctx)
//	//spew.Dump(err)
//	//spew.Dump(token)
//
//	ctx := context.Background()
//
//	conf := &oauth2.Config{
//		ClientID:     "mizxg8mxfibi9cy35xc1z8m9ge",
//		ClientSecret: "9rixankt4ff5fqdu74dcbowimy",
//		Endpoint: oauth2.Endpoint{
//			TokenURL: "http://localhost:8065/oauth2/access_token",
//			AuthURL:  "http://localhost:8065/oauth2/authorize",
//		},
//		RedirectURL: "http://localhost:3000/mattermost/callback",
//	}
//
//	// Redirect user to consent page to ask for permission
//	// for the scopes specified above.
//	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
//	fmt.Printf("Visit the URL for the oauth dialog: %v", url)
//
//	// Use the authorization code that is pushed to the redirect
//	// URL. Exchange will do the handshake to retrieve the
//	// initial access token. The HTTP Client returned by
//	// conf.Client will refresh the token as necessary.
//	var code string
//	if _, err := fmt.Scan(&code); err != nil {
//		log.Fatal(err)
//	}
//
//	// Use the custom HTTP client when requesting a token.
//	httpClient := &http.Client{Timeout: 2 * time.Second}
//	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
//
//	tok, err := conf.Exchange(ctx, code)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	client := conf.Client(ctx, tok)
//	spew.Dump(client)
//}

package main

import (
	"fmt"

	"goscrum/server/controllers"
	"goscrum/server/db"
	"goscrum/server/gateway"
	"goscrum/server/services"
)

func main() {
	fmt.Println("Request Initiated")

	db := db.DbClient(true)

	defer db.Close()

	service := services.NewWorkspaceService(db)

	controller := controllers.NewAuthController(service)
	router := gateway.NewAPIRouter()

	router.Get("/oauth/mattermost/{workspaceId}/login", controller.MattermostLogin)
	router.Get("/oauth/mattermost/callback", controller.MattermostOauth)

	apiGateway := gateway.NewGateway()
	apiGateway.StartAPI(router)
}
