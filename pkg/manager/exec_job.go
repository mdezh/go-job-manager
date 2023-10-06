package manager

import "context"

func (m *manager) execJob(j *jobRecord) {
	defer m.wg.Done()

	j.setWorking(true)
	defer j.setWorking(false)

	if m.isStopped() {
		return
	}

	ctx, cancel := context.WithTimeout(m.ctx, j.cfg.Timeout())
	defer cancel()

	if err := j.job(ctx); err != nil {
		m.logger.Println(j.name, err)
	}
}
