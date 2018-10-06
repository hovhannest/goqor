package apiProvider

import (
	"encoding/json"
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/qor/utils"
	"github.com/qor/session"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var DefaultAuthorizeHandler = func(context *auth.Context) {

	provider, _ := context.Provider.(*APIProvider)
	r := context.Request
	w := context.Writer
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Invalid Account!"))
		return
	}
	// Unmarshal
	var msg LoginData
	err = json.Unmarshal(b, &msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Invalid Account!"))
		return
	}

	// Login
	var (
		authInfo    auth_identity.Basic
		tx          = context.Auth.GetDB(r)
	)

	authInfo.Provider = "password"
	authInfo.UID = strings.TrimSpace(msg.Login)

	if tx.Model(context.Auth.AuthIdentityModel).Where(authInfo).Scan(&authInfo).RecordNotFound() {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Invalid Account!"))
		return
	}

	if provider.Config.Confirmable && authInfo.ConfirmedAt == nil {
		currentUser, _ := context.Auth.UserStorer.Get(authInfo.ToClaims(), context)
		provider.Config.ConfirmMailer(authInfo.UID, context, authInfo.ToClaims(), currentUser)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Account Unconfirmed!"))
		return
	}

	if err := provider.Encryptor.Compare(authInfo.EncryptedPassword, strings.TrimSpace(msg.Password)); err == nil {
		cl := authInfo.ToClaims()
		tk :=  context.SessionStorer.SignedToken(cl)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Logged in ! token = " + tk))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Wrong Password!"))
	return
}


// DefaultRegisterHandler default register handler
var DefaultRegisterHandler = func(context *auth.Context) {
	var (
		err         error
		currentUser interface{}
		schema      auth.Schema
		authInfo    auth_identity.Basic
		r         	= context.Request
		w 			= context.Writer
		tx          = context.Auth.GetDB(r)
		provider, _ = context.Provider.(*APIProvider)
	)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Invalid Account!"))
		return
	}
	// Unmarshal
	var msg LoginData
	err = json.Unmarshal(b, &msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Invalid Account!"))
		return
	}

	authInfo.Provider = "password"
	authInfo.UID = strings.TrimSpace(msg.Login)

	if !tx.Model(context.Auth.AuthIdentityModel).Where(authInfo).Scan(&authInfo).RecordNotFound() {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Invalid Account!"))
		return
	}

	if authInfo.EncryptedPassword, err = provider.Encryptor.Digest(strings.TrimSpace(msg.Password)); err == nil {
		schema.Provider = authInfo.Provider
		schema.UID = authInfo.UID
		schema.Email = authInfo.UID
		schema.RawInfo = r

		currentUser, authInfo.UserID, err = context.Auth.UserStorer.Save(&schema, context)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something gone wrong!"))
			return
		}

		// create auth identity
		authIdentity := reflect.New(utils.ModelType(context.Auth.Config.AuthIdentityModel)).Interface()
		if err = tx.Where(authInfo).FirstOrCreate(authIdentity).Error; err == nil {
			if provider.Config.Confirmable {
				context.SessionStorer.Flash(w, r, session.Message{Message: ConfirmFlashMessage, Type: "success"})
				err = provider.Config.ConfirmMailer(schema.Email, context, authInfo.ToClaims(), currentUser)
			}
			cl := authInfo.ToClaims()
			tk :=  context.SessionStorer.SignedToken(cl)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 - Logged in ! token = " + tk))
			return
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something gone wrong!"))
	return
}
