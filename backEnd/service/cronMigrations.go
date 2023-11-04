package service

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func (s *Service) RunDailyMigrations() {
	//making go routines synchronous
	wg.Add(2)
	ch := make(chan int)
	log.Println("running migrations")
	go s.CheckOutDatedTokensSetFalse(ch)
	go s.DeleteNotValidTokens(ch)
	wg.Wait()
}

func (s *Service) CheckOutDatedTokensSetFalse(ch chan int) {

	ch <- 1
	all_tokens, err := s.daos.GetAllTokens()
	if err != nil {
		log.Panic(err)
	}
	var time_now = time.Now()
	for _, token := range all_tokens {

		if token.ValidTill-time_now.Unix() < 0 {
			s.daos.SetTokenFalse(token.Token)
		}

	}
	defer wg.Done()
}

func (s *Service) DeleteNotValidTokens(ch chan int) {

	<-ch
	err := s.daos.RunMigrationsForRemovingOutDatedTokens()
	if err != nil {
		log.Panic(err)
	}
	defer wg.Done()
}
