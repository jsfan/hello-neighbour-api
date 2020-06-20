package dal

import "github.com/jsfan/hello-neighbour/internal/storage/models"

func (dal *Dal) SelectUserByEmail(email string) (user *models.UserProfile, errVal error) {
	var userProfile models.UserProfile
	err := dal.tx.QueryRowContext(
		dal.ctx,
		"SELECT pub_id,"+
			"email,"+
			"first_name,"+
			"last_name,"+
			"gender,"+
			"date_of_birth "+
			"FROM user WHERE active is TRUE and email = ?", email).Scan(
		&userProfile.PubId,
		&userProfile.Email,
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.Gender,
		&userProfile.DateOfBirth,
	)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}
