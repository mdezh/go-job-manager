package manager

import "time"

func (m *manager) jobLoop(j *jobRecord) {
	defer m.wg.Done()

	timer := time.NewTimer(j.cfg.Interval())

	m.wg.Add(1)
	go m.execJob(j)

	for {
		select {
		case <-timer.C:
			// recalculate interval on the each tick
			timer = time.NewTimer(j.cfg.Interval())

			if !j.isWorking() {
				m.wg.Add(1)
				go m.execJob(j)
			}
		case <-m.stopped:
			timer.Stop()
			return
		}
	}
}
