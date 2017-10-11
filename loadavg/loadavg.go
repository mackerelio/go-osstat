package loadavg

// Get load average
func Get() (*Loadavg, error) {
	return get()
}

// Loadavg represents load average values
type Loadavg struct {
	Loadavg1, Loadavg5, Loadavg15 float64
}
