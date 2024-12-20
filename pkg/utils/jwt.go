package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("yijiansanlian")


type Claims struct {
    ID        uint   `json:"id"`
    UserName  string `json:"username"`
    Authority int    `json:"authority"`
    jwt.StandardClaims
}

func GenerateToken(id uint, userName string, authority int)(string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1*time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),   
			Issuer:    "Gin-Mall",       
			IssuedAt:  nowTime.Unix(),      
			Subject:   "user-token",        
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims!=nil{
		if claims, ok := tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			return claims, nil
		}
	}
	return nil, err 
}