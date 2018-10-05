package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/media"
	"github.com/qor/media/oss"
	"time"
)

type User struct {
	gorm.Model
	Email                  string `gorm:"column:email;unique_index;not null"`
	Password               string
	Name                   string `form:"name"`
	Gender                 string
	Role                   string `gorm:"default:'newUser'"`
	Birthday               *time.Time
	Balance                float32
	DefaultBillingAddress  uint `form:"default-billing-address"`
	DefaultShippingAddress uint `form:"default-shipping-address"`
	Addresses              []Address
	Avatar                 AvatarImageStorage

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time

	// Accepts
	AcceptPrivate bool `form:"accept-private"`
	AcceptLicense bool `form:"accept-license"`
	AcceptNews    bool `form:"accept-news"`

	AuthIdentity []UserAuthIdentity `gorm:"foreignkey:UID"`
}

type UserAuthIdentity struct {
	gorm.Model
	auth_identity.Basic
	auth_identity.SignLogs
}


func (UserAuthIdentity) TableName() string {
	return "basics"
}


func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US"}
}

type AvatarImageStorage struct{ oss.OSS }

func (AvatarImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 50, Height: 50},
		"middle": {Width: 120, Height: 120},
		"big":    {Width: 320, Height: 320},
	}
}