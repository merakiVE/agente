package types

import "encoding/json"

/*

 */
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

func (this *AgentMessage) LoadData(data_encode []byte) (error) {
	return json.Unmarshal(data_encode, this)
}

/*


 */
type ProcedureRequest struct {
	ProcedureID string `json:"procedure_id"`
	UserID      string `json:"user_id"`
	Params      map[string]interface{} `json:"params"`
}

func (this *ProcedureRequest) Serialize() ([]byte, error) {
	return json.Marshal(this)
}

func (this *ProcedureRequest) LoadData(data_encode []byte) (error) {
	return json.Unmarshal(data_encode, this)
}