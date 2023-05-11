package opportunity

import (
	"net/http"
	"strconv"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/opportunity"
	opportunityModel "app.eirc/internal/interactor/models/opportunities"
	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Control interface {
	Create(ctx *gin.Context)
	GetByList(ctx *gin.Context)
	GetBySingle(ctx *gin.Context)
	GetBySingleCampaigns(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type control struct {
	Manager opportunity.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: opportunity.Init(db),
	}
}

// Create
// @Summary 新增商機
// @description 新增商機
// @Tags opportunity
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body opportunities.Create true "新增商機"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities [post]
func (c *control) Create(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &opportunityModel.Create{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//input.CreatedBy = ctx.MustGet("user_id").(string)
	codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByList
// @Summary 取得全部商機
// @description 取得全部商機
// @Tags opportunity
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param search query string false "搜尋"
// @success 200 object code.SuccessfulMessage{body=opportunities.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &opportunityModel.Fields{}
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	input.Limit, _ = strconv.ParseInt(limit, 10, 64)
	input.Page, _ = strconv.ParseInt(page, 10, 64)

	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= constant.DefaultLimit {
		input.Limit = constant.DefaultLimit
	}

	codeMessage := c.Manager.GetByList(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetBySingle
// @Summary 取得單一商機
// @description 取得單一商機
// @Tags opportunity
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityID path string true "商機ID"
// @success 200 object code.SuccessfulMessage{body=opportunities.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities/{opportunityID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	opportunityID := ctx.Param("opportunityID")
	input := &opportunityModel.Field{}
	input.OpportunityID = opportunityID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetBySingleCampaigns
// @Summary 取得單一商機含影響的行銷活動
// @description 取得單一商機含影響的行銷活動
// @Tags opportunity
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityID path string true "商機ID"
// @success 200 object code.SuccessfulMessage{body=opportunities.SingleCampaigns} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities/campaigns/{opportunityID} [get]
func (c *control) GetBySingleCampaigns(ctx *gin.Context) {
	opportunityID := ctx.Param("opportunityID")
	input := &opportunityModel.Field{}
	input.OpportunityID = opportunityID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.GetBySingleCampaigns(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一商機
// @description 刪除單一商機
// @Tags opportunity
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityID path string true "商機ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities/{opportunityID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	opportunityID := ctx.Param("opportunityID")
	input := &opportunityModel.Field{}
	input.OpportunityID = opportunityID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Update
// @Summary 更新單一商機
// @description 更新單一商機
// @Tags opportunity
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param opportunityID path string true "商機ID"
// @param * body opportunities.Update true "更新商機"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /opportunities/{opportunityID} [patch]
func (c *control) Update(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	opportunityID := ctx.Param("opportunityID")
	input := &opportunityModel.Update{}
	input.OpportunityID = opportunityID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	codeMessage := c.Manager.Update(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
