package dal

import "github.com/jsfan/hello-neighbour/internal/storage/models"

func (dal *DAL) SelectUserByEmail(email string) (user *models.UserProfile, errVal error) {
	var userProfile models.UserProfile
	err := dal.tx.QueryRowContext(
		dal.ctx,
		`SELECT u.pub_id,
			u.email,
       		u.password_hash,
			u.first_name,
			u.last_name,
			u.gender,
			u.date_of_birth,
			u.role,
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
		&userProfile.ChurchUUID,
	)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}
