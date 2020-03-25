package main

import (
    "github.com/dgrijalva/jwt-go"
	"fmt"
	"io"
	"os"
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"crypto/rsa"
)

var path = "./sample_key"

func main() {
	pemString := pemKey()
	fmt.Println(reflect.TypeOf(pemString))
    mySigningKey := pemString

    // Create the token
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    // Set some claims
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
    fmt.Println(tokenString, err)
}

func readFile(){
    // buka file
    var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
    if isError(err) { return }
    defer file.Close()

    // baca file
    var text = make([]byte, 1024)
    for {
        n, err := file.Read(text)
        if err != io.EOF {
            if isError(err) { break }
        }
        if n == 0 {
            break
        }
    }
    if isError(err) { return }

    fmt.Println("==> file berhasil dibaca")
    fmt.Println(string(text))
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func pemKey() *rsa.PrivateKey {
	pemString := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDLets8+7M+iAQAqN/5BVyCIjhTQ4cmXulL+gm3v0oGMWzLupUS
v8KPA+Tp7dgC/DZPfMLaNH1obBBhJ9DhS6RdS3AS3kzeFrdu8zFHLWF53DUBhS92
5dCAEuJpDnNizdEhxTfoHrhuCmz8l2nt1pe5eUK2XWgd08Uc93h5ij098wIDAQAB
AoGAHLaZeWGLSaen6O/rqxg2laZ+jEFbMO7zvOTruiIkL/uJfrY1kw+8RLIn+1q0
wLcWcuEIHgKKL9IP/aXAtAoYh1FBvRPLkovF1NZB0Je/+CSGka6wvc3TGdvppZJe
rKNcUvuOYLxkmLy4g9zuY5qrxFyhtIn2qZzXEtLaVOHzPQECQQDvN0mSajpU7dTB
w4jwx7IRXGSSx65c+AsHSc1Rj++9qtPC6WsFgAfFN2CEmqhMbEUVGPv/aPjdyWk9
pyLE9xR/AkEA2cGwyIunijE5v2rlZAD7C4vRgdcMyCf3uuPcgzFtsR6ZhyQSgLZ8
YRPuvwm4cdPJMmO3YwBfxT6XGuSc2k8MjQJBAI0+b8prvpV2+DCQa8L/pjxp+VhR
Xrq2GozrHrgR7NRokTB88hwFRJFF6U9iogy9wOx8HA7qxEbwLZuhm/4AhbECQC2a
d8h4Ht09E+f3nhTEc87mODkl7WJZpHL6V2sORfeq/eIkds+H6CJ4hy5w/bSw8tjf
sz9Di8sGIaUbLZI2rd0CQQCzlVwEtRtoNCyMJTTrkgUuNufLP19RZ5FpyXxBO5/u
QastnN77KfUwdj3SJt44U/uh1jAIv4oSLBr8HYUkbnI8
-----END RSA PRIVATE KEY-----`

	block, _ := pem.Decode([]byte(pemString))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}