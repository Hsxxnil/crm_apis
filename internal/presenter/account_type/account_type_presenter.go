package account_type

import (
	"net/http"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/account_type"
	accountTypesModel "app.eirc/internal/interactor/models/account_types"
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
	Manager account_type.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: account_type.Init(db),
	}
}

// Create
// @Summary 新增帳戶類型
// @description 新增帳戶類型
// @Tags account_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body account_types.Create true "新增帳戶類型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /crm/v1.0/account_types [post]
func (c *control) Create(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &accountTypesModel.Create{}
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
// @Summary 取得全部帳戶類型
// @description 取得全部帳戶類型
// @Tags account_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=account_types.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /crm/v1.0/account_types [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &accountTypesModel.Fields{}

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
// @Summary 取得單一帳戶類型
// @description 取得單一帳戶類型
// @Tags account_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountTypeID path string true "帳戶類型ID"
// @success 200 object code.SuccessfulMessage{body=account_types.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /crm/v1.0/account_types/{accountTypeID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	accountTypeID := ctx.Param("accountTypeID")
	input := &accountTypesModel.Field{}
	input.AccountTypeID = accountTypeID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一帳戶類型
// @description 刪除單一帳戶類型
// @Tags account_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountTypeID path string true "帳戶類型ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /crm/v1.0/account_types/{accountTypeID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	accountTypeID := ctx.Param("accountTypeID")
	input := &accountTypesModel.Field{}
	input.AccountTypeID = accountTypeID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Update
// @Summary 更新單一帳戶類型
// @description 更新單一帳戶類型
// @Tags account_type
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountTypeID path string true "帳戶類型ID"
// @param * body account_types.Update true "更新帳戶類型"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /crm/v1.0/account_types/{accountTypeID} [patch]
func (c *control) Update(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	accountTypeID := ctx.Param("accountTypeID")
	input := &accountTypesModel.Update{}
	input.AccountTypeID = accountTypeID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	codeMessage := c.Manager.Update(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
