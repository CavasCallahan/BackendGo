package services

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "cebolasLindasOk"

type TokenDetails struct {
	AcessToken string
	AcessUuid  string
	AtExpires  int64
}

type AcessDetails struct {
	AcessUuid string
	AuthId    string
}

func GenerateToken(auth_id string) (*TokenDetails, error) { //generates the acess_token
	token_details := &TokenDetails{}
	token_details.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	token_details.AcessUuid = "djwudhwudhwu"

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["acess_uuid"] = token_details.AcessUuid
	atClaims["auth_id"] = auth_id
	atClaims["exp"] = token_details.AtExpires
	token_unsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token_details.AcessToken, err = token_unsigned.SignedString([]byte(secretKey))

	if err != nil {
		return nil, err
	}

	return token_details, nil

}

func extractToken(request *http.Request) string { //Extrats the bearer token
	bearToken := request.Header.Get("Authorization")

	strSplit := strings.Split(bearToken, " ")

	if len(strSplit) == 2 {
		return strSplit[1]
	}

	return ""
}

func VerifyToken(request *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(request)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValid(request *http.Request) error {
	token, err := VerifyToken(request)

	fmt.Print(token)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return err
	}

	return nil
}

func ExtractTokenMetaData(request *http.Request) (*AcessDetails, error) {
	token, err := VerifyToken(request)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		acessUuid, ok := claims["acess_uuid"].(string)

		if !ok {
			return nil, err
		}

		auth_id, ok := claims["auth_id"].(string)

		if !ok {
			return nil, err
		}

		return &AcessDetails{
			AcessUuid: acessUuid,
			AuthId:    auth_id,
		}, nil
	}

	return nil, err
}
