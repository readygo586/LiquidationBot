package scanner

func (s *Scanner) LiquidationLoop() {
	defer s.wg.Done()

	for {
		select {
		case <-s.quitCh:
			return

		case pending := <-s.liquidationCh:
			logger.Printf("receive priority liquidation:%v\n", pending)
			s.liquidate(pending)
		}
	}
}

func (s *Scanner) liquidate(liquidation *Liquidation) error {
	//TODO(keep)
	return nil
}
