package dal

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jsfan/hello-neighbour-api/internal/storage/models"
	"github.com/jsfan/hello-neighbour-api/pkg"
)

func (dalInstance *DAL) InsertChurch(ctx context.Context, churchIn *pkg.ChurchIn) (church *models.ChurchProfile, errVal error) {
	var churchProfile models.ChurchProfile
	err := sq.
		Insert("church").
		Columns(
			"name",
			"description",
			"address",
			"website",
			"email",
			"phone",
			"group_size",
			"same_gender",
			"min_age",
			"member_basic_info_update",
			"active").
		Values(
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
		).
		Suffix(`RETURNING
			pub_id,
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
		`).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		QueryRowContext(ctx).
		Scan(
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

func (dalInstance *DAL) SelectChurchByEmail(ctx context.Context, email string) (church *models.ChurchProfile, errVal error) {
	var churchProfile models.ChurchProfile
	err := sq.Select(
		"pub_id",
		"name",
		"description",
		"address",
		"website",
		"email",
		"phone",
		"group_size",
		"same_gender",
		"min_age",
		"member_basic_info_update",
		"active",
	).
		From("church").
		Where(sq.Eq{"email": email}).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		QueryRowContext(ctx).
		Scan(
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

func (dalInstance *DAL) UpdateChurchActivationStatus(ctx context.Context, churchPubId *uuid.UUID, isActive bool) error {
	_, err := sq.
		Update("church").
		Set("active", isActive).
		Where("pub_id", churchPubId).
		PlaceholderFormat(sq.Dollar).
		RunWith(dalInstance.db()).
		ExecContext(ctx)
	return err
}
