package manager

func (m *manager) catchCancel() {
	defer m.wg.Done()

	select {
	case <-m.ctx.Done():
		m.stopOnce.Do(func() {
			m.logger.Println("job manager received cancel, gracefully shutting down...")
			close(m.stop)
		})
	case <-m.stop:
	}
}
