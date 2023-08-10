package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	//router start
	  router := gin.Default()

	router.GET("/persons/:id", getPersonById)
	
	router.Run("localhost:8080")
}

func getPersonById(c *gin.Context) {
    id := c.Param("id")
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//dto implementation
	var person Person

	err = conn.QueryRow(context.Background(), "select name, job from persons where id=$1", id).Scan(&person.Name, &person.Job)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
		return
	}
	person.Id = id

    c.IndentedJSON(http.StatusOK, person)
    return
}

