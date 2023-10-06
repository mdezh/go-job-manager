package manager

func (m *manager) toStarted() bool {
	m.mx.Lock()
	defer m.mx.Unlock()

	if m.started {
		return false
	}

	m.started = true
	return true
}
