package manager

import "errors"

func (m *manager) Stop() error {
	if !m.isStarted() {
		return errors.New("failed to stop job manager: should be started")
	}

	if m.isStopped() {
		return errors.New("failed to stop job manager: already stopped")
	}

	m.stopOnce.Do(func() {
		m.logger.Println("job manager gracefully shutting down...")
		close(m.stop)
	})

	<-m.done

	return nil
}
