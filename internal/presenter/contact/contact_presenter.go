package contact

import (
	"net/http"
	"strconv"

	"app.eirc/internal/interactor/pkg/util"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/contact"
	contactModel "app.eirc/internal/interactor/models/contacts"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	GetByAccountIDListNoPagination(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager contact.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: contact.Init(db),
	}
}

// Create
// @Summary 新增聯絡人
// @description 新增聯絡人
// @Tags contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body contacts.Create true "新增聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contacts [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &contactModel.Create{}
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
// @Summary 取得全部聯絡人
// @description 取得全部聯絡人
// @Tags contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param search query string false "搜尋"
// @param * body contacts.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=contacts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contacts/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &contactModel.Fields{}
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

// GetByAccountIDListNoPagination
// @Summary 取得全部聯絡人(不用page和limit)
// @description 取得全部聯絡人(不用page和limit)
// @Tags contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=contacts.ListNoPagination} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contacts/get-by-account/{accountID} [get]
func (c *control) GetByAccountIDListNoPagination(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	input := &contactModel.Field{}
	input.AccountID = util.PointerString(accountID)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetByListNoPagination(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一聯絡人
// @description 取得單一聯絡人
// @Tags contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param contactID path string true "聯絡人ID"
// @success 200 object code.SuccessfulMessage{body=contacts.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contacts/{contactID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	contactID := ctx.Param("contactID")
	input := &contactModel.Field{}
	input.ContactID = contactID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一聯絡人
// @description 刪除單一聯絡人
// @Tags contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param contactID path string true "聯絡人ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contacts/{contactID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	contactID := ctx.Param("contactID")
	input := &contactModel.Field{}
	input.ContactID = contactID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一聯絡人
// @description 更新單一聯絡人
// @Tags contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param contactID path string true "聯絡人ID"
// @param * body contacts.Update true "更新聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contacts/{contactID} [patch]
func (c *control) Update(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	contactID := ctx.Param("contactID")
	input := &contactModel.Update{}
	input.ContactID = contactID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(trx, input)
	ctx.JSON(httpCode, codeMessage)
}
