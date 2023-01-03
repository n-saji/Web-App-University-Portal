package service

func (s *Service) RunDailyMigrations() {

	go s.CheckOutDatedTokensSetFalse()
	go s.DeleteNotValidTokens()
}

func (s *Service) CheckOutDatedTokensSetFalse() {

}

func (s *Service) DeleteNotValidTokens() {}
