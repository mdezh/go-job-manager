package manager

func (m *manager) isStarted() bool {
	select {
	case <-m.started:
		return true
	default:
		return false
	}
}
