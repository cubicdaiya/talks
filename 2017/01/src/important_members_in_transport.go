// MaxIdleConns controls the maximum number of idle (keep-alive)
// connections across all hosts. Zero means no limit.
MaxIdleConns int

// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
// (keep-alive) connections to keep per-host. If zero,
// DefaultMaxIdleConnsPerHost is used.
MaxIdleConnsPerHost int

// IdleConnTimeout is the maximum amount of time an idle
// (keep-alive) connection will remain idle before closing
// itself.
// Zero means no limit.
IdleConnTimeout time.Duration
