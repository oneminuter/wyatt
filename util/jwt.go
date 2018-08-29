package util

import (
	"log"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	signingKey = "wyatt.oneminuter.com!@#$%^&*()_+"
)

type Token struct {
	UserId int64
	UUID   string
	Status int64
}

//生成token
func NewToken(userId, status int64, uuid string) string {
	claims := jwt.StandardClaims{
		Id:        strconv.FormatInt(userId, 10),
		Audience:  uuid,
		IssuedAt:  status,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := withClaims.SignedString([]byte(signingKey))
	if err != nil {
		log.Println(err)
		return ""
	}
	return s
}

//解析token
func ParseToken(tokenString string) (t *Token, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return generateToken(claims), nil
	} else {
		return new(Token), err
	}
}

//从解析token的结果构造token
func generateToken(claims *jwt.StandardClaims) *Token {
	userId, _ := strconv.ParseInt(claims.Id, 10, 64)
	return &Token{
		UserId: userId,
		UUID:   claims.Audience,
		Status: claims.IssuedAt,
	}
}
