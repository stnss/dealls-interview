package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stnss/dealls-interview/internal/consts"
	"github.com/stnss/dealls-interview/internal/entity"
	"github.com/stnss/dealls-interview/pkg/databasex"
	"github.com/stnss/dealls-interview/pkg/helper"
	"github.com/stnss/dealls-interview/pkg/logger"
	"time"
)

type repository struct {
	db databasex.Adapter
}

func NewUserRepository(db databasex.Adapter) Repository {
	return &repository{db: db}
}

func (r *repository) GetUserById(ctx context.Context, userID string) (*entity.User, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("user_repository.get_user_by_id"),
			logger.String("user_id", userID),
		)
		user entity.User
	)

	logger.InfoWithContext(ctx, "getting user by id", lf...)

	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE id = :id LIMIT 1;`,
		consts.TableNameUser,
	)

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		logger.ErrorWithContext(ctx, err)
		return nil, err
	}

	err = stmt.GetContext(ctx, &user, map[string]any{
		"id": userID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		logger.WarnWithContext(ctx, "user not found", lf...)
		return nil, consts.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var (
		lf = logger.NewFields(
			logger.EventName("user_repository.get_user_by_email"),
			logger.String("email", email),
		)
		user entity.User
	)

	logger.InfoWithContext(ctx, "getting user by email", lf...)

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE email = :email LIMIT 1;",
		consts.TableNameUser,
	)

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		logger.ErrorWithContext(ctx, err)
		return nil, err
	}

	err = stmt.GetContext(ctx, &user, map[string]any{
		"email": email,
	})
	if errors.Is(err, sql.ErrNoRows) {
		logger.WarnWithContext(ctx, "user not found", lf...)
		return nil, consts.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, data *entity.User) error {
	var (
		lf = logger.NewFields(
			logger.EventName("user_repository.create_user"),
			logger.String("data.name", data.Name),
			logger.String("data.emai", data.Email),
		)
		err error
	)

	logger.InfoWithContext(ctx, "creating user", lf...)

	if data.ID == "" {
		id, _ := uuid.NewV7()
		data.ID = id.String()
	}
	data.CreatedAt = time.Now()

	query, val, err := helper.StructQueryInsert(data, consts.TableNameUser, "db", false)

	_, err = r.db.QueryX(
		ctx,
		query,
		val...,
	)
	if err != nil {
		logger.ErrorWithContext(ctx, err)
		return r.db.ParseSQLError(err)
	}

	return nil
}
