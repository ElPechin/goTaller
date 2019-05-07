package myml

import (
	"github.com/mercadolibre/academy19/myml/src/api/apierrors"
	"github.com/mercadolibre/academy19/myml/src/api/domain/myml"
	)



func GetUrl( usrId int64 ) (*myml.User, *apierrors.ApiError) {

	//creo una variable user de tipo myml.User estatica

	user := &myml.User{
		ID: usrId,
	}

	user.Get()
	user.SiteInfo, _ = GetSite(user.SiteID)


	return user, nil
	//devuelvo el contenido del puntero si uso la forma de declaracion "var user myml.User"
	// return &user
}

func GetSite (siteId string) (myml.Site, *apierrors.ApiError){

	site := myml.Site{
		ID: siteId,
	}

	site.GetSites()
	site.Categories, _ = site.GetCategory()

	return site, nil
}

/*func GetCategories (siteId string) (myml.Category, *apierrors.ApiError){

	site.GetCategory(siteId)

	return category, nil
}
*/
