package manager

import "errors"

func (m *manager) AddJob(name string, j Job, cfg JobConfiguration) error {
	if m.started {
		return errors.New("failed to add job: job manager is already started")
	}

	m.jobs = append(m.jobs, &jobRecord{
		name: name,
		job:  j,
		cfg:  cfg,
	})

	return nil
}
