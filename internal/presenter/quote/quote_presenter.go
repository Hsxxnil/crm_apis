package quote

import (
	"net/http"
	"strconv"

	"crm/internal/interactor/pkg/util"

	quoteModel "crm/internal/interactor/models/quotes"

	constant "crm/internal/interactor/constants"

	"crm/internal/interactor/manager/quote"
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
	GetByOpportunityIDSingle(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager quote.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: quote.Init(db),
	}
}

// Create
// @Summary 新增報價
// @description 新增報價
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body quotes.Create true "新增報價"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &quoteModel.Create{}
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
// @Summary 取得全部報價
// @description 取得全部報價
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param * body quotes.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=quotes.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &quoteModel.Fields{}
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
// @Summary 取得單一報價
// @description 取得單一報價
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteID path string true "報價ID"
// @success 200 object code.SuccessfulMessage{body=quotes.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes/{quoteID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	quoteID := ctx.Param("quoteID")
	input := &quoteModel.Field{}
	input.QuoteID = quoteID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingleProducts
// @Summary 取得單一報價含產品
// @description 取得單一報價含產品
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteID path string true "報價ID"
// @success 200 object code.SuccessfulMessage{body=quotes.SingleProducts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes/products/{quoteID} [get]
func (c *control) GetBySingleProducts(ctx *gin.Context) {
	quoteID := ctx.Param("quoteID")
	input := &quoteModel.Field{}
	input.QuoteID = quoteID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingleProducts(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetByOpportunityIDSingle
// @Summary 透過商機ID取得最終單一報價含產品
// @description 透過商機ID取得最終單一報價含產品
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param OpportunityID path string true "商機ID"
// @success 200 object code.SuccessfulMessage{body=quotes.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes/get-by-opportunity/{opportunityID} [get]
func (c *control) GetByOpportunityIDSingle(ctx *gin.Context) {
	opportunityID := ctx.Param("opportunityID")
	input := &quoteModel.Field{}
	input.OpportunityID = util.PointerString(opportunityID)
	input.IsFinal = util.PointerBool(true)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingleProducts(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一報價
// @description 刪除單一報價
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteID path string true "報價ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes/{quoteID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	quoteID := ctx.Param("quoteID")
	input := &quoteModel.Field{}
	input.QuoteID = quoteID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一報價
// @description 更新單一報價
// @Tags quote
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param quoteID path string true "報價ID"
// @param * body quotes.Update true "更新報價"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /quotes/{quoteID} [patch]
func (c *control) Update(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	quoteID := ctx.Param("quoteID")
	input := &quoteModel.Update{}
	input.QuoteID = quoteID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(trx, input)
	ctx.JSON(httpCode, codeMessage)
}
