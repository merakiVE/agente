package pool

import "github.com/garyburd/redigo/redis"

type ClientManager struct {
	pool *redis.Pool
}

func NewClientManager(connPool *redis.Pool) *ClientManager {
	return &ClientManager{pool: connPool}
}

func (this *ClientManager) SendProcedureRequest(procedure string, params map[string]interface{}) {
	conn := this.pool.Get()

	r := ProcedureRequest{procedure: procedure, params: params}
	data, _ := r.Serialize()
	conn.Do("PUBLISH", CHANNEL_PROCEDURE_REQUEST, data)
	defer conn.Close()
}
