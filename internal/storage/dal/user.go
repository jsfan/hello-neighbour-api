package dal

import (
	"context"
	"database/sql"

	"github.com/jsfan/hello-neighbour-api/internal/rest/model"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
)

func (dalInstance *DAL) SelectUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	var userProfile models.UserProfile
	var churchPubId sql.NullString

	err := sq.Select(
		"u.pub_id",
		"u.email",
		"u.password_hash",
		"u.first_name",
		"u.last_name",
		"u.gender",
		"u.date_of_birth",
		"u.role",
		"u.description",
		"u.active",
		"c.pub_id",
	).
		From("app_user u").
		LeftJoin("church c ON u.church_id = c.id").
		Where(sq.And{sq.Eq{"u.email": email}, sq.Eq{"u.active": true}}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		QueryRowContext(ctx).
		Scan(
			&userProfile.PubId,
			&userProfile.Email,
			&userProfile.PasswordHash,
			&userProfile.FirstName,
			&userProfile.LastName,
			&userProfile.Gender,
			&userProfile.DateOfBirth,
			&userProfile.Role,
			&userProfile.Description,
			&userProfile.Active,
			&churchPubId,
		)

	if err != nil {
		return nil, err
	}
	if churchPubId.Valid {
		userProfile.ChurchUUID = churchPubId.String
	}
	return &userProfile, nil
}

func (dalInstance *DAL) InsertUser(ctx context.Context, userIn *model.UserIn) error {
	church := &userIn.Church
	if *church == "" {
		church = nil
	}

	_, err := sq.Insert("app_user").Columns(
		"church_id",
		"email",
		"password_hash",
		"first_name",
		"last_name",
		"gender",
		"date_of_birth",
		"description",
		"role",
		"active",
	).
		Values(
			church,
			userIn.Email,
			userIn.Password,
			userIn.FirstName,
			userIn.LastName,
			userIn.Gender,
			userIn.DateOfBirth,
			userIn.Description,
			userIn.Role,
			true,
		).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		ExecContext(ctx)
	return err
}

func (dalInstance *DAL) DeleteUserByPubId(ctx context.Context, userPubId *uuid.UUID) error {
	_, err := sq.
		Delete("app_user").
		Where(sq.Eq{"pub_id": userPubId}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		ExecContext(ctx)
	return err
}

func (dalInstance *DAL) MakeLeader(ctx context.Context, churchPubId *uuid.UUID, userPubId *uuid.UUID) error {
	_, err := sq.
		Update("app_user").
		Set("church_id", sq.Select("id").
			From("church").
			Where(sq.Eq{"pub_id": churchPubId})).
		Where(sq.Eq{"pub_id": userPubId}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		ExecContext(ctx)
	return err
}
