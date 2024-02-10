package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/services"
	"github.com/xALEGORx/go-expression-calculator/pkg/response"
	"strconv"
)

type Task struct {
	Route *gin.RouterGroup
}

type TaskCreateRequest struct {
	Expression string `json:"expression" binding:"required"`
}

// @Summary Get all tasks
// @Tags Worker
// @ID task-index
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=[]repositories.AgentModel}
// @Router /task [get]
func (p *Task) Index(ctx *gin.Context) {
	tasks, err := repositories.TaskRepository().GetAllTasks()
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Data(ctx, tasks)
}

// @Summary Create task
// @Tags Worker
// @ID task-store
// @Accept json
// @Produce json
// @Param input body TaskCreateRequest true "fields"
// @Success 200 {object} response.SuccessResponse{data=repositories.AgentModel}
// @Router /task [post]
func (p *Task) Store(ctx *gin.Context) {
	var request TaskCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.BadRequest(ctx, "невалидный expression")
		return
	}

	task, err := services.TaskService().Create(request.Expression)
	if err != nil {
		response.InternalServerError(ctx, err.Error())
		return
	}

	response.Data(ctx, task)
}

// @Summary Get all tasks
// @Tags Worker
// @ID task-show
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse{data=repositories.AgentModel}
// @Router /task/:id [get]
func (p *Task) Show(ctx *gin.Context) {
	taskIdUrl := ctx.Param("id")
	taskId, err := strconv.Atoi(taskIdUrl)
	if taskId == 0 || err != nil {
		response.BadRequest(ctx, "введите корректный task_id")
		return
	}

	tasks, err := repositories.TaskRepository().GetById(taskId)
	if err != nil {
		response.NotFound(ctx, err.Error())
		return
	}

	response.Data(ctx, tasks)
}
