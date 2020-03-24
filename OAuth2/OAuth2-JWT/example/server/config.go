package main

import (
	"time"

	"gopkg.in/oauth2.v3/manage"
)

// TokenStore variable for token store file name
const TokenStore = "db_token"

// JwtKey variable for Jwt Key
const JwtKey = "00000000"

// AuthorizeCodeTokenCfg for Mange config
var AuthorizeCodeTokenCfg = &manage.Config{AccessTokenExp: time.Hour * 2, RefreshTokenExp: time.Hour * 24 * 3, IsGenerateRefresh: true}
