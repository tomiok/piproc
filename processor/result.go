package processor

type Result struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func createResult(port int, host string) Result {
	return Result{
		Port: port,
		Host: host,
	}
}
