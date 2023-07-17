package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	pasteo       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		pasteo:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (m *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(username, duration)
	token, err := m.pasteo.Encrypt(m.symmetricKey, payload, nil)
	return token, payload, err
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := m.pasteo.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	// 校验token是否过期
	if time.Now().After(payload.ExpiredAt) {
		return nil, ErrExpiredToken
	}
	return payload, nil
}
