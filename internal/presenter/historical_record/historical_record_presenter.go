package historical_record

import (
	"net/http"
	"strconv"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/pkg/util"

	"app.eirc/internal/interactor/manager/historical_record"
	historicalRecordModel "app.eirc/internal/interactor/models/historical_records"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	GetByList(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
}

type control struct {
	Manager historical_record.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: historical_record.Init(db),
	}
}

// GetByList
// @Summary 透過來源ID取得全部歷程記錄
// @description 透過來源ID取得全部歷程記錄
// @Tags historical_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sourceID path string true "來源ID"
// @success 200 object code.SuccessfulMessage{body=historical_records.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /historical-records/list/{sourceID} [post]
func (c *control) GetByList(ctx *gin.Context) {
	sourceID := ctx.Param("sourceID")
	input := &historicalRecordModel.Fields{}
	input.SourceID = util.PointerString(sourceID)
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	input.Limit, _ = strconv.ParseInt(limit, 10, 64)
	input.Page, _ = strconv.ParseInt(page, 10, 64)

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
// @Summary 取得單一歷程記錄
// @description 取得單一歷程記錄
// @Tags historical_record
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param historicalRecordID path string true "歷程記錄ID"
// @success 200 object code.SuccessfulMessage{body=historical_records.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /historical-records/{historicalRecordID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	historicalRecordID := ctx.Param("historicalRecordID")
	input := &historicalRecordModel.Field{}
	input.HistoricalRecordID = historicalRecordID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}
