package service

import (
	"log"
	"time"
)

func (s *Service) RunDailyMigrations() {

	go s.CheckOutDatedTokensSetFalse()
	go s.DeleteNotValidTokens()
}

func (s *Service) CheckOutDatedTokensSetFalse() {

	all_tokens, err := s.daos.GetAllTokens()
	if err != nil {
		log.Panic(err)
	}
	var time_now = time.Now()
	for _, token := range all_tokens {

		if token.ValidTill.Sub(time_now) < 0 {
			s.daos.SetTokenFalse(token.Token)
		}

	}
}

func (s *Service) DeleteNotValidTokens() {
	err := s.daos.RunMigrationsForRemovingOutDatedTokens()
	if err != nil {
		log.Panic(err)
	}
}
