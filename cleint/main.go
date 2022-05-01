package main

import (
	"fmt"
	api "github.com/eminoz/grpc-api/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := api.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}
		req := &api.Request{
			A: int64(a),
			B: int64(b)}
		response, err := client.Add(context, req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
	})
	g.GET("/mult/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}
		req := &api.Request{
			A: int64(a),
			B: int64(b),
		}
		response, err := client.Multiply(context, req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		context.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
