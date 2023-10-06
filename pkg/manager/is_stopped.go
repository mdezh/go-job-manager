package manager

func (m *manager) isStopped() bool {
	select {
	case <-m.stop:
		return true
	default:
		return false
	}
}
