package manager

import (
	"errors"
)

func (m *manager) Start() error {
	if len(m.jobs) == 0 {
		return errors.New("failed to start job manager: no jobs added")
	}

	if m.isStarted() {
		return errors.New("failed to start job manager: already started")
	}

	go m.startOnce.Do(func() {
		close(m.started)

		m.wg.Add(len(m.jobs))
		for _, j := range m.jobs {
			go m.jobLoop(j)
		}

		m.wg.Add(1)
		go m.catchCancel()

		m.wg.Wait()

		m.logger.Println("job manager gracefully done")

		close(m.done)
	})

	return nil
}
