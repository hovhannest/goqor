package apiProvider

import (
	"fmt"
	"github.com/qor/auth"
)

func New(config *Config) *APIProvider {
	return &APIProvider{}
}

// Config facebook Config
type Config struct {
}

// PhoneProvider provide login with phone method
type APIProvider struct {
}

// GetName return provider name
func (APIProvider) GetName() string {
	return "api"
}

// ConfigAuth config auth
func (APIProvider) ConfigAuth(*auth.Auth) {
	fmt.Print("APIProvider: ConfigAuth")
}

// Login implemented login with phone provider
func (APIProvider) Login(context *auth.Context) {
	fmt.Print("APIProvider: Login ", context)
}

// Logout implemented logout with phone provider
func (APIProvider) Logout(context *auth.Context) {
	fmt.Print("APIProvider: Logout ", context)
}

// Register implemented register with phone provider
func (APIProvider) Register(context *auth.Context) {
	fmt.Print("APIProvider: Register ", context)
}

// Callback implement Callback with phone provider
func (APIProvider) Callback(context *auth.Context) {
	fmt.Print("APIProvider: Callback ", context)
}

// ServeHTTP implement ServeHTTP with phone provider
func (APIProvider) ServeHTTP(context *auth.Context) {
	fmt.Print("APIProvider: ServeHTTP", context)
}

