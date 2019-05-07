package myml

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/academy19/myml/src/api/apierrors"
	"github.com/mercadolibre/academy19/myml/src/api/services/myml"
	"net/http"
	"strconv"
)

const (
	paramUserID = "userId"
)

func GetUserId(context *gin.Context) {


	userId := context.Param(paramUserID)
	id, err := strconv.ParseInt(userId, 10, 64)

	if err != nil{
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}

	 user, apiErr := myml.GetUrl(id)
	 if(apiErr != nil){

	 	context.JSON(apiErr.Status, apiErr)
	 	return
	 }
	//devolucion de la llamada exitosa del metodo myml.GetUrl(userId)
	context.JSON( http.StatusOK , user)
}