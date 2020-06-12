package main

import (
	"contex"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

type AuthInterceptor struct {
	token string
}

// JWTAccessClaims jwt claims
type JWTAccessClaims struct {
	jwt.StandardClaims
}

// JwtKey for jwt key
const publicKey = `-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgHUt0lYhu3sdEKBbvuDyEygHgUxi
vuQdLDJJki7bbuxoqQcqw9tSwlbS8HHeypFblDJxnpmWM7WKoiKIBfbm8EIX8rHb
+5zDr99EB6N2bnjZYxssfPVI6O/syH2RusciA+mFjBtedJLIZ4ayW9MYjZZUusP6
TpPJEqtxY0v4Css/AgMBAAE=
-----END PUBLIC KEY-----`

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICWgIBAAKBgHUt0lYhu3sdEKBbvuDyEygHgUxivuQdLDJJki7bbuxoqQcqw9tS
wlbS8HHeypFblDJxnpmWM7WKoiKIBfbm8EIX8rHb+5zDr99EB6N2bnjZYxssfPVI
6O/syH2RusciA+mFjBtedJLIZ4ayW9MYjZZUusP6TpPJEqtxY0v4Css/AgMBAAEC
gYAT3GRePPFt+Iss4uADeIROsJb1v3bjax5ml5mzX3X5k/OyR0lTDz/pP2+JnV1w
Lbl8o01ictbo2kqI181K//GQxSFI9o4J9bXGf6WGuTtNW9R/4P/78whYpGscdqyN
oeNIlZ0AbRZbE8dXcIzfDOm4haN9x7tmOa8wHIxLWVmFYQJBAOa9yc5/qUl5e16U
/gyKSn/DxFVCXL7ZGqZxna2w5tRkS6iyw6fz3oeLadS+FI5AkLgUmWh/zn8h+G0l
EEbn0h0CQQCCAZpX/ko64wyn3fztNZ0V0rGEW2Rigd/bxzqPexLDyNviESkEy+yH
mk9i0UkN7KeY3oFVlUs1uA5JX4O+TJQLAkA61trHVbaNjfTvwX6nI/yrZ/iA5VNR
N3ucRgy3Zgz2zDye4DcUdROh7OMn5PDI9Z3x0w7rnuvBb6Ax9OmZqzPhAkA9BNOU
iV391aOTBrY3//oCzClVni+3rI2Nci0iOvnLuvK5YKSIA864tsyb7O8FAhaHwaei
q7fpNsI93f7PytGRAkBpFRMEKHN16uzSYUIfkS8VJ/J4ZpGuVmJQRE/i6NUOK067
nQRURIRpQ+cLRHwR/6QKeg18M3SeflnBciA7A5d/
-----END RSA PRIVATE KEY-----`

func NewAuthInterceptor(token string) *AuthInterceptor {
	return &AuthInterceptor{token}
}

func (interceptor *AuthInterceptor) verifyToken() (*JWTAccessClaims, error) {
	token, err := jwt.ParseWithClaims(
		interceptor.token,
		&JWTAccessClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodRSA)
			if !ok {
				return nil, fmt.Errorf("unexpected token sign")
			}
			verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
			if err != nil {
				return nil, err
			}
			return verifyKey, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*JWTAccessClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx contex.Contex,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) {
		log.Println("-> unary interceptor: ", info.FullMethod)
		return handler(ctx, req)
	}
}
