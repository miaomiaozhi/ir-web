package v1

type EngineResponse struct {
	Title []string `json:"Title"`
	Urls  []string `json:"Urls"`
	Time  int32    `json:"Time"`
}
