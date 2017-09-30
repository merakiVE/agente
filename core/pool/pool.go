package pool

import (
	"fmt"
	"encoding/json"
)

const (
	JOB_FINISH    = "FINISH"
	JOB_PAUSED    = "PAUSED"
	JOB_MANUAL    = "MANUAL"
	JOB_AUTOMATIC = "JOB_AUTOMATIC"

	CHANNEL_PROCEDURE_REQUEST = "pool.agents.request_procedure"
)

type ProcedureRequest struct {
	procedure string `json:"procedure"`
	params    map[string]interface{} `json:"params"`
}

func (this *ProcedureRequest) Serialize() ([]byte, error) {
	return json.Marshal(this)
}

func CreateNameChannel(channel, group, action string) string {
	return fmt.Sprintf("%s.%s.%s", channel, group, action)
}
