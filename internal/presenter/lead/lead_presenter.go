package lead

import (
	"net/http"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/lead"
	leadModel "app.eirc/internal/interactor/models/leads"
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
	Manager lead.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: lead.Init(db),
	}
}

// Create
// @Summary 新增商機線索
// @description 新增商機線索
// @Tags lead
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body leads.Create true "新增商機線索"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/lead [post]
func (c *control) Create(ctx *gin.Context) {
	// Todo 將UUID改成登入的商機線索
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &leadModel.Create{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByList
// @Summary 取得全部商機線索
// @description 取得全部商機線索
// @Tags lead
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=leads.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/lead [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &leadModel.Fields{}

	if err := ctx.ShouldBindQuery(input); err != nil {
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
// @Summary 取得單一商機線索
// @description 取得單一商機線索
// @Tags lead
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param leadID path string true "商機線索ID"
// @success 200 object code.SuccessfulMessage{body=leads.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/lead/{leadID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	leadID := ctx.Param("leadID") // 跟router對應
	input := &leadModel.Field{}
	input.LeadID = leadID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一商機線索
// @description 刪除單一商機線索
// @Tags lead
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param leadID path string true "商機線索ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/lead/{leadID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的商機線索
	leadID := ctx.Param("leadID")
	input := &leadModel.Field{}
	input.LeadID = leadID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Update
// @Summary 更新單一商機線索
// @description 更新單一商機線索
// @Tags lead
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param leadID path string true "商機線索ID"
// @param * body leads.Update true "更新商機線索"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/lead/{leadID} [patch]
func (c *control) Update(ctx *gin.Context) {
	// Todo 將UUID改成登入的商機線索
	leadID := ctx.Param("leadID")
	input := &leadModel.Update{}
	input.LeadID = leadID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Update(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
