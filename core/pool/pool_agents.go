package pool

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
)

type PoolAgents struct {
	context    PoolAgentsContext
	workerPool *work.WorkerPool
}

func NewPoolAgents(concurrency uint, namespace string, pool *redis.Pool) (*PoolAgents) {
	pa := &PoolAgents{}

	pa.context = PoolAgentsContext{}
	pa.workerPool = work.NewWorkerPool(pa.context, concurrency, namespace, pool)

	for key, value := range Middlewares {
		pa.registerJob(key, value)
	}

	return pa
}

func (this *PoolAgents) registerMiddleware(fn interface{}) {
	this.workerPool.Middleware(fn)
}

func (this *PoolAgents) registerJob(name_job string, fn interface{}) {
	this.workerPool.Job(name_job, fn)
}

func (this *PoolAgents) Start() {
	this.workerPool.Start()
}

func (this *PoolAgents) Stop() {
	this.workerPool.Stop()
}
