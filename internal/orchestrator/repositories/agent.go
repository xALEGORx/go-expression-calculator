package repositories

import (
	"context"
	"time"

	"github.com/xALEGORx/go-expression-calculator/pkg/database"
)

type Agent struct {
}

type AgentModel struct {
	AgentID  string    `json:"agent_id"`
	LastPing time.Time `json:"last_ping"`
}

// Get all agents in database
func (a *Agent) GetAllAgents() ([]AgentModel, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT * FROM agents ORDER BY last_ping DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	agents := []AgentModel{}

	for rows.Next() {
		var agent AgentModel
		if err = rows.Scan(&agent.AgentID, &agent.LastPing); err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return agents, nil
}

// Create new row with agent
func (a *Agent) Create(agentId string) error {
	query := "INSERT INTO agents (agent_id, last_ping) VALUES ($1, $2)"
	if _, err := database.DB.Query(context.Background(), query, agentId, time.Now()); err != nil {
		return err
	}

	return nil
}

// Update last ping by agent id
func (a *Agent) SetLastPing(agentId string) error {
	query := "UPDATE agents SET last_ping = $1 WHERE agent_id = $2"
	if _, err := database.DB.Query(context.Background(), query, time.Now(), agentId); err != nil {
		return err
	}

	return nil
}

func AgentRepository() *Agent {
	return &Agent{}
}
