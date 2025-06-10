package models

type AgentResponse struct {
	Items []*AgentFile `json:"items"`
}

type AgentFile struct {
	FileName string `json:"file_name"`
	Code     string `json:"code"`
}
