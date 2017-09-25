package pool

import (
	"github.com/garyburd/redigo/redis"
)

type AgentManager struct {
	PubSubConn *redis.PubSubConn
}

func (this *AgentManager) SubscribeChannel(channel string) error {
	return this.PubSubConn.PSubscribe(channel)
}

func (this *AgentManager) Close() error {
	return this.PubSubConn.Close()
}

func (this *AgentManager) ReceiveMessage() interface{} {
	return this.PubSubConn.Receive()
}

func NewAgentManager(c redis.Conn) *AgentManager {
	return &AgentManager{
		PubSubConn: &redis.PubSubConn{Conn: c},
	}
}
