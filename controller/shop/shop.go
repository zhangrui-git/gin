package shop

import (
	"by/user/bean/shop"
	"by/user/model"
	"by/user/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(ctx *gin.Context) {
	var err error
	var s model.Shop
	err = ctx.ShouldBindJSON(&s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
		return
	}

	if s.Create().Error != nil {
		ctx.JSON(http.StatusOK, response.Failed(400000))
	} else {
		ctx.JSON(http.StatusOK, response.Success(s))
	}
}

func Del(ctx *gin.Context) {
	var err error
	var id shop.Pk
	err = ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
		return
	}

	var s = model.Shop{ShopId: id.Val}
	if s.Delete().Error != nil {
		ctx.JSON(http.StatusOK, response.Failed(400000))
	} else {
		ctx.JSON(http.StatusOK, response.Success(struct{}{}))
	}
}

func Save(ctx *gin.Context) {
	var err error
	var id shop.Pk
	var s model.Shop
	err = ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
		return
	}
	err = ctx.ShouldBindJSON(&s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
		return
	}
	s.ShopId = id.Val

	var fields = []string{"ShopName"}
	if s.Update(fields).Error != nil {
		ctx.JSON(http.StatusOK, response.Failed(400000))
	} else {
		ctx.JSON(http.StatusOK, response.Success(s))
	}
}

// GetById 通过ID获取单个店铺信息
func GetById(ctx *gin.Context) {
	var sid shop.Pk
	err := ctx.ShouldBindUri(&sid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
		return
	}

	var s model.Shop
	if s.GetBySid(sid).Error != nil {
		ctx.JSON(http.StatusOK, response.Failed(400000))
	} else {
		ctx.JSON(http.StatusOK, response.Success(s))
	}
}

// GetListById 通过ID获取多个店铺信息
func GetListById(ctx *gin.Context) {
	var err error
	var sidSet shop.PkSet
	err = ctx.ShouldBindJSON(&sidSet)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
		return
	}

	//var page common.Page
	//err = ctx.ShouldBindQuery(&page)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, response.New(400001, err.Error(), struct{}{}))
	//}

	var ss model.Shops
	if ss.GetListBySid(sidSet).Error != nil {
		ctx.JSON(http.StatusOK, response.Failed(400000))
	} else {
		ctx.JSON(http.StatusOK, response.Success(ss))
	}
}
