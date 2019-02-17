package httpserver

type errorOutput struct {
	Message string `json:"message"`
}

func errorOutputFrom(err error) (string, interface{}) {
	return "error", errorOutput{Message: err.Error()}
}
