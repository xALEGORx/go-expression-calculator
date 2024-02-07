package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/pkg/response"
)

type Agent struct {
	Route *gin.RouterGroup
}

// @Summary Get all agents
// @Tags Agent
// @ID agent-index
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]repositories.AgentModel}
// @Router /agent [get]
func (a *Agent) Index(ctx *gin.Context) {
	agents, err := repositories.AgentRepository().GetAllAgents()
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Data(ctx, agents)
}
