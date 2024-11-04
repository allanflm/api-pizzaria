package handler

import (
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": data.Pizzas,
	})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	c.JSON(201, newPizza)
}

func GetPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error()})
		return
	}
	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			c.JSON(200, gin.H{"message": "Pizza deleted"})
			return
		}
	}

	c.JSON(200, gin.H{"message": "Pizza not found"})

}
func UpdatePizzaById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	var updatePizza models.Pizza
	if err := c.ShouldBindBodyWithJSON(&updatePizza); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatePizza
			data.Pizzas[i].ID = id
			data.SavePizza()
			c.JSON(200, data.Pizzas[i])
			return
		}
	}

	c.JSON(404, gin.H{"method": "pizza not found"})
}
