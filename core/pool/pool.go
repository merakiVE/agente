package pool

const (
	JOB_FINISH    = "FINISH"
	JOB_PAUSED    = "PAUSED"
	JOB_MANUAL    = "MANUAL"
	JOB_AUTOMATIC = "AUTOMATIC"

	POOL               = "pool"
	POOL_AGENTS        = "agents"
	ACTION_NEW_REQUEST = "new_request_procedure"
	ACTION_NEW_UPDATE  = "new_update"

	CHANNEL_NEW_PROCEDURE_REQUEST = POOL + "." + POOL_AGENTS + "." + ACTION_NEW_REQUEST
	CHANNEL_NEW_PROCEDURE_UPDATE  = POOL + "." + POOL_AGENTS + "." + ACTION_NEW_UPDATE
)
