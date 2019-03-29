package request

import (
	"net/http"
	"time"
)

// Call encapsulates all information about a request to and response from an API.
type Call struct {
	RequestMethod   string
	RequestURL      string
	RequestHeaders  http.Header
	RequestBody     string
	ResponseStatus  int
	ResponseHeaders http.Header
	ResponseBody    string
	Took            time.Duration
}
