package Controller

import (
	"github.com/gin-gonic/gin"
	"gorm-tutorial/Model"
	"gorm-tutorial/Service"
	"gorm-tutorial/Token"
	"net/http"
	"strconv"
)

type Controller struct {
	userService *Service.UserService
}

func NewUserController(userService *Service.UserService) *Controller {
	return &Controller{
		userService: userService,
	}
}

// AddUserHandler thêm một người dùng mới từ dữ liệu yêu cầu.
func (uc *Controller) AddUserHandler(c *gin.Context) {
	var user Model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.userService.AddUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// GetUserByIDHandler trả về thông tin người dùng theo ID.
func (uc *Controller) GetUserByIDHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	user, err := uc.userService.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// UpdateUserHandler cập nhật thông tin người dùng từ dữ liệu yêu cầu.
func (uc *Controller) UpdateUserHandler(c *gin.Context) {
	var user Model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.userService.Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// SearchUserHandler tìm kiếm người dùng theo tên.
func (uc *Controller) SearchUserHandler(c *gin.Context) {
	keyword := c.PostForm("keyword")
	user, err := uc.userService.SearchUser(keyword)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uc *Controller) GetAllUser(c *gin.Context) {
	users, err := uc.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Users not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (uc *Controller) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	_, err := uc.userService.Login(email, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := Token.GenerateToken(email)
	refresh, err := Token.GenerateRefreshToken(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"refresh": refresh,
	})

}
