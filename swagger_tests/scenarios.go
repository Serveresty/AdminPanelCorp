package swagger_tests

import (
	apiclient "AdminPanelCorp/docs/output_swag/client"
	"AdminPanelCorp/docs/output_swag/client/auth"
	"AdminPanelCorp/docs/output_swag/client/page"
	modl "AdminPanelCorp/docs/output_swag/models"
	"fmt"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
)

func Scenario1(clientt *apiclient.AdminPanelAPI) error {
	a := &modl.ModelsUser{
		Email:    "cs.go.12228@gmail.com",
		Password: "tCN9yjOuh5",
	}

	signinact, err := clientt.Auth.PostAuthLoginForm(auth.NewPostAuthLoginFormParams().WithInput(a))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", signinact)

	hom, err := clientt.Page.GetPageHomepage(page.NewGetPageHomepageParams(), httptransport.BearerToken(strings.Split(signinact.Authorization, " ")[1]))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", hom)

	logoutt, err := clientt.Auth.PostAuthLogoutForm(auth.NewPostAuthLogoutFormParams(), httptransport.BearerToken(strings.Split(signinact.Authorization, " ")[1]))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", logoutt)
	return nil
}

func Scenario2(clientt *apiclient.AdminPanelAPI) error {
	a := &modl.ModelsUser{
		Email:    "cs.go.12228@gmail.com",
		Password: "tCN9yjOuh5",
	}

	signinact, err := clientt.Auth.PostAuthLoginForm(auth.NewPostAuthLoginFormParams().WithInput(a))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", signinact)

	admp, err := clientt.Page.GetPageAdmin(page.NewGetPageAdminParams(), httptransport.BearerToken("sda"))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", admp)

	logoutt, err := clientt.Auth.PostAuthLogoutForm(auth.NewPostAuthLogoutFormParams(), httptransport.BearerToken(strings.Split(signinact.Authorization, " ")[1]))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", logoutt)
	return nil
}

func Scenario3(clientt *apiclient.AdminPanelAPI) error {
	a := &modl.ModelsUser{
		Email:    "cs.go.12228@gmail.com",
		Password: "tCN9yjOuh5",
	}

	signinact, err := clientt.Auth.PostAuthLoginForm(auth.NewPostAuthLoginFormParams().WithInput(a))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", signinact)

	admp, err := clientt.Page.GetPageAdmin(page.NewGetPageAdminParams(), httptransport.BearerToken(strings.Split(signinact.Authorization, " ")[1]))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", admp)

	logoutt, err := clientt.Auth.PostAuthLogoutForm(auth.NewPostAuthLogoutFormParams(), httptransport.BearerToken(strings.Split(signinact.Authorization, " ")[1]))
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", logoutt)
	return nil
}
