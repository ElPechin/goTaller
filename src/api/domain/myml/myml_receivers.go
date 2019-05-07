package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/academy19/myml/src/api/apierrors"
	"io/ioutil"
	"net/http"
)
const(
	urlUsers = "https://api.mercadolibre.com/users/"
	urlSite = "https://api.mercadolibre.com/sites/"
)

func (user *User) Get()  *apierrors.ApiError{



	if user.ID == 0{
		return &apierrors.ApiError{
			Message: "userId is Empty",
			Status: http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%d", urlUsers, user.ID)
	response, err := http.Get(final)
	//forma estatica de declarar un user
	//var user myml.User

	//Meto el error que recibo en vez de esta comparacion
	//siempre pregunto el camino con error y corto la ejecucion en ese momento.
	if err != nil{
		println(&err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	jsonByte, err := ioutil.ReadAll(response.Body)

	if err != nil{
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(jsonByte), &user); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	//devuelvo el puntero
	return nil

}

func (site *Site) GetSites()  *apierrors.ApiError{



	if site.ID == ""{
		return &apierrors.ApiError{
			Message: "siteId is Empty",
			Status: http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s", urlSite, site.ID)
	response, err := http.Get(final)
	fmt.Print("\n"+final+"\n")


	//forma estatica de declarar un user
	//var user myml.User

	//Meto el error que recibo en vez de esta comparacion
	//siempre pregunto el camino con error y corto la ejecucion en ese momento.
	if err != nil{
		println(&err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	jsonByte, err := ioutil.ReadAll(response.Body)

	if err != nil{
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(jsonByte), &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	//devuelvo el puntero
	return nil

}

func (site *Site) GetCategory()  (Category,*apierrors.ApiError){



	if site.ID == ""{
		return nil,&apierrors.ApiError{
			Message: "siteId is Empty",
			Status: http.StatusBadRequest,
		}
	}

	final := fmt.Sprintf("%s%s/categories", urlSite, site.ID)
	response, err := http.Get(final)
	fmt.Print("\n"+final+"\n")


	//forma estatica de declarar un user
	//var user myml.User

	//Meto el error que recibo en vez de esta comparacion
	//siempre pregunto el camino con error y corto la ejecucion en ese momento.
	if err != nil{
		println(&err)
		return nil,&apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	jsonByte, err := ioutil.ReadAll(response.Body)

	if err != nil{
		return nil,&apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	//creo variable de tipo struc category, donde luego el unmarshal
	var category Category

	if err := json.Unmarshal([]byte(jsonByte), &category); err != nil {
		return nil,&apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	//devuelvo el puntero
	return category, nil

}