package auth

import (
	"errors"
	"go-webapp/common"
	"go-webapp/middleware/session"
	"go-webapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Authorizer structures contain the store of user session cookies a reference
// to a backend storage system.
type Authorizer struct {
	backend AuthBackend //auth backend TODO
}

//todo https://github.com/apexskier/httpauth/blob/master/auth.go

// NewAuthorizer returns a new Authorizer given an AuthBackend, a session store
func NewAuthorizer(b AuthBackend) *Authorizer {
	a := new(Authorizer)

	a.backend = b
	log.WithFields(log.Fields{
		"backend": a.backend,
	}).Info("Authorizer--NewAuthorizer")
	return a
}

// Login logs a user in with status code 200
func (a Authorizer) Login(context *gin.Context, email string, password string) {

	userModel, err := a.backend.FetchUser(&models.User{Email: email})

	if err != nil {
		context.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.CheckPassword(password) != nil {
		context.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	//Manage session here
	if sessionError := session.Authenticate(context, userModel); !sessionError {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

//Register the user to the system
func (a Authorizer) Register(data interface{}) error {
	err := a.backend.SaveUser(data)
	return err
}

func (a Authorizer) Logout(context *gin.Context) {

}

func (a Authorizer) Update(context *gin.Context) {

}
