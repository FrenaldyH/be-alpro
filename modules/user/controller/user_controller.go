package controller

import (
	"net/http"
	"strconv"

	"github.com/Mobilizes/materi-be-alpro/modules/user/service"
	"github.com/Mobilizes/materi-be-alpro/modules/user/validation"
	"github.com/Mobilizes/materi-be-alpro/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
    service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
    return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    req, err := validation.ValidateCreateUser(c)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    user, err := ctrl.service.CreateUser(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
    users, err := ctrl.service.GetAllUser()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal fetch data users")
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Berhasil fetch data users", users)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
    idParam := c.Param("id")

    id, err := strconv.Atoi(idParam) 
    if err != nil || id <= 0 {
        // Development: ---
        utils.ErrorResponse(c, http.StatusBadRequest, "id tidak valid")
        return 
    }

    user, err := ctrl.service.GetUserByID(uint(id))
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "id tidak ditemukan")
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "berhasil fetch data user", user)
}
