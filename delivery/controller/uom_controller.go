package controller

import (
	"github.com/eulbyvan/go-enigma-laundry/model"
	"github.com/eulbyvan/go-enigma-laundry/usecase"
	"github.com/eulbyvan/go-enigma-laundry/utils"
	"github.com/gin-gonic/gin"
)

type UomController struct {
	uomUseCase usecase.UomUseCase
}

func (pc *UomController) GetAllUom(ctx *gin.Context) {
	products, err := pc.uomUseCase.GetAllUom()
	if err != nil {
		utils.HandleInternalServerError(ctx, err.Error())
	} else {
		utils.HandleSuccess(ctx, products, "Success Get All data products")
	}
}

func (uc *UomController) CreateNewUom(ctx *gin.Context) {
	var newUom *model.Uom
	errBinding := ctx.ShouldBindJSON(&newUom)
	if errBinding == nil {
		err := uc.uomUseCase.RegisterUom(newUom)
		if err != nil {
			utils.HandleInternalServerError(ctx, err.Error())
		} else {
			utils.HandleSuccessCreated(ctx, newUom, "Success create new product")
		}
	} else {
		utils.HandleBadRequest(ctx, errBinding.Error())
	}
}

func NewUomController(router *gin.Engine, productUc usecase.UomUseCase) *UomController {
	newUomController := UomController{
		uomUseCase: productUc,
	}
	router.GET("/uoms", newUomController.GetAllUom)
	router.POST("/uoms", newUomController.CreateNewUom)
	return &newUomController
}
