package auth

import (
	"github.com/go-money-transfer/internal/balance"
	"github.com/go-money-transfer/internal/errors"
	"github.com/go-money-transfer/internal/user"
	"github.com/go-money-transfer/pkg/hash"
	"github.com/go-money-transfer/pkg/jwt"
)

type AuthService struct {
	UsersRepository   user.Repository
	BalanceRepository balance.Repository
}

func (s *AuthService) AuthUser(login Login) (string, error) {
	user, err := s.UsersRepository.FindUserByUsername(login.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.ErrInvalidUsernameOrPassword
	}
	validUser, err := hash.ValidPassword(user.Password, login.Password)
	if err != nil {
		return "", err
	}
	if !validUser {
		return "", nil
	}

	balance, err := s.BalanceRepository.FindByUserId(user.ID)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.ID, balance.ID)
	if err != nil {
		return "", err
	}

	return token, err
}
