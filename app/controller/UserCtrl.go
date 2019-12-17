package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vid/app/controller/exception"
	"vid/app/database"
	"vid/app/database/dao"
	"vid/app/middleware"
	"vid/app/model/dto"
	"vid/app/model/enum"
	"vid/app/model/vo"
)

type userCtrl struct{}

var UserCtrl = new(userCtrl)

// GET /user?page (Admin)
func (u *userCtrl) QueryAllUsers(c *gin.Context) {
	pageString := c.Query("page")
	page, err := strconv.Atoi(pageString)
	if err != nil || page == 0 {
		page = 1
	}
	users, count := dao.UserDao.QueryAll(page)
	c.JSON(http.StatusOK,
		dto.Result{}.Ok().SetPage(count, page, users))
}

// GET /user/:uid (Non-Auth)
func (u *userCtrl) QueryUser(c *gin.Context) {
	uidString := c.Param("uid")
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.Result{}.Error(http.StatusBadRequest).SetMessage(exception.RouteParamError.Error()))
		return
	}

	user := dao.UserDao.QueryByUid(uid)
	if user == nil {
		c.JSON(http.StatusNotFound,
			dto.Result{}.Error(http.StatusNotFound).SetMessage(exception.UserNotFoundError.Error()))
		return
	}

	isSelfOrAdmin := middleware.GetAuthUser(c) == nil || user.Authority == enum.AuthAdmin
	extraInfo, status := dao.UserDao.QueryUserExtraInfo(isSelfOrAdmin, user)
	if status == database.DbNotFound {
		c.JSON(http.StatusNotFound,
			dto.Result{}.Error(http.StatusNotFound).SetMessage(exception.UserNotFoundError.Error()))
		return
	}

	c.JSON(http.StatusOK,
		dto.Result{}.Ok().PutData("user", user).PutData("extra", extraInfo))
}

// PUT /user/update (Auth)
func (u *userCtrl) UpdateUser(c *gin.Context) {
	user := middleware.GetAuthUser(c)

	username := c.DefaultPostForm("username", user.Username)
	sex := enum.StringToSex(c.DefaultPostForm("sex", string(user.Sex)))
	profile := c.DefaultPostForm("profile", user.Profile)
	birthTime := vo.JsonDate{}.Parse(c.DefaultPostForm("birth_time", user.BirthTime.String()), user.BirthTime)
	phoneNumber := c.DefaultPostForm("phone_number", user.PhoneNumber)

	user.Username = username
	user.Sex = sex
	user.Profile = profile
	user.BirthTime = birthTime
	user.PhoneNumber = phoneNumber

	status := dao.UserDao.Update(user)
	if status == database.DbNotFound {
		c.JSON(http.StatusNotFound,
			dto.Result{}.Error(http.StatusNotFound).SetMessage(exception.UserNotFoundError.Error()))
		return
	} else if status == database.DbFailed {
		c.JSON(http.StatusInternalServerError,
			dto.Result{}.Error(http.StatusInternalServerError).SetMessage(exception.UserUpdateError.Error()))
		return
	}

	c.JSON(http.StatusOK,
		dto.Result{}.Ok().SetData(user))
}

// DELETE /user/delete (Auth)
func (u *userCtrl) DeleteUser(c *gin.Context) {
	user := middleware.GetAuthUser(c)
	user, status := dao.UserDao.Delete(user.Uid)
	if status == database.DbNotFound {
		c.JSON(http.StatusNotFound,
			dto.Result{}.Error(http.StatusNotFound).SetMessage(exception.UserNotFoundError.Error()))
		return
	} else if status == database.DbFailed {
		c.JSON(http.StatusInternalServerError,
			dto.Result{}.Error(http.StatusInternalServerError).SetMessage(exception.UserDeleteError.Error()))
		return
	}

	c.JSON(http.StatusOK,
		dto.Result{}.Ok().SetData(user))
}
