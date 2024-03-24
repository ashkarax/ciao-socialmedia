package jwttoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func TempTokenForOtpVerification(securityKey string, email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(securityKey))
	if err != nil {
		fmt.Println(err, "error at creating jwt token ")
	}
	return tokenString, err
}

func GenerateRefreshToken(secretKey string) (string, error) {

	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 3600000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error occured while creating token:", err)
		return "", err
	}

	return signedToken, nil

}

func GenerateAcessToken(securityKey string, id string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 300,
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(securityKey))
	if err != nil {
		fmt.Println(err, "Error creating acesss token ")
		return "", err
	}
	return tokenString, nil
}

func UnbindEmailFromClaim(tokenString string, tempVerificationKey string) (string, error) {

	secret := []byte(tempVerificationKey)
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		fmt.Println(err)
		return "", err
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	return email, nil
}

func VerifyRefreshToken(accesToken string, securityKey string) error {
	key := []byte(securityKey)
	_, err := jwt.Parse(accesToken, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("-------", err)
		return err
	}

	return nil
}

func VerifyAccessToken(token string, secretkey string) (string, error) {
	key := []byte(secretkey)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	// If parsing failed, check the specific error and handle accordingly
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// Token is malformed
				return "", errors.New("malformed token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				claims, ok := parsedToken.Claims.(jwt.MapClaims)
				if !ok {
					return "", errors.New("failed to extract claims")
				}

				id, ok := claims["id"].(string)
				if !ok {
					return "", errors.New("ID claim not found or not a string")
				}

				// Token is expired or not valid yet
				return id, errors.New("expired token")
			} else {
				// Other validation errors
				return "", errors.New("validation error")
			}
		} else {
			// Other parsing errors
			return "", err
		}
	}

	// If the token is valid, extract claims and return the ID
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to extract claims")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("ID claim not found or not a string")
	}

	return id, nil
}
