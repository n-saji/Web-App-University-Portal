package service

import (
	"CollegeAdministration/daos"
)

type Service struct {
	daos *daos.AdministrationCloud
}

func New(db *daos.AdministrationCloud) *Service {
	return &Service{
		daos: db,
	}
}
