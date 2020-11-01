package dal

import (
	"github.com/jsfan/hello-neighbour/internal/storage/models"
	"github.com/jsfan/hello-neighbour/pkg"
)

func (dalInstance *DAL) InsertChurch(churchIn *pkg.ChurchIn) error {
	// TODO: Implement me
	return nil
}

func (dalInstance *DAL) SelectChurchByEmail(email string) (church *models.ChurchProfile, errVal error) {
	// TODO: Implement me
	return &models.ChurchProfile{}, nil
}
