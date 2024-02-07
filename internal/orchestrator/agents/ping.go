package agents

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"time"
)

var agents = make(map[string]time.Time)

func InitAgents() error {
	// loading current agents from database
	agentsDb, err := repositories.AgentRepository().GetAllAgents()
	if err != nil {
		return err
	}

	for _, agent := range agentsDb {
		agents[agent.AgentID] = agent.LastPing
	}

	return nil
}

func HandlePing(message amqp.Delivery) {
	// handle ping from agent (create row or update last_ping)
	agent := string(message.Body)
	_, ok := agents[agent]

	if !ok {
		// it is new agent, create row
		if err := repositories.AgentRepository().Create(agent); err != nil {
			logrus.Fatalf("Failed create a new agent %s: %s", agent, err.Error())
		}
	}

	if err := repositories.AgentRepository().SetLastPing(agent); err != nil {
		// error in update database
		logrus.Fatalf("Failed update a last ping for %s: %s", agent, err.Error())
	}

	agents[agent] = time.Now()
	logrus.Debugf("Update last ping for %s", agent)
}
