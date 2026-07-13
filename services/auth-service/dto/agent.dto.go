package dto

type AgentDto struct {
	AgentName   string `json:"agent_name"`
	AgentIP     string `json:"agent_ip"`
	Description string `json:"description"`
}
