package pool

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type ClientAgentManager struct {
	pool *redis.Pool
}

func NewClientAgentManager(connPool *redis.Pool) *ClientAgentManager {
	return &ClientAgentManager{pool: connPool}
}

func (this *ClientAgentManager) SendProcedureRequest(id_procedure string, params map[string]interface{}) {
	conn := this.pool.Get()

	r := ProcedureRequest{IDProcedure: id_procedure, Params: params}
	data, err := r.Serialize()

	if err != nil {
		log.Fatal("Error, serialize Procedure Request")
	}

	conn.Do("PUBLISH", CHANNEL_PROCEDURE_REQUEST, data)
	defer conn.Close()
}
