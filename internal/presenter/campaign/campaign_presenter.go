package campaign

import (
	"net/http"
	"strconv"

	"crm/internal/interactor/pkg/util"

	constant "crm/internal/interactor/constants"

	"crm/internal/interactor/manager/campaign"
	campaignModel "crm/internal/interactor/models/campaigns"
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
	GetBySingleOpportunities(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager campaign.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: campaign.Init(db),
	}
}

// Create
// @Summary 新增行銷活動
// @description 新增行銷活動
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body campaigns.Create true "新增行銷活動"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &campaignModel.Create{}
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
// @Summary 取得全部行銷活動
// @description 取得全部行銷活動
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param * body campaigns.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=campaigns.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &campaignModel.Fields{}
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
// @Summary 取得全部行銷活動(不用page和limit)
// @description 取得全部行銷活動(不用page和limit)
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=campaigns.ListNoPagination} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns [get]
func (c *control) GetByListNoPagination(ctx *gin.Context) {
	input := &campaignModel.Field{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetByListNoPagination(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingle
// @Summary 取得單一行銷活動
// @description 取得單一行銷活動
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param campaignID path string true "行銷活動ID"
// @success 200 object code.SuccessfulMessage{body=campaigns.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns/{campaignID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	campaignID := ctx.Param("campaignID")
	input := &campaignModel.Field{}
	input.CampaignID = campaignID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// GetBySingleOpportunities
// @Summary 取得單一行銷活動含受影響的商機
// @description 取得單一行銷活動含受影響的商機
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param campaignID path string true "行銷活動ID"
// @success 200 object code.SuccessfulMessage{body=campaigns.SingleOpportunities} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns/opportunities/{campaignID} [get]
func (c *control) GetBySingleOpportunities(ctx *gin.Context) {
	campaignID := ctx.Param("campaignID")
	input := &campaignModel.Field{}
	input.CampaignID = campaignID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingleOpportunities(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一行銷活動
// @description 刪除單一行銷活動
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param campaignID path string true "行銷活動ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns/{campaignID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	campaignID := ctx.Param("campaignID")
	input := &campaignModel.Field{}
	input.CampaignID = campaignID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一行銷活動
// @description 更新單一行銷活動
// @Tags campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param campaignID path string true "行銷活動ID"
// @param * body campaigns.Update true "更新行銷活動"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /campaigns/{campaignID} [patch]
func (c *control) Update(ctx *gin.Context) {
	campaignID := ctx.Param("campaignID")
	input := &campaignModel.Update{}
	input.CampaignID = campaignID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
