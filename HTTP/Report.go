package HTTP

type DirSearchReport struct {
	TargetURL string
	Endpoints []Endpoint

	Priority  int
}

type Endpoint struct{
	Entity string
	StatusCode int
}

