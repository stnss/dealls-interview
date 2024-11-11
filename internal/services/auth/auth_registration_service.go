package auth

import (
	"context"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stnss/dealls-interview/internal/entity"
	"github.com/stnss/dealls-interview/internal/presentation"
	"github.com/stnss/dealls-interview/pkg/logger"
	"github.com/stnss/dealls-interview/pkg/util"
)

func (s *service) Registration(ctx context.Context, param *presentation.RegistrationRequest) error {
	var (
		lf = logger.NewFields(
			logger.EventName("service.user.store"),
		)

		userdata entity.User
	)

	logger.InfoWithContext(ctx, "storing user...", lf...)

	password, err := s.crypto.DecryptRSAWithBase64(s.cfg.RSAKey.PrivateKey, param.Password)
	if err != nil {
		return err
	}

	passwordConf, err := s.crypto.DecryptRSAWithBase64(s.cfg.RSAKey.PrivateKey, param.PasswordConfirmation)
	if err != nil {
		return err
	}

	if password != passwordConf {
		logger.WarnWithContext(ctx, "password not match with confirm password", lf...)
		return consts.ErrPasswordNotMatch
	}

	_ = util.CopyStruct(*param, &userdata, "json")

	userdata.Password, err = s.crypto.BcryptHash(password)
	if err != nil {
		logger.ErrorWithContext(ctx, err, lf...)
		return err
	}

	if err := s.userRepo.CreateUser(ctx, &userdata); err != nil {
		logger.ErrorWithContext(ctx, err, lf...)
		return err
	}

	return nil
}
