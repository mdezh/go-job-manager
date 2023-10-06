package manager

import (
	"errors"
)

func (m *manager) Start() error {
	for {
		if m.isStarted() {
			return errors.New("failed to start job manager: already started")
		}

		if len(m.jobs) == 0 {
			return errors.New("failed to start job manager: no jobs added")
		}

		if m.toStarted() {
			break
		}
	}

	go m.startOnce.Do(func() {
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
