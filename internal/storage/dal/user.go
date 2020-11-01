package dal

import (
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

func (dalInstance *DAL) SelectUserByEmail(email string) (user *models.UserProfile, errVal error) {
	var userProfile models.UserProfile
	err := dalInstance.tx.QueryRowContext(
		dalInstance.ctx,
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
			JOIN church c ON u.church_id = c.id
			WHERE u.active IS TRUE AND c.active IS TRUE AND u.email = $1`, email).Scan(
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
		&userProfile.ChurchUUID,
	)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (dalInstance *DAL) InsertUser(userIn *pkg.UserIn) error {
	_, err := dalInstance.tx.ExecContext(
		dalInstance.ctx,
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
		userIn.Church,
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

func (dalInstance *DAL) DeleteUserByPubId(userPubId *uuid.UUID) error {
	_, err := dalInstance.tx.ExecContext(
		dalInstance.ctx,
		`DELETE FROM app_user WHERE pub_id = $1`,
		userPubId,
	)
	return err
}
