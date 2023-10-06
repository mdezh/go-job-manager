package manager

func (j *jobRecord) setWorking(w bool) {
	j.mx.Lock()
	defer j.mx.Unlock()

	j.working = w
}
