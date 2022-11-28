package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type todo struct {
	ID         string `json:"id"`
	Item       string `json:"item"`
	Completed  bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Write Curd API for golang", Completed: false},
	{ID: "2", Item:"Read book", Completed: false},
	{ID: "3", Item: "Record Videos", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusOK, newTodo)
}

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("golang-testing"),
		newrelic.ConfigLicense("78508c11275de05cf562c807fe4fa9fba142NRAL"),
        newrelic.ConfigAppLogForwardingEnabled(true),
	)

	fmt.Printf("%s",err, app);


	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")
}