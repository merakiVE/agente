package types

type MessageAgent struct {
	AgentID     string `json:"agent_id"`
	JobID       string `json:"job_id"`
	ProcedureID string `json:"procedure_id"`
	ActivityID  string `json:"activity_id"`
	Status      string `json:"status"`
}
