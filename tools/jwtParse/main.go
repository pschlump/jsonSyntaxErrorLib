package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/pschlump/ReadConfig"
	"github.com/pschlump/dbgo"
)

/*
3. Add in code that parses jwttoken (../tools/jwtParse/main)
	Read in ./cfg.json
	Token on CLI
	Spit out contents of signed token
*/

var Cfg = flag.String("cfg", "cfg.json", "config file for this call")
var DbFlagParam = flag.String("db_flag", "", "Additional Debug Flags")
var InputToken = flag.String("token", "", "Input token to parse")

type GlobalConfigData struct {
	JwtKey string `json:"jwt_key_password" default:"$ENV$QR_JWT_KEY"`
}

var DbOn map[string]bool = make(map[string]bool)
var gCfg GlobalConfigData

func main() {

	// CLI Args
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "qr_svr2 : Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Printf("Extra arguments are not supported [%s]\n", fns)
		os.Exit(1)
	}

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	// ------------------------------------------------------------------------------
	// Read in Configuration
	// ------------------------------------------------------------------------------
	err := ReadConfig.ReadFile(*Cfg, &gCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	// ------------------------------------------------------------------------------
	// Parse the JWT
	// ------------------------------------------------------------------------------
	s := ParseAuthToken(*InputToken)

	// ------------------------------------------------------------------------------
	// Output Results
	// ------------------------------------------------------------------------------
	fmt.Printf("%s\n", s)
}

type JwtClaims struct {
	AuthToken string `json:"auth_token"`
	jwt.StandardClaims
}

func ParseAuthToken(jwtToken string) (AuthToken string) {

	// jwtTok := val

	// Parse and Validate the JWT Berrer
	// Extract the auth_token
	// Validate in the d.b. that this is a valid auth_token
	// get the user_id  -- if have user_id, then ...
	// if valid - then reutrn it.

	// Initialize a new instance of `Claims`
	claims := &JwtClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		data, err := hex.DecodeString(gCfg.JwtKey)
		if err != nil {
			return []byte{}, err
		}
		return data, nil // return gCfg.JwtKey, nil
	})
	if err != nil {
		dbgo.Printf("X-Authentication - %(LF)\n")
		if err == jwt.ErrSignatureInvalid {
			dbgo.Printf("X-Authentication - %(LF)\n")
			return
		}
		dbgo.Printf("X-Authentication - %(LF)\n")
		return
	}
	if !tkn.Valid {
		dbgo.Printf("X-Authentication - %(LF)\n")
		return
	}

	AuthToken = claims.AuthToken
	// dbgo.Printf("X-Authentication - AuthToken ->%s<- %(LF)\n", AuthToken)
	dbgo.Printf("AuthToken ->%s<-\n", AuthToken)

	//	/*
	//		CREATE TABLE if not exists q_qr_auth_tokens (
	//			auth_token_id 	serial primary key not null,
	//			user_id 			int not null,
	//			token			 	uuid not null,
	//			expires 			timestamp not null
	//		);
	//	*/
	//	var v2 []*SQLIntType
	//	stmt := `select t1.user_id as "x"
	//			from q_qr_users as t1
	//				join q_qr_auth_tokens as t2 on ( t1.user_id = t2.user_id )
	//			where t2.token = $1
	//		      and ( t1.start_date < current_timestamp or t1.start_date is null )
	//		      and ( t1.end_date > current_timestamp or t1.end_date is null )
	//			  and t1.email_verified = 'y'
	//		      and t1.setup_complete_2fa = 'y'
	//			  and t2.expires > current_timestamp
	//		`
	//	err = pgxscan.Select(ctx, conn, &v2, stmt, AuthToken)
	//	if err != nil {
	//		LogSQLError(www, req, stmt, err, AuthToken)
	//		return
	//	}
	//	dbgo.Printf("X-Authentication - len(v2) = %d %(LF)\n", len(v2))
	//	if len(v2) > 0 {
	//		dbgo.Printf("X-Authentication - %(LF)\n")
	//		if v2[0].X == nil {
	//			dbgo.Printf("X-Authentication - %(LF)\n")
	//			UserId = 0
	//			AuthToken = ""
	//		} else {
	//			dbgo.Printf("X-Authentication - %(LF)\n")
	//			UserId = *v2[0].X
	//			jjwt, err := GetJWTAuthResponceWriterFromWWW(www, req)
	//			if err != nil {
	//				dbgo.Printf("X-Authentication - %(LF)\n")
	//				return
	//			}
	//			dbgo.Printf("X-Authentication - %(LF)\n")
	//			jjwt.StateVars["__is_logged_in__"] = "y"
	//			jjwt.StateVars["__user_id__"] = fmt.Sprintf("%d", UserId)
	//			jjwt.StateVars["__auth_token__"] = AuthToken
	//		}
	//	} else {
	//		dbgo.Printf("X-Authentication - %(LF)\n")
	//		UserId = 0
	//		AuthToken = ""
	//	}
	//	dbgo.Printf("X-Authentication - %(LF)\n")

	return
}
