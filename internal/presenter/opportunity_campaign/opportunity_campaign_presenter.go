package opportunity_campaign

import (
	"net/http"

	"crm/internal/interactor/pkg/util"

	constant "crm/internal/interactor/constants"

	"crm/internal/interactor/manager/opportunity_campaign"
	opportunityCampaignModel "crm/internal/interactor/models/opportunity_campaigns"
	"crm/internal/interactor/pkg/util/code"
	"crm/internal/interactor/pkg/util/log"

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
	Manager opportunity_campaign.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: opportunity_campaign.Init(db),
	}
}

// Create
// @Summary 新增商機行銷活動
// @description 新增商機行銷活動
// @Tags opportunity-campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body opportunity_campaigns.Create true "新增商機行銷活動"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities-campaigns [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &opportunityCampaignModel.Create{}
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
// @Summary 取得全部商機行銷活動
// @description 取得全部商機行銷活動
// @Tags opportunity-campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=opportunity_campaigns.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities-campaigns [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &opportunityCampaignModel.Fields{}

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
// @Summary 取得單一商機行銷活動
// @description 取得單一商機行銷活動
// @Tags opportunity-campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityCampaignID path string true "商機行銷活動ID"
// @success 200 object code.SuccessfulMessage{body=opportunity_campaigns.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities-campaigns/{opportunityCampaignID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	opportunityCampaignID := ctx.Param("opportunityCampaignID")
	input := &opportunityCampaignModel.Field{}
	input.OpportunityCampaignID = opportunityCampaignID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一商機行銷活動
// @description 刪除單一商機行銷活動
// @Tags opportunity-campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityCampaignID path string true "商機行銷活動ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities-campaigns/{opportunityCampaignID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	opportunityCampaignID := ctx.Param("opportunityCampaignID")
	input := &opportunityCampaignModel.Field{}
	input.OpportunityCampaignID = opportunityCampaignID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一商機行銷活動
// @description 更新單一商機行銷活動
// @Tags opportunity-campaign
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityCampaignID path string true "商機行銷活動ID"
// @param * body opportunity_campaigns.Update true "更新商機行銷活動"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities-campaigns/{opportunityCampaignID} [patch]
func (c *control) Update(ctx *gin.Context) {
	opportunityCampaignID := ctx.Param("opportunityCampaignID")
	input := &opportunityCampaignModel.Update{}
	input.OpportunityCampaignID = opportunityCampaignID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
