package http

import "net/http"

// GeneratorResponse is the response payload for the Article data model.
// See NOTE above in ArticleRequest as well.
//
// In the GeneratorResponse object, first a Render() is called on itself,
// then the next field, and so on, all the way down the tree.
// Render is called in top-down order, like a http handler middleware chain.
type GeneratorResponse struct {
	SensorData string `json:"sensor_data"`

	CookieString string `json:"cookies"`
}

// NewGeneratorRequest is used to return a new request
func NewGeneratorRequest(sensorData, cookieString string) *GeneratorResponse {
	resp := &GeneratorResponse{SensorData: sensorData, CookieString: cookieString}

	return resp
}

func (rd *GeneratorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
