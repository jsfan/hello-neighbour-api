package dal

import (
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

func (dalInstance *DAL) InsertChurch(churchIn *pkg.ChurchIn) error {
	_, err := dalInstance.tx.ExecContext(
		dalInstance.ctx,
		`INSERT INTO church (
			name, 
			description, 
			address, 
			website, 
			email, 
			phone, 
			group_size, 
			same_gender, 
			min_age, 
			member_basic_info_update, 
			active
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		churchIn.Name,
		churchIn.Description,
		churchIn.Address,
		churchIn.Website,
		churchIn.Email,
		churchIn.Phone,
		churchIn.GroupSize,
		churchIn.SameGender,
		churchIn.MinAge,
		churchIn.MemberBasicInfoUpdate,
		false,
	)
	return err
}

func (dalInstance *DAL) SelectChurchByEmail(email string) (church *models.ChurchProfile, errVal error) {
	var churchProfile models.ChurchProfile
	err := dalInstance.tx.QueryRowContext(
		dalInstance.ctx,
		`SELECT pub_id,
			name,
			description,
			address,
			website,
			email,
			phone,
			group_size,
			same_gender,
			min_age,
			member_basic_info_update,
			active
			FROM church
			WHERE email = $1`, email).Scan(
		&churchProfile.PubId,
		&churchProfile.Name,
		&churchProfile.Description,
		&churchProfile.Address,
		&churchProfile.Website,
		&churchProfile.Email,
		&churchProfile.Phone,
		&churchProfile.GroupSize,
		&churchProfile.SameGender,
		&churchProfile.MinAge,
		&churchProfile.MemberBasicInfoUpdate,
		&churchProfile.Active,
	)
	if err != nil {
		return nil, err
	}
	return &churchProfile, nil
}

func (dalInstance *DAL) UpdateChurchActivationStatus(churchPubId *uuid.UUID, isActive bool) error {
	_, err := dalInstance.tx.ExecContext(
		dalInstance.ctx,
		`UPDATE church SET active = $1 WHERE pub_id = $2`,
		isActive,
		churchPubId,
	)
	return err
}
