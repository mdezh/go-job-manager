package manager

func (m *manager) Done() <-chan empty {
	return m.done
}
