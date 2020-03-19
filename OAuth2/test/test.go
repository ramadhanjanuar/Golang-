package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/models"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var EXPIRES_AT = time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix()
var TIME_NOW = time.Now().Unix()
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("CORE 3.0")
var ISSUER = "https://auth.99group.co"

func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		HandlerLogin(w, r, clientStore)
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9096", nil))
}

// HandlerLogin handler
func HandlerLogin(w http.ResponseWriter, r *http.Request, clientStore *store.ClientStore) {
	if r.Method != "POST" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	user := r.FormValue("user")
	password := r.FormValue("password")
	authenticated := authentication(user, password)
	if !authenticated {
		return
	}

	clientID := uuid.New().String()[:8]
	clientSecret := uuid.New().String()[:8]
	err := clientStore.Set(clientID, &models.Client{
		ID:     clientID,
		Secret: clientSecret,
	})

	claims := jwt.StandardClaims{
		Subject:   "uuid of the user",
		Audience:  clientID,
		Issuer:    ISSUER,
		ExpiresAt: EXPIRES_AT,
		NotBefore: TIME_NOW,
		IssuedAt:  TIME_NOW,
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	tokenString := generateJwtToken(claims)
	json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientID, "CLIENT_SECRET": clientSecret, "JWT": tokenString})
	return
}

func authentication(user string, password string) bool {
	_ = user
	salt := "rXI$L@bdND"
	hashQwerty := "$2y$08$mHMknlurxKSvJ.Eh63kOD.ySX4XzfRUGU8tc2uQQWjwpPEkTRYam6"
	match := CheckPasswordHash(salt+password, hashQwerty)

	return match
}

// generatedJwt method
func generateJwtToken(claims jwt.StandardClaims) string {
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

	if err != nil {
		return string(http.StatusBadRequest)
	}
	return signedToken
}

// CheckPasswordHash method
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
