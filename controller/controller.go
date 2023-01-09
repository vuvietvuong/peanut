package controller

import (
	"github.com/google/uuid"
	"path/filepath"
	"peanut/pkg/apierrors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func bindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bindQueryParams(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindQuery(obj)

	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}
func bindForm(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBind(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}
func checkError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
	return true
}
func saveUploadedFile(ctx *gin.Context, name string, dir string) (path string, err error) {
	// Save uploaded file
	file, _ := ctx.FormFile(name)
	dst := uuid.New().String() + filepath.Ext(file.Filename)
	contentPath := dir + dst
	err = ctx.SaveUploadedFile(file, contentPath)
	if err != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, err.Error())
		return
	}

	return
}
