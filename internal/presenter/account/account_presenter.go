package account

import (
	"net/http"
	"strconv"

	"crm/internal/interactor/pkg/util"

	constant "crm/internal/interactor/constants"

	"crm/internal/interactor/manager/account"
	accountModel "crm/internal/interactor/models/accounts"
	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	GetByListNoPagination(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	GetBySingleContacts(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager account.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: account.Init(db),
	}
}

// Create
// @Summary 新增帳戶
// @description 新增帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body accounts.Create true "新增帳戶"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &accountModel.Create{}
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
// @Summary 取得全部帳戶
// @description 取得全部帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param * body accounts.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=accounts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &accountModel.Fields{}
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

// GetByListNoPagination
// @Summary 取得全部帳戶(不用page和limit)
// @description 取得全部帳戶(不用page和limit)
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body accounts.FilterNoPagination false "搜尋"
// @success 200 object code.SuccessfulMessage{body=accounts.ListNoPagination} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts/list/no-pagination [post]
func (c *control) GetByListNoPagination(ctx *gin.Context) {
	input := &accountModel.FieldsNoPagination{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetByListNoPagination(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一帳戶
// @description 取得單一帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @success 200 object code.SuccessfulMessage{body=accounts.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts/{accountID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	input := &accountModel.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingleContacts
// @Summary 取得單一帳戶含聯絡人
// @description 取得單一帳戶含聯絡人
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @success 200 object code.SuccessfulMessage{body=accounts.SingleContacts} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts/contacts/{accountID} [get]
func (c *control) GetBySingleContacts(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	input := &accountModel.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingleContacts(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一帳戶
// @description 刪除單一帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts/{accountID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	input := &accountModel.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一帳戶
// @description 更新單一帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @param * body accounts.Update true "更新帳戶"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /accounts/{accountID} [patch]
func (c *control) Update(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	accountID := ctx.Param("accountID")
	input := &accountModel.Update{}
	input.AccountID = accountID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(trx, input)
	ctx.JSON(httpCode, codeMessage)
}
