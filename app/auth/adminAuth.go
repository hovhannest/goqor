package auth

import (
	"fmt"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
	"net/http"
)


func init() {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil
	})
}

type AdminAuth struct {
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := App.Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
		}
		return qorCurrentUser
	}
	return nil
}