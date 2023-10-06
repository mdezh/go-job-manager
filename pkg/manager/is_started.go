package manager

func (m *manager) isStarted() bool {
	m.mx.RLock()
	defer m.mx.RUnlock()

	return m.started
}
