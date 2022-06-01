package gometr

type HealthCheck struct {
	ServiceID string
	Status    MyType
}

type MyType string

const (
	PassStatus MyType = "pass"
	FailStatus MyType = "fail"
)

type GoMetrClient struct {
	url     string
	secunda int
}

func NewGoMetrClient(url string, sec int) GoMetrClient {
	return GoMetrClient{
		url,
		sec,
	}
}

func (g GoMetrClient) GetMetrics() string {
	return string(g.secunda)
}

func (g GoMetrClient) Ping() error {
	return nil
}

func (g GoMetrClient) GetID() string {
	return g.url
}

func (g GoMetrClient) Health() bool {
	return g.GetHealth().Status == PassStatus
}

func (g GoMetrClient) GetHealth() HealthCheck {

	if g.url[0]%2 == 0 {
		return HealthCheck{g.url, PassStatus}
	} else {
		return HealthCheck{g.url, FailStatus}
	}

}
