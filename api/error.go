package api

type Error struct {
	Code    string `json:"code,omitempty"`
	Status  int32  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
