package pool

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"github.com/merakiVE/agente/core/types"
	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/koinos/db"
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

func (this *AgentManager) findProcedure(id string, p *models.ProcedureModel) {
	db.Model(p, db.GetCurrentDatabase()).FindOne(p, "v.id == '"+id+"'")
}

func (this *AgentManager) listenMessages() {
	for {
		select {
		case <-this.stopChan:
			fmt.Println("Stop listen messages")
			return
		default:
			switch v := this.ReceiveMessage().(type) {

			case redis.PMessage:

				switch v.Channel {

				case CHANNEL_NEW_PROCEDURE_REQUEST:

					d := types.ProcedureRequest{}
					err := d.LoadData(v.Data)

					ps := models.ProcedureSessionModel{}

					//ps.ID = "1"
					ps.UserID = d.UserID
					ps.ProcedureID = d.ProcedureID
					ps.CurrentStage = 1

					a := db.Model(ps, db.GetCurrentDatabase()).Create(&ps)

					fmt.Println(a)

					if err != nil {
						log.Fatal("Error parsing data")
					}

					var pm models.ProcedureModel
					var nm models.NeuronModel

					db.Model(pm, db.GetCurrentDatabase()).FindOne(&pm, "v.id == '"+d.ProcedureID+"'")

					activity, err := pm.GetFirstActivity()

					if err != nil {
						log.Fatal(err.Error())
					}

					db.Model(nm, db.GetCurrentDatabase()).FindOne(&nm, "v.id == '"+activity.NeuronKey+"'")

					fmt.Println(pm, nm)

					//this.EnqueueJob("job_procesor_activity", map[string]interface{}{"session_id": ps.ID})

				case CHANNEL_NEW_PROCEDURE_UPDATE:
					break

				}

				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)

			case redis.Message:
				fmt.Printf("%s: message: %s %s\n", v.Channel, v.Data)
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
