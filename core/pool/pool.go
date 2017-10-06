package pool

import (
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
	IDProcedure string `json:"procedure"`
	Params      map[string]interface{} `json:"params"`
}

func (this *ProcedureRequest) Serialize() ([]byte, error) {
	return json.Marshal(this)
}