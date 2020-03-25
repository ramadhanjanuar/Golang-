package main

import (
	"time"

	"gopkg.in/oauth2.v3/manage"
)

// AuthorizeCodeTokenCfg for Mange config
var AuthorizeCodeTokenCfg = &manage.Config{
	AccessTokenExp:    time.Hour * 2,
	RefreshTokenExp:   time.Hour * 24 * 3,
	IsGenerateRefresh: false}

// TokenStore for token store file name
const TokenStore = "db_token"

// JwtKey for Jwt Key
const JwtKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCtsH9uCywIxzFtdx4LjDbq8ONLaaYzyPedk/Rd80HE84nRSTTj
P2jbSO7JEv8oES94ANU4Q/H+pGIioitLgmgR37KJXAQ371iRZxFdD3z/iMHzL8/S
lV9hpbqj4CBp4A0O0Nag2AomZR3FSPAe9WHirBmewTO5A3nNrC4SRjGN8QIDAQAB
AoGAWP13LMCYnR7B2l4PjMcYVCN7sWW9/AZZp+joaUJ7ThehYxNo/427ga7KeUc8
aCD1+zwiAqP9VwSOK0R/fj6gZrzXrkKwmiGqXAii/xibMfCTcWN7YokpO4K7f14A
cTJtHhR3Mm1U/aQHMbfe+98kbVioG83uVSFG8QQIAYePmAkCQQD5mnz8m6CfnOIs
r2ILdr3hr6s/QpNZkbzum0/lWx9CzmFTx11tHANl+1x/qK//uSP4XN5yrRHrY1j/
ZBjggrnjAkEAsiP6qrZVsmo7Rilhtu2M51TeoR5c4w+tf5yXGv3GxdFMosn0Ovq9
TKY3NOUGODb5rGRS9mQbk7/EqLHXADaxGwJBAKkc3b/85XgfQDKMZX5k5kPp1LnY
rqMKFhsICbrY4TZPRCwSMJ+DUl4fT02DCjPVyaQ1r9W8ox57wknPFPfQOocCQAUL
f83V3BdxarxR77J+h78FaEin03mauA8lICJjjX2Mr97sWT5SW4oQKwAUfFibNfbr
4G2Z6C4HAErGvpAC/IECQB2XnUuwBa/n2OFgTi/dQIMjFwGe9EQ9TU2EuiGvCoKD
kVHOsMBwFxtD4j4e3l8uocZMNIRcq/jr7Tpnm36ZEo8=
-----END RSA PRIVATE KEY-----`

// Salt for hashing password
const Salt = "rXI$L@bdND"

// ClientID for client_id
const ClientID = "000000"

// ClientSecret for client_secret
const ClientSecret = "999999"

// Domain for domain
const Domain = "https://auth.99group.co"
