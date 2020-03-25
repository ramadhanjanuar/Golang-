package main

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	ginserver "github.com/go-oauth2/gin-server"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

const (
	userkey = "user"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore(TokenStore))

	manager.MapAccessGenerate(NewJWTAccessGenerate([]byte(JwtKey), jwt.SigningMethodRS256))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set(ClientID, &models.Client{
		ID:     ClientID,
		Secret: ClientSecret,
		Domain: Domain,
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service

	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
	ginserver.SetPasswordAuthorizationHandler(AuthHandler)

	g := gin.Default()

	auth := g.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	api := g.Group("/api")
	{
		api.Use(ginserver.HandleTokenVerify())
		api.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusOK, "not found")
		})
	}

	g.Run(":9096")
}

// AuthHandler handler for authentication
func AuthHandler(username, password string) (userID string, err error) {
	hashQwerty := "$2y$08$mHMknlurxKSvJ.Eh63kOD.ySX4XzfRUGU8tc2uQQWjwpPEkTRYam6"
	passwordMatch := CheckPasswordHash(password, hashQwerty)
	if username == "istiq" && passwordMatch {
		userID = "istiq"
	}
	return
}

// CheckPasswordHash function for checking password match or not
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Salt+password))
	return err == nil
}
