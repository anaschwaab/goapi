package main

import (
	"net/http" // Pacote para funcionalidades HTTP
	"github.com/gin-gonic/gin" // Pacote Gin para construir a API
)

// Estrutura que define uma tarefa com ID, Título e Status
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// Variável global que contém uma lista de tarefas predefinidas
var tasks = []Task{
	{ID: 1, Title: "Task One", Status: "pending"},
	{ID: 2, Title: "Task Two", Status: "completed"},
}

func main() {
	// Cria um roteador Gin com as configurações padrão
	r := gin.Default()

	// Define a rota para obter todas as tarefas
	r.GET("/tasks", func(c *gin.Context) {
		// Retorna a lista de tarefas como JSON com status HTTP 200 (OK)
		c.JSON(http.StatusOK, tasks)
	})

	// Define a rota para obter uma tarefa específica pelo ID
	r.GET("/tasks/:id", func(c *gin.Context) {
		// Pega o ID da tarefa da URL
		id := c.Param("id")
		// Procura a tarefa com o ID correspondente
		for _, task := range tasks {
			if string(task.ID) == id {
				// Se encontrada, retorna a tarefa como JSON com status 200 (OK)
				c.JSON(http.StatusOK, task)
				return
			}
		}
		// Se não encontrada, retorna uma mensagem de erro com status 404 (Not Found)
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})

	// Define a rota para criar uma nova tarefa
	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task
		// Vincula os dados JSON do corpo da requisição à estrutura Task
		if err := c.ShouldBindJSON(&newTask); err != nil {
			// Se houver erro na vinculação, retorna um erro com status 400 (Bad Request)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Define o ID da nova tarefa como o próximo na sequência
		newTask.ID = len(tasks) + 1
		// Adiciona a nova tarefa à lista de tarefas
		tasks = append(tasks, newTask)
		// Retorna a nova tarefa como JSON com status 201 (Created)
		c.JSON(http.StatusCreated, newTask)
	})

	// Define a rota para atualizar uma tarefa existente
	r.PUT("/tasks/:id", func(c *gin.Context) {
		// Pega o ID da tarefa da URL
		id := c.Param("id")
		var updatedTask Task
		// Vincula os dados JSON do corpo da requisição à estrutura Task
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			// Se houver erro na vinculação, retorna um erro com status 400 (Bad Request)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Procura a tarefa com o ID correspondente
		for i, task := range tasks {
			if string(task.ID) == id {
				// Atualiza o título e o status da tarefa
				tasks[i].Title = updatedTask.Title
				tasks[i].Status = updatedTask.Status
				// Retorna a tarefa atualizada como JSON com status 200 (OK)
				c.JSON(http.StatusOK, tasks[i])
				return
			}
		}
		// Se não encontrada, retorna uma mensagem de erro com status 404 (Not Found)
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})

	// Define a rota para excluir uma tarefa pelo ID
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		// Pega o ID da tarefa da URL
		id := c.Param("id")
		// Procura a tarefa com o ID correspondente
		for i, task := range tasks {
			if string(task.ID) == id {
				// Remove a tarefa da lista de tarefas
				tasks = append(tasks[:i], tasks[i+1:]...)
				// Retorna uma mensagem de confirmação com status 200 (OK)
				c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
				return
			}
		}
		// Se não encontrada, retorna uma mensagem de erro com status 404 (Not Found)
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})

	// Inicia o servidor na porta padrão (8080)
	r.Run()
}
