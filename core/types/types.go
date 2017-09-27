package types

import "encoding/json"

type AgentMessage struct {
	AgentID     string `json:"agent_id"`
	JobID       string `json:"job_id"`
	ProcedureID string `json:"procedure_id"`
	ActivityID  string `json:"activity_id"`
	Status      string `json:"status"`
}

func (this *AgentMessage) Serialize() ([]byte, error) {
	return json.Marshal(this)
}
