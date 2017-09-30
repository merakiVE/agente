package pool

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"

	"github.com/merakiVE/agente/core/types"

	"fmt"
)

type AgentManager struct {
	pubSubConn *redis.PubSubConn
	enqueuer   *work.Enqueuer
	started    bool
	stopChan   chan bool
	dataChan   chan []byte
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

func (this *AgentManager) Stop() {
	this.stopChan <- true
	this.pubSubConn.Close()
}

func (this *AgentManager) ReceiveMessage() interface{} {
	return this.pubSubConn.Receive()
}

func (this *AgentManager) Start() {

	if this.started {
		return
	}
	this.started = true

	go this.listenMessages()
	go this.administratorAgent()
}

func (this *AgentManager) listenMessages() {
	for {
		select {
		case <-this.stopChan:
			fmt.Println("Stop listen messages")
			return
		default:
			switch v := this.ReceiveMessage().(type) {
			case redis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			case redis.PMessage:
				fmt.Printf("%s: message: %s %s\n", v.Channel, v.Data, v.Pattern)

				this.dataChan <- v.Data

			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				panic(v)
			}
		}
	}
}

func (this *AgentManager) administratorAgent() {
	for {
		select {
		case data := <-this.dataChan:

			x := types.AgentMessage{}
			x.LoadData(data)

			fmt.Println("data received ", x)

		case <-this.stopChan:
			fmt.Println("Stop administrator")
			return
		}
	}
}

func NewAgentManager(cnn redis.Conn, cnn_pool *redis.Pool, namespace string) *AgentManager {
	return &AgentManager{
		pubSubConn: &redis.PubSubConn{Conn: cnn},
		enqueuer:   work.NewEnqueuer(namespace, cnn_pool),
		dataChan:   make(chan []byte),
		stopChan:   make(chan bool),
		started:    false,
	}
}
