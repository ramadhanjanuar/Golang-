package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"golang.org/x/oauth2/jwt"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	g := gin.Default()

	auth := g.Group("/oauth2")
	{
		ctx := context.Background()
		conf := &jwt.Config{
			Email: "xxx@developer.com",
			// The contents of your RSA private key or your PEM file
			// that contains a private key.
			// If you have a p12 file instead, you
			// can use `openssl` to export the private key into a pem file.
			//
			//    $ openssl pkcs12 -in key.p12 -out key.pem -nodes
			//
			// It only supports PEM containers with no passphrase.
			PrivateKey: []byte("-----BEGIN RSA PRIVATE KEY-----..."),
			Subject:    "user@example.com",
			TokenURL:   "https://provider.com/o/oauth2/token",
		}
		// Initiate an http.Client, the following GET request will be
		// authorized and authenticated on the behalf of user@example.com.
		client := conf.Client(ctx)
		fmt.Print("teeeeee", client)
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

func ExampleJWTConfig() {
	ctx := context.Background()
	conf := &jwt.Config{
		Email: "xxx@developer.com",
		// The contents of your RSA private key or your PEM file
		// that contains a private key.
		// If you have a p12 file instead, you
		// can use `openssl` to export the private key into a pem file.
		//
		//    $ openssl pkcs12 -in key.p12 -out key.pem -nodes
		//
		// It only supports PEM containers with no passphrase.
		PrivateKey: []byte("-----BEGIN RSA PRIVATE KEY-----..."),
		Subject:    "user@example.com",
		TokenURL:   "https://provider.com/o/oauth2/token",
	}
	// Initiate an http.Client, the following GET request will be
	// authorized and authenticated on the behalf of user@example.com.
	client := conf.Client(ctx)
}
