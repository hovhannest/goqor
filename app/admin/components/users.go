package components

import (
	"github.com/qor/admin"
	"github.com/qor/roles"
	"goqor1.0/app/models"
)

func ConfigUsers(Admin *admin.Admin) {
	user := Admin.AddResource(&models.User{})

	// Set attributes will be shown for the edit page, similar to new page
	user.EditAttrs("-Password", "-role")

	user.Meta(&admin.Meta{Name: "Email", Permission: roles.Allow(roles.Update, "newRole")})

}
