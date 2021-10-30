package dal

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

func (dalInstance *DAL) SelectUserByEmail(ctx context.Context, email string) (user *models.UserProfile, errVal error) {
	var userProfile models.UserProfile
	var churchPubId sql.NullString

	err := dalInstance.db().QueryRowContext(
		ctx,
		`SELECT u.pub_id,
			u.email,
       		u.password_hash,
			u.first_name,
			u.last_name,
			u.gender,
			u.date_of_birth,
			u.role,
			u.description,
			u.active,
			c.pub_id
			FROM app_user u
			LEFT JOIN church c ON u.church_id = c.id
			WHERE u.email = $1 AND u.active IS TRUE`, email).Scan(
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

func (dalInstance *DAL) InsertUser(ctx context.Context, userIn *pkg.UserIn) error {
	church := &userIn.Church
	if *church == "" {
		church = nil
	}

	_, err := dalInstance.db().ExecContext(
		ctx,
		`INSERT INTO app_user (
			church_id,
			email,
			password_hash,
			first_name,
			last_name,
			gender,
			date_of_birth,
			description,
			role,
			active
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
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
	)
	return err
}

func (dalInstance *DAL) DeleteUserByPubId(ctx context.Context, userPubId *uuid.UUID) error {
	_, err := dalInstance.db().ExecContext(
		ctx,
		`DELETE FROM app_user WHERE pub_id = $1`,
		userPubId,
	)
	return err
}

func (dalInstance *DAL) MakeLeader(ctx context.Context, churchPubId *uuid.UUID, userPubId *uuid.UUID) error {
	_, err := dalInstance.db().ExecContext(
		ctx,
		`UPDATE app_user
		SET church_id = (
				SELECT id
				FROM church
				WHERE pub_id = $1
			),
			role = 'leader'
		WHERE pub_id = $2`,
		churchPubId,
		userPubId,
	)
	return err
}
