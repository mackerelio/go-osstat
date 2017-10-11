package loadavg

// Get load average
func Get() (Loadavg, error) {
	return get()
}

// Loadavg represents load average values
type Loadavg struct {
	Loadavg1  float64
	Loadavg5  float64
	Loadavg15 float64
}
