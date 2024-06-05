package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	ParamsGenerateJWT struct {
		ExpiredInMinute int
		SecretKey       string
		UserID          int64
	}

	ResultGenerateJWT struct {
		Token  string
		Expire int64
	}

	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}

	Claims struct {
		jwt.StandardClaims
		UserID int64 `json:"user_id,omitempty"`
	}
)

func GenerateJWT(p *ParamsGenerateJWT) (ResultGenerateJWT, error) {
	expiredAt := time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix()
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
		UserID: p.UserID,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(p.SecretKey))

	return ResultGenerateJWT{
		signedToken,
		expiredAt,
	}, err
}

func ValidateJWT(p *ParamsValidateJWT) (jwt.MapClaims, error) {
	return jwt.MapClaims{}, nil
}
