package pool

import (
	"log"
	"github.com/merakiVE/agente/core/types"
	"github.com/garyburd/redigo/redis"
)

type ClientAgentManager struct {
	pool *redis.Pool
}

func NewClientAgentManager(connPool *redis.Pool) *ClientAgentManager {
	return &ClientAgentManager{pool: connPool}
}

func (this *ClientAgentManager) SendProcedureRequest(user_id, procedure_id string, params map[string]interface{}) {

	r := types.ProcedureRequest{UserID: user_id, ProcedureID: procedure_id, Params: params}
	data, err := r.Serialize()

	if err != nil {
		log.Fatal("Error, Serialize Procedure Request")
	}

	conn := this.pool.Get()
	conn.Do("PUBLISH", CHANNEL_NEW_PROCEDURE_REQUEST, data)
	defer conn.Close()
}
