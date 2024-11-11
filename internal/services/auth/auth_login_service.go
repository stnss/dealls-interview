package auth

import (
	"context"
	"github.com/stnss/dealls-interview/internal/presentation"
	"github.com/stnss/dealls-interview/pkg/jwtx"
	"github.com/stnss/dealls-interview/pkg/logger"
	"time"
)

func (s *service) Login(ctx context.Context, param *presentation.LoginRequest) (*presentation.LoginResponse, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("service.auth.login"),
		)
	)

	logger.InfoWithContext(ctx, "processing login...", lf...)

	pass, err := s.crypto.DecryptRSAWithBase64(s.cfg.RSAKey.PrivateKey, param.Password)
	if err != nil {
		logger.ErrorWithContext(ctx, err, lf...)
		return nil, err
	}
	param.Password = pass

	user, err := s.userRepo.GetUserByEmail(ctx, param.Email)
	if err != nil {
		logger.ErrorWithContext(ctx, err, lf...)
		return nil, err
	}

	if err = s.crypto.BcryptValidate(user.Password, param.Password); err != nil {
		logger.ErrorWithContext(ctx, err, lf...)
		return nil, err
	}

	accessToken, accessTokenExp := s.jwt.GenerateJWT(jwtx.UserClaim{
		Uid:  user.ID,
		Name: user.Name,
	}, s.cfg.JWT.AccessSecret, s.cfg.JWT.ExpiredTime)
	refreshToken, _ := s.jwt.GenerateJWT(jwtx.UserClaim{
		Uid: user.ID,
	}, s.cfg.JWT.RefreshSecret, s.cfg.JWT.ExpiredTime)

	return &presentation.LoginResponse{
		AccessToken:  accessToken,
		ExpiresIn:    int(accessTokenExp.Sub(time.Now()).Seconds()),
		RefreshToken: refreshToken,
	}, nil
}
