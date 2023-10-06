package manager

import "context"

func (m *manager) execJob(j *jobRecord) {
	j.setWorking(true)
	defer m.wg.Done()
	defer j.setWorking(false)

	ctx, cancel := context.WithTimeout(m.ctx, j.cfg.Timeout())
	defer cancel()

	if err := j.job(ctx); err != nil {
		m.logger.Println(j.name, ":", err)
	}
}
