package event_user_attendee

import (
	"net/http"

	"app.eirc/internal/interactor/pkg/util"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/event_user_attendee"
	eventUserAttendeeModel "app.eirc/internal/interactor/models/event_user_attendees"
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
	Manager event_user_attendee.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: event_user_attendee.Init(db),
	}
}

// Create
// @Summary 新增事件參與人員
// @description 新增事件參與人員
// @Tags event_user_attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body event_user_attendees.Create true "新增事件參與人員"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-users-attendee [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &eventUserAttendeeModel.Create{}
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
// @Summary 取得全部事件參與人員
// @description 取得全部事件參與人員
// @Tags event_user_attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=event_user_attendees.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-users-attendee [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &eventUserAttendeeModel.Fields{}
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
// @Summary 取得單一事件參與人員
// @description 取得單一事件參與人員
// @Tags event_user_attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventUserAttendeeID path string true "事件參與人員ID"
// @success 200 object code.SuccessfulMessage{body=event_user_attendees.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-users-attendee/{eventUserAttendeeID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	eventUserAttendeeID := ctx.Param("eventUserAttendeeID")
	input := &eventUserAttendeeModel.Field{}
	input.EventUserAttendeeID = eventUserAttendeeID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一事件參與人員
// @description 刪除單一事件參與人員
// @Tags event_user_attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventUserAttendeeID path string true "事件參與人員ID"
// @param * body event_user_attendees.Update true "更新事件參與人員"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-users-attendee/{eventUserAttendeeID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	eventUserAttendeeID := ctx.Param("eventUserAttendeeID")
	input := &eventUserAttendeeModel.Update{}
	input.EventUserAttendeeID = eventUserAttendeeID
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
// @Summary 更新單一事件參與人員
// @description 更新單一事件參與人員
// @Tags event_user_attendee
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param eventUserAttendeeID path string true "事件參與人員ID"
// @param * body event_user_attendees.Update true "更新事件參與人員"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /events-users-attendee/{eventUserAttendeeID} [patch]
func (c *control) Update(ctx *gin.Context) {
	eventUserAttendeeID := ctx.Param("eventUserAttendeeID")
	input := &eventUserAttendeeModel.Update{}
	input.EventUserAttendeeID = eventUserAttendeeID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
