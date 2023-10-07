package manager

import "time"

func (m *manager) jobLoop(j *jobRecord) {
	defer m.wg.Done()

	ticker := time.NewTicker(j.cfg.Interval())
	defer ticker.Stop()

	m.wg.Add(1)
	go m.execJob(j)

	for {
		select {
		case <-ticker.C:
			if !j.isWorking() {
				m.wg.Add(1)
				go m.execJob(j)
			}
		case <-m.stopped:
			return
		}
	}
}
