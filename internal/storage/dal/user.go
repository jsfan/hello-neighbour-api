package dal

import "github.com/jsfan/hello-neighbour/internal/storage/models"

func (dal *Dal) SelectUserByEmail(email string) (user *models.UserProfile, errVal error) {
	var userProfile models.UserProfile
	err := dal.tx.QueryRowContext(
		dal.ctx,
		"SELECT user.pub_id," +
			"user.email," +
			"user.first_name," +
			"user.last_name," +
			"user.gender," +
			"user.date_of_birth " +
			"FROM user " +
			"JOIN church ON user.church_id = church.id" +
			"WHERE active is TRUE and email = ?", email).Scan(
		&userProfile.PubId,
		&userProfile.Email,
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.Gender,
		&userProfile.DateOfBirth,
		&userProfile.ChurchUUID,
	)
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}
