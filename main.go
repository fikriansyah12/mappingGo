package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Representasi data yang akan kita gunakan
type Todo struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Status bool   `json:"status"`
}

// Inisiasi data sementara
var todos = []Todo{
    {ID: "1", Title: "Belajar Golang", Status: false},
    {ID: "2", Title: "Belajar Gin", Status: false},
}

func main() {
    router := gin.Default()

    // Inisiasi route endpoint
    router.GET("/todos", getTodos)
    router.POST("/todos", createTodo)
    router.GET("/todos/:id", getTodoByID)
    router.PUT("/todos/:id", updateTodo)
    router.DELETE("/todos/:id", deleteTodo)

    // Jalankan server
    router.Run("localhost:8080")
}

// Handler untuk mendapatkan semua todo
func getTodos(c *gin.Context) {
    c.JSON(http.StatusOK, todos)
}

// Handler untuk membuat todo baru
func createTodo(c *gin.Context) {
    var newTodo Todo
    if err := c.BindJSON(&newTodo); err != nil {
        return
    }
    todos = append(todos, newTodo)
    c.JSON(http.StatusCreated, newTodo)
}

// Handler untuk mendapatkan todo berdasarkan ID
func getTodoByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range todos {
        if a.ID == id {
            c.JSON(http.StatusOK, a)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// Handler untuk mengupdate todo
func updateTodo(c *gin.Context) {
    id := c.Param("id")
    var updatedTodo Todo

    if err := c.BindJSON(&updatedTodo); err != nil {
        return
    }

    for i, a := range todos {
        if a.ID == id {
            todos[i] = updatedTodo
            c.JSON(http.StatusOK, updatedTodo)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// Handler untuk menghapus todo
func deleteTodo(c *gin.Context) {
    id := c.Param("id")

    for i, a := range todos {
        if a.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}