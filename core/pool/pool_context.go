package pool

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"fmt"
)

var Middlewares map[string]interface{}

func middlewareLog(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: " + job.Name + " with id: " + job.ID)
	return next()
}

func init() {
	Middlewares = map[string]interface{}{
		"middleware_log": middlewareLog,
	}
}

type PoolAgentsContext struct {
	PoolManager *redis.Pool
}
