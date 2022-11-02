package service

import (
	"CollegeAdministration/daos"
)

type Service struct {
	daos *daos.AdminstrationCloud
}

func New(db *daos.AdminstrationCloud) *Service {
	return &Service{
		daos: db,
	}
}
