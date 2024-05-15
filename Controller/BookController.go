package Controller

import (
	"github.com/gin-gonic/gin"
	"gorm-tutorial/Model"
	"gorm-tutorial/Service"
	"net/http"
)

type BookController struct {
	Bsc *Service.BookService
}

func NewBookController(bookService *Service.BookService) *BookController {
	return &BookController{Bsc: bookService}
}

func (bc *BookController) AddBook(c *gin.Context) {
	var book Model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bc.Bsc.AddBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
func (bc *BookController) UpdateBook(c *gin.Context) {
	var book Model.Book
	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := bc.Bsc.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func (bc *BookController) SreachBook(c *gin.Context) {
	var book []Model.Book
	keyword := c.PostForm("keyword")
	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err, _ := bc.Bsc.SreachBook(keyword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}
