package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string // 用于给加密算法签名的字符串,使用的是对称加密，所以只需要一个密钥
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < 32 {
		return nil, fmt.Errorf("invalid key size: must be at least 32 characters")
	}
	return &JWTMaker{secretKey}, nil
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(username, duration)
	claims := NewMyClaims(payload)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(m.secretKey))
	return token, payload, err
}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验token的签名算法是否是HS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	claims, ok := jwtToken.Claims.(*MyClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claims.Payload, nil

}

type MyClaims struct {
	Payload *Payload
	jwt.RegisteredClaims
}

func NewMyClaims(payload *Payload) *MyClaims {
	return &MyClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(payload.ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(payload.CreatedAt),
			ID:        payload.ID.String(),
		},
	}
}
