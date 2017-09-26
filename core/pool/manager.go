package pool

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
)

type AgentManager struct {
	pubSubConn *redis.PubSubConn
	enqueuer   *work.Enqueuer
}

func (this *AgentManager) GetEnqueuer() *work.Enqueuer {
	return this.enqueuer
}

func (this *AgentManager) EnqueueJob(name_job string, params map[string]interface{}) (*work.Job, error) {
	return this.enqueuer.Enqueue(name_job, params)
}

func (this *AgentManager) SubscribeChannel(channel string) error {
	return this.pubSubConn.PSubscribe(channel)
}

func (this *AgentManager) Close() error {
	return this.pubSubConn.Close()
}

func (this *AgentManager) ReceiveMessage() interface{} {
	return this.pubSubConn.Receive()
}

func NewAgentManager(cnn_pub_sub *redis.Conn, cnn_pool *redis.Pool, namespace string) *AgentManager {
	return &AgentManager{
		pubSubConn: &redis.PubSubConn{Conn: cnn_pub_sub},
		enqueuer:   work.NewEnqueuer(namespace, cnn_pool),
	}
}
