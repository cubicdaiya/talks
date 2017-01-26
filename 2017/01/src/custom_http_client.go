HttpClient := http.Client{
	Timeout: time.Duration(Config.Timeout) * time.Second,
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   Config.DialTimeout * time.Second,
			KeepAlive: Config.DialKeepAlive * time.Second,
		}).DialContext,
		MaxIdleConns:          Config.MaxIdleConns,
		MaxIdleConnsPerHost:   Config.MaxIdleConnsPerHost,
		DisableCompression:    Config.DisableCompression,
		IdleConnTimeout:       time.Duration(Config.IdleConnTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(Config.ProxyReadTimeout) * time.Second,
	},
}
