package apiProvider

import (
	"fmt"
	"github.com/qor/auth"
	"github.com/qor/auth/claims"
	"github.com/qor/auth/providers/password/encryptor"
	"github.com/qor/auth/providers/password/encryptor/bcrypt_encryptor"
	"net/http"
)

func New(config *Config) *APIProvider {
	if config == nil {
	config = &Config{}
}

	if config.Encryptor == nil {
		config.Encryptor = bcrypt_encryptor.New(&bcrypt_encryptor.Config{})
	}

	provider := &APIProvider{Config: config}

	//if config.ConfirmMailer == nil {
	//	config.ConfirmMailer = DefaultConfirmationMailer
	//}
	//
	//if config.ConfirmHandler == nil {
	//	config.ConfirmHandler = DefaultConfirmHandler
	//}
	//
	//if config.ResetPasswordMailer == nil {
	//	config.ResetPasswordMailer = DefaultResetPasswordMailer
	//}
	//
	//if config.ResetPasswordHandler == nil {
	//	config.ResetPasswordHandler = DefaultResetPasswordHandler
	//}
	//
	//if config.RecoverPasswordHandler == nil {
	//	config.RecoverPasswordHandler = DefaultRecoverPasswordHandler
	//}

	if config.AuthorizeHandler == nil {
		config.AuthorizeHandler = DefaultAuthorizeHandler
	}

	if config.RegisterHandler == nil {
		config.RegisterHandler = DefaultRegisterHandler
	}

	return provider
}

type LoginData struct {
	Login string `json:"Login"`
	Password string `json:"Password"`
}

// Config facebook Config
type Config struct {
	Confirmable    bool

	ConfirmMailer  func(email string, context *auth.Context, claims *claims.Claims, currentUser interface{}) error
	ConfirmHandler func(*auth.Context) error

	ResetPasswordMailer    func(email string, context *auth.Context, claims *claims.Claims, currentUser interface{}) error
	ResetPasswordHandler   func(*auth.Context) error
	RecoverPasswordHandler func(*auth.Context) error

	Encryptor        encryptor.Interface
	AuthorizeHandler func(*auth.Context)
	RegisterHandler  func(*auth.Context)
}

// PhoneProvider provide login with phone method
type APIProvider struct {
	*Config
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
func (provider APIProvider) Login(context *auth.Context) {
	if context.Request.Method != "POST" {
		context.Writer.WriteHeader(http.StatusNotFound)
		context.Writer.Write([]byte("404 - Use POST for login!"))
		return
	}
	provider.AuthorizeHandler(context)
}

// Logout implemented logout with phone provider
func (APIProvider) Logout(context *auth.Context) {
	fmt.Print("APIProvider: Logout ", context)
}

// Register implemented register with phone provider
func (provider APIProvider) Register(context *auth.Context) {
	if context.Request.Method != "POST" {
		context.Writer.WriteHeader(http.StatusNotFound)
		context.Writer.Write([]byte("404 - Use POST for Register!"))
		return
	}
	provider.RegisterHandler(context)
}

// Callback implement Callback with phone provider
func (APIProvider) Callback(context *auth.Context) {
	fmt.Print("APIProvider: Callback ", context)
}

// ServeHTTP implement ServeHTTP with phone provider
func (APIProvider) ServeHTTP(context *auth.Context) {
	fmt.Print("APIProvider: ServeHTTP", context)
}

