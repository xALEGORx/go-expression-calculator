package agents

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/pkg/config"
	"github.com/xALEGORx/go-expression-calculator/pkg/websocket"
	"time"
)

var agents = make(map[string]*repositories.AgentModel)

func HandlePing(message amqp.Delivery) {
	// handle ping from agent (create row or update last_ping)
	agent := string(message.Body)
	_, ok := agents[agent]

	if !ok {
		// it is new agent, create row
		if err := repositories.AgentRepository().Create(agent); err != nil {
			logrus.Fatalf("Failed create a new agent %s: %s", agent, err.Error())
			return
		}

		agents[agent] = &repositories.AgentModel{
			AgentID:  agent,
			LastPing: time.Now(),
			Status:   repositories.AGENT_CONNECTED,
		}
		sendToWebsocket(*agents[agent])

		logrus.Infof("Connected new agent #%s", agent)
	}

	if err := repositories.AgentRepository().SetLastPing(agent); err != nil {
		// error in update database
		logrus.Fatalf("Failed update a last ping for %s: %s", agent, err.Error())
		return
	}

	if ok {
		agents[agent].LastPing = time.Now()
	}

	logrus.Debugf("Update last ping for %s", agent)
}

func HandleTimeoutAgents() {
	// check every second which agents is disconnected
	ticker := time.NewTicker(time.Second)
	timeout := time.Duration(config.Get().AgentTimeout) * time.Second
	pingTime := time.Duration(config.Get().AgentPing) * time.Second
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for agentID, agent := range agents {
				if time.Now().Add(-timeout).After(agent.LastPing) {
					// agent is disconnected more than 10 minutes - delete from database
					agents[agentID].Status = repositories.AGENT_DELETED
					sendToWebsocket(*agents[agentID])

					delete(agents, agentID)
					if err := repositories.AgentRepository().Delete(agentID); err != nil {
						logrus.Errorf("Failed delete agent #%s: %s", agent, err.Error())
						continue
					}

					logrus.Infof("Agent #%s timeout (deletion)", agentID)
				}

				if time.Now().Add(-pingTime).After(agent.LastPing) {
					// agent is disconnected less than 10 minutes - set status disconnected

					if agent.Status != repositories.AGENT_CONNECTED {
						// already sent message about agent was disconnected
						continue
					}

					agents[agentID].Status = repositories.AGENT_DISCONNECTED
					sendToWebsocket(*agents[agentID])
					logrus.Infof("Agent #%s disconnected", agentID)
				} else if agent.Status != repositories.AGENT_CONNECTED {
					// agent has been reconnected

					agents[agentID].Status = repositories.AGENT_CONNECTED
					sendToWebsocket(*agents[agentID])
					logrus.Infof("Agent #%s has been reconnected", agentID)
				}
			}
		}
	}
}

func InitAgents() error {
	// loading current agents from database
	agentsDb, err := repositories.AgentRepository().GetAllAgents()
	if err != nil {
		return err
	}

	for _, agent := range agentsDb {
		agents[agent.AgentID] = &agent
	}

	return nil
}

func sendToWebsocket(agent repositories.AgentModel) {
	wsData := websocket.WSData{
		Action: "update_agent",
		Id:     agent.AgentID,
		Data:   agent,
	}
	if err := websocket.Broadcast(wsData); err != nil {
		logrus.Errorf("Failed send message to websocket about agent #%s: %s", agent.AgentID, err.Error())
		return
	}
}
