package manager

func (m *manager) isStopped() bool {
	select {
	case <-m.stopped:
		return true
	default:
		return false
	}
}
