package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

// AuthInterceptor server
type AuthInterceptor struct{}

// JWTAccessClaims jwt claims
type JWTAccessClaims struct {
	jwt.StandardClaims
	Scope string
}

// ContextKey for key context
type ContextKey int

// ClaimsKey claims key
const ClaimsKey ContextKey = iota

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

var staticToken = map[string]JWTAccessClaims{
	"fmk9i0UkN7KeY3oFVlUs1uA5JX4OTJQLAkA61trHVbaNjfTvwX6nIyrZiA5VNR": JWTAccessClaims{
		Scope: "INTERNAL",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 0,
		},
	},
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

// NewServerWithAuthInterceptor auth interceptor
func NewServerWithAuthInterceptor() *grpc.Server {
	interceptor := NewAuthInterceptor()
	return grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
	)
}

// Unary return unary
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx, err := interceptor.authorize(ctx)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metada is not provided")
	}
	return interceptor.getMetaDataAuthorization(ctx, md)
}

func (interceptor *AuthInterceptor) getMetaDataAuthorization(
	ctx context.Context,
	md metadata.MD,
) (context.Context, error) {
	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	return interceptor.getClaimsFromVerifyToken(ctx, values)
}

func (interceptor *AuthInterceptor) getClaimsFromVerifyToken(
	ctx context.Context,
	values []string,
) (context.Context, error) {
	token := values[0]
	claims, err := interceptor.verifyToken(token)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "access token is invalid: %v", err)
	}
	return interceptor.addClaimsToContext(ctx, ClaimsKey, claims)
}

func (interceptor *AuthInterceptor) addClaimsToContext(
	ctx context.Context,
	key,
	val interface{},
) (context.Context, error) {
	ctx = context.WithValue(ctx, key, val)
	return ctx, nil
}

func (interceptor *AuthInterceptor) verifyToken(token string) (*JWTAccessClaims, error) {
	// if token is static token return claims static token
	if claims, ok := staticToken[token]; ok {
		return &claims, nil
	}
	// if token is not static token, token will validate using jwt verify token
	return interceptor.getValidJwtClaims(token)
}

func (interceptor *AuthInterceptor) getValidJwtClaims(token string) (*JWTAccessClaims, error) {
	claims, err := interceptor.verifyJwt(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	return claims, nil
}

func (interceptor *AuthInterceptor) verifyJwt(token string) (*JWTAccessClaims, error) {
	jwtToken, err := interceptor.parseWithClaimsJwt(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	return interceptor.getClaimsJwt(jwtToken)
}

func (interceptor *AuthInterceptor) getClaimsJwt(jwtToken *jwt.Token) (*JWTAccessClaims, error) {
	claims, ok := jwtToken.Claims.(*JWTAccessClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func (interceptor *AuthInterceptor) parseWithClaimsJwt(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(
		token,
		&JWTAccessClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return interceptor.verifyJwtKey(token)
		},
	)
}

func (interceptor *AuthInterceptor) verifyJwtKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodRSA)
	if !ok {
		return nil, fmt.Errorf("unexpected token sign")
	}
	return interceptor.getVerifyKey([]byte(publicKey))
}

func (interceptor *AuthInterceptor) getVerifyKey(publicKey []byte) (interface{}, error) {
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}
