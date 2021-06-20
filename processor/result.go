package processor

type Result struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Protocol string `json:"protocol"`
}

func createResult(port int, host, protocol string) Result {
	return Result{
		Port:     port,
		Host:     host,
		Protocol: protocol,
	}
}
