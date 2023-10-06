package manager

func (j *jobRecord) isWorking() bool {
	j.mx.RLock()
	defer j.mx.RUnlock()

	return j.working
}
