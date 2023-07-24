package event

import (
	"net/http"

	"app.eirc/internal/interactor/pkg/util"

	"app.eirc/internal/interactor/manager/event"
	eventModel "app.eirc/internal/interactor/models/events"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager event.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: event.Init(db),
	}
}

// Create
// @Summary 新增事件
// @description 新增事件
// @Tags event
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body events.Create true "新增事件"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &eventModel.Create{}
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
// @Summary 取得全部事件
// @description 取得全部事件
// @Tags event
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body events.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=events.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &eventModel.Fields{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetByList(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一事件
// @description 取得單一事件
// @Tags event
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventID path string true "事件ID"
// @success 200 object code.SuccessfulMessage{body=events.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events/{eventID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	eventID := ctx.Param("eventID")
	input := &eventModel.Field{}
	input.EventID = eventID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一事件
// @description 刪除單一事件
// @Tags event
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventID path string true "事件ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events/{eventID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	eventID := ctx.Param("eventID")
	input := &eventModel.Field{}
	input.EventID = eventID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一事件
// @description 更新單一事件
// @Tags event
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventID path string true "事件ID"
// @param * body events.Update true "更新事件"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events/{eventID} [patch]
func (c *control) Update(ctx *gin.Context) {
	eventID := ctx.Param("eventID")
	input := &eventModel.Update{}
	input.EventID = eventID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
