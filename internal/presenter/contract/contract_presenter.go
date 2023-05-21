package contract

import (
	"net/http"
	"strconv"

	"app.eirc/internal/interactor/pkg/util"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/contract"
	contractModel "app.eirc/internal/interactor/models/contracts"
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
	Manager contract.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: contract.Init(db),
	}
}

// Create
// @Summary 新增契約
// @description 新增契約
// @Tags contract
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body contracts.Create true "新增契約"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contracts [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &contractModel.Create{}
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
// @Summary 取得全部契約
// @description 取得全部契約
// @Tags contract
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @param sort query string false "排序"
// @param direction query string false "排序方式"
// @param * body contracts.Filter false "搜尋"
// @success 200 object code.SuccessfulMessage{body=contracts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contracts/list [post]
func (c *control) GetByList(ctx *gin.Context) {
	input := &contractModel.Fields{}
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

// GetBySingle
// @Summary 取得單一契約
// @description 取得單一契約
// @Tags contract
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param contractID path string true "契約ID"
// @success 200 object code.SuccessfulMessage{body=contracts.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contracts/{contractID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	contractID := ctx.Param("contractID")
	input := &contractModel.Field{}
	input.ContractID = contractID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一契約
// @description 刪除單一契約
// @Tags contract
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param contractID path string true "契約ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contracts/{contractID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	contractID := ctx.Param("contractID")
	input := &contractModel.Field{}
	input.ContractID = contractID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Delete(input)
	ctx.JSON(httpCode, codeMessage)
}

// Update
// @Summary 更新單一契約
// @description 更新單一契約
// @Tags contract
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param contractID path string true "契約ID"
// @param * body contracts.Update true "更新契約"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /contracts/{contractID} [patch]
func (c *control) Update(ctx *gin.Context) {
	contractID := ctx.Param("contractID")
	input := &contractModel.Update{}
	input.ContractID = contractID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
