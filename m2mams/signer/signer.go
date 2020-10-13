package m2mams_signer

import (
	"github.com/dgrijalva/jwt-go"
	m2mamspkp "github.com/flsusp/m2mams-go-client/m2mams/pkp"
	"time"
)

type Signer struct {
	privateKeyProvider m2mamspkp.PrivateKeyProvider
	context            string
	keyPair            string
}

func (s Signer) generateSignedToken() (string, error) {
	key, err := s.privateKeyProvider.LoadKey(s.context, s.keyPair)
	if err != nil {
		return "", err
	}
	uid, err := s.privateKeyProvider.LoadKeyUid(s.context, s.keyPair)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"uid": uid,
		"kp":  s.keyPair,
		"ts":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
