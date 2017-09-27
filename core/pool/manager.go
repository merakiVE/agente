package pool

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"fmt"
)

type AgentManager struct {
	pubSubConn *redis.PubSubConn
	enqueuer   *work.Enqueuer

	started bool
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

func (this *AgentManager) closeConnPubSub() error {
	return this.closeConnPubSub()
}

func (this *AgentManager) Stop() error {
	return this.pubSubConn.Close()
}

func (this *AgentManager) ReceiveMessage() interface{} {
	return this.pubSubConn.Receive()
}

func (this *AgentManager) Start() {

	if !this.started {
		return
	}

	this.started = true

	this.initLoop()
}

func (this *AgentManager) initLoop() {
	for {
		switch v := this.ReceiveMessage().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.PMessage:
			fmt.Printf("%s: message: %s %s\n", v.Channel, v.Data, v.Pattern)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			panic(v)
		}
	}
}

func NewAgentManager(cnn redis.Conn, cnn_pool *redis.Pool, namespace string) *AgentManager {
	return &AgentManager{
		pubSubConn: &redis.PubSubConn{Conn: cnn},
		enqueuer:   work.NewEnqueuer(namespace, cnn_pool),
	}
}
