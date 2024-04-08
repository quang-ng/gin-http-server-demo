package controller

import (
	"net/http"
	"strconv"
	"todolist/api/service"
	"todolist/models"
	"todolist/util"

	"github.com/gin-gonic/gin"
)

// TodoController -> TodoController
type TodoController struct {
	service service.TodoService
}

// NewTodoController : NewTodoController
func NewTodoController(s service.TodoService) TodoController {
	return TodoController{
		service: s,
	}
}

// GetTodos : GetTodos controller
func (p TodoController) GetTodos(ctx *gin.Context) {
	var todos models.Todo

	keyword := ctx.Query("keyword")

	data, total, err := p.service.FindAll(todos, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Todo result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddTodo : AddTodo controller
func (p *TodoController) AddTodo(ctx *gin.Context) {
	var todo models.Todo
	ctx.ShouldBindJSON(&todo)

	if todo.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if todo.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}
	err := p.service.Save(todo)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create post")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Post")
}

// GetTodo : get Todo by id
func (p *TodoController) GetTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	var todo models.Todo
	todo.ID = id
	foundTodo, err := p.service.Find(todo)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Error Finding Post")
		return
	}
	response := foundTodo.ResponseMap()

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Post",
		Data:    &response})

}

func (p *TodoController) DeleteTodo(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.Delete(id)

	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to delete Post")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	c.JSON(http.StatusOK, response)
}

// UpdateTodo : get update by id
func (p TodoController) UpdateTodo(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var todo models.Todo
	todo.ID = id

	todoRecord, err := p.service.Find(todo)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "todo with given id not found")
		return
	}
	ctx.ShouldBindJSON(&todoRecord)

	if todoRecord.Title == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
		return
	}
	if todoRecord.Description == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
		return
	}

	if err := p.service.Update(todoRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store todo")
		return
	}
	response := todoRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated todo",
		Data:    response,
	})
}
