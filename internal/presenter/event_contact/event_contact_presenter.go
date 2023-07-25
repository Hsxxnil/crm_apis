package event_contact

import (
	"net/http"

	"app.eirc/internal/interactor/pkg/util"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/event_contact"
	eventContactModel "app.eirc/internal/interactor/models/event_contacts"
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
	Manager event_contact.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: event_contact.Init(db),
	}
}

// Create
// @Summary 新增事件聯絡人
// @description 新增事件聯絡人
// @Tags event_contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body event_contacts.Create true "新增事件聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-contacts [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &eventContactModel.Create{}
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
// @Summary 取得全部事件聯絡人
// @description 取得全部事件聯絡人
// @Tags event_contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=event_contacts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-contacts [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &eventContactModel.Fields{}
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
// @Summary 取得單一事件聯絡人
// @description 取得單一事件聯絡人
// @Tags event_contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventContactID path string true "事件聯絡人ID"
// @success 200 object code.SuccessfulMessage{body=event_contacts.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-contacts/{eventContactID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	eventContactID := ctx.Param("eventContactID")
	input := &eventContactModel.Field{}
	input.EventContactID = eventContactID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一事件聯絡人
// @description 刪除單一事件聯絡人
// @Tags event_contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventContactID path string true "事件聯絡人ID"
// @param * body event_contacts.Update true "更新事件聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-contacts/{eventContactID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	eventContactID := ctx.Param("eventContactID")
	input := &eventContactModel.Update{}
	input.EventContactID = eventContactID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一事件聯絡人
// @description 更新單一事件聯絡人
// @Tags event_contact
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventContactID path string true "事件聯絡人ID"
// @param * body event_contacts.Update true "更新事件聯絡人"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-contacts/{eventContactID} [patch]
func (c *control) Update(ctx *gin.Context) {
	eventContactID := ctx.Param("eventContactID")
	input := &eventContactModel.Update{}
	input.EventContactID = eventContactID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
