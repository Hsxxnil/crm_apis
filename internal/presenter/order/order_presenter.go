package order

import (
	"net/http"
	"strconv"

	"crm/internal/interactor/pkg/util"

	constant "crm/internal/interactor/constants"

	"crm/internal/interactor/manager/order"
	orderModel "crm/internal/interactor/models/orders"
	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	GetBySingleProducts(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager order.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: order.Init(db),
	}
}

// Create
// @Summary 新增訂單
// @description 新增訂單
// @Tags order
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body orders.Create true "新增訂單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &orderModel.Create{}
	input.CreatedBy = ctx.MustGet("user_id").(string)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(httpCode, codeMessage)
}

// GetByList
// @Summary 取得全部訂單
// @description 取得全部訂單
// @Tags order
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param * body orders.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=orders.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &orderModel.Fields{}
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	input.Limit, _ = strconv.ParseInt(limit, 10, 64)
	input.Page, _ = strconv.ParseInt(page, 10, 64)

	if err := ctx.ShouldBindJSON(input); err != nil {
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
// @Summary 取得單一訂單
// @description 取得單一訂單
// @Tags order
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param orderID path string true "訂單ID"
// @success 200 object code.SuccessfulMessage{body=orders.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders/{orderID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	input := &orderModel.Field{}
	input.OrderID = orderID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingleProducts
// @Summary 取得單一訂單含產品
// @description 取得單一訂單含產品
// @Tags order
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param orderID path string true "訂單ID"
// @success 200 object code.SuccessfulMessage{body=orders.SingleProducts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders/products/{orderID} [get]
func (c *control) GetBySingleProducts(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	input := &orderModel.Field{}
	input.OrderID = orderID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingleProducts(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一訂單
// @description 刪除單一訂單
// @Tags order
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param orderID path string true "訂單ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders/{orderID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	input := &orderModel.Field{}
	input.OrderID = orderID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一訂單
// @description 更新單一訂單
// @Tags order
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param orderID path string true "訂單ID"
// @param * body orders.Update true "更新訂單"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /orders/{orderID} [patch]
func (c *control) Update(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	orderID := ctx.Param("orderID")
	input := &orderModel.Update{}
	input.OrderID = orderID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(trx, input)
	ctx.JSON(httpCode, codeMessage)
}
