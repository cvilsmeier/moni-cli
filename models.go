package monibot

// Watchdog holds data for a Watchdog.
type Watchdog struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	IntervalMillis int64  `json:"intervalMillis"`
}

// Machine holds data for a Machine.
type Machine struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// A MachineSample holds data for a machine resource usage sample.
type MachineSample struct {

	// Unix time millis since 1970-01-01T00:00:00Z, always UTC, never local time.
	Tstamp int64

	// Loadavg 1 minute.
	Load1 float64

	// Loadavg 5 minutes.
	Load5 float64

	// Loadavg 15 minutes.
	Load15 float64

	// CPU usage percent 0..100 since last sample.
	CpuPercent int

	// Memory usage percent 0..100.
	MemPercent int

	// Disk usage percent 0..100.
	DiskPercent int

	// Number of disk bytes read since last sample.
	DiskRead int64

	// Number of disk bytes written since last sample.
	DiskWrite int64

	// Number of network bytes received since last sample.
	NetRecv int64

	// Number of network bytes sent since last sample.
	NetSend int64
}

// Metric holds data for a Metric.
type Metric struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"` // 0=Counter, 1=Gauge, 2=Histogram
}
