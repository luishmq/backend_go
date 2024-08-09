package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type PredictionRequest struct {
    Features []float64 `json:"features"`
}

type PredictionResponse struct {
    Prediction float64 `json:"prediction"`
}

func main() {
    r := gin.Default()

    r.POST("/predict", func(c *gin.Context) {
        var req PredictionRequest

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        prediction := 42.0

        response := PredictionResponse{
            Prediction: prediction,
        }

        c.JSON(http.StatusOK, response)
    })

    r.Run(":8080")
}