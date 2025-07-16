package services

import (
	"context"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/models"
	"time"

	"github.com/pkg/errors"
)

type RefreshTokenService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error) {
	response := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, "token", tokenClaim.Email, time.Now())
	if err != nil {
		return response, errors.Wrap(err, "Failed generate token")
	}

	err = s.UserRepository.UpdateRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return response, errors.Wrap(err, "failed to update token by refresh token")
	}
	response.Token = token
	return response, nil
}
