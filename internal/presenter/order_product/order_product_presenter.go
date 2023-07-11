package order_product

import (
	"net/http"

	"app.eirc/internal/interactor/pkg/util"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/order_product"
	orderProductModel "app.eirc/internal/interactor/models/order_products"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager order_product.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: order_product.Init(db),
	}
}

// Create
// @Summary 新增多筆訂單產品
// @description 新增多筆訂單產品
// @Tags order-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body order_products.CreateList true "新增多筆訂單產品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders-products [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &orderProductModel.CreateList{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	for _, value := range input.OrderProducts {
		value.CreatedBy = ctx.MustGet("user_id").(string)
	}

	httpCode, codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(httpCode, codeMessage)
}

// GetByList
// @Summary 取得全部訂單產品
// @description 取得全部訂單產品
// @Tags order-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=order_products.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders-products [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &orderProductModel.Fields{}

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= constant.DefaultLimit {
		input.Limit = constant.DefaultLimit
	}

	httpCode, codeMessage := c.Manager.GetByList(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一訂單產品
// @description 取得單一訂單產品
// @Tags order-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param orderProductID path string true "訂單產品ID"
// @success 200 object code.SuccessfulMessage{body=order_products.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders-products/{orderProductID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	orderProductID := ctx.Param("orderProductID")
	input := &orderProductModel.Field{}
	input.OrderProductID = orderProductID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除多筆訂單產品
// @description 刪除多筆訂單產品
// @Tags order-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders-products [delete]
func (c *control) Delete(ctx *gin.Context) {
	inputIDList := &orderProductModel.DeleteList{}
	if err := ctx.ShouldBindQuery(inputIDList); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	input := &orderProductModel.UpdateList{
		OrderProducts: make([]*orderProductModel.Update, len(inputIDList.OrderProducts)),
	}
	for i, OrderProductID := range inputIDList.OrderProducts {
		input.OrderProducts[i] = &orderProductModel.Update{
			OrderProductID: OrderProductID,
		}
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新多筆訂單產品
// @description 更新多筆訂單產品
// @Tags order-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body order_products.Update true "更新訂單產品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders-products [patch]
func (c *control) Update(ctx *gin.Context) {
	input := &orderProductModel.UpdateList{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	for _, value := range input.OrderProducts {
		value.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
