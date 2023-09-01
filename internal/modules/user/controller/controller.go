package controller

import (
	"context"
	"github.com/KuYaki/waffler_server/internal/infrastructure/component"
	"github.com/KuYaki/waffler_server/internal/infrastructure/responder"
	"github.com/KuYaki/waffler_server/internal/infrastructure/tools/cryptography"
	"github.com/KuYaki/waffler_server/internal/models"
	"github.com/KuYaki/waffler_server/internal/modules/message"
	"github.com/KuYaki/waffler_server/internal/modules/user/service"
	"github.com/gin-gonic/gin"
	"github.com/ptflp/godecoder"
	"net/http"
)

type Userer interface {
	Info(c *gin.Context)
	Save(c *gin.Context)
}
type User struct {
	service service.Userer
	jwt     cryptography.TokenManager
	responder.Responder
	godecoder.Decoder
}

func NewUserController(service service.Userer, components *component.Components) Userer {
	return &User{service: service, Responder: components.Responder, Decoder: components.Decoder, jwt: components.TokenManager}
}

func (a *User) Info(c *gin.Context) {
	claims, err := a.jwt.ParseTokenForHTTP(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	userInfo, err := a.service.GetUserInfo(context.Background(), claims.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, userInfo)
}

func (a *User) Save(c *gin.Context) {
	claims, err := a.jwt.ParseTokenForHTTP(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	var userSave message.UserInfo
	err = c.BindJSON(&userSave)
	if err != nil {
		return
	}

	err = a.service.Update(context.Background(), models.User{
		ID:          claims.ID,
		ParserToken: userSave.Parser.Token,
		ParserType:  userSave.Parser.Type,
		Locale:      userSave.Locale,
	})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
