package quote_product

import (
	"net/http"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/quote_product"
	quoteProductModel "app.eirc/internal/interactor/models/quote_products"
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
	Manager quote_product.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: quote_product.Init(db),
	}
}

// Create
// @Summary 新增報價產品
// @description 新增報價產品
// @Tags quote-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body quote_products.Create true "新增報價產品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes-products [post]
func (c *control) Create(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &quoteProductModel.Create{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//input.CreatedBy = ctx.MustGet("user_id").(string)
	httpCode, codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(httpCode, codeMessage)
}

// GetByList
// @Summary 取得全部報價產品
// @description 取得全部報價產品
// @Tags quote-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=quote_products.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes-products [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &quoteProductModel.Fields{}

	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= constant.DefaultLimit {
		input.Limit = constant.DefaultLimit
	}

	httpCode, codeMessage := c.Manager.GetByList(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一報價產品
// @description 取得單一報價產品
// @Tags quote-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteProductID path string true "報價產品ID"
// @success 200 object code.SuccessfulMessage{body=quote_products.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes-products/{quoteProductID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	quoteProductID := ctx.Param("quoteProductID")
	input := &quoteProductModel.Field{}
	input.QuoteProductID = quoteProductID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一報價產品
// @description 刪除單一報價產品
// @Tags quote-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteProductID path string true "報價產品ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes-products/{quoteProductID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	quoteProductID := ctx.Param("quoteProductID")
	input := &quoteProductModel.Field{}
	input.QuoteProductID = quoteProductID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一報價產品
// @description 更新單一報價產品
// @Tags quote-product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteProductID path string true "報價產品ID"
// @param * body quote_products.Update true "更新報價產品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes-products/{quoteProductID} [patch]
func (c *control) Update(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	quoteProductID := ctx.Param("quoteProductID")
	input := &quoteProductModel.Update{}
	input.QuoteProductID = quoteProductID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
