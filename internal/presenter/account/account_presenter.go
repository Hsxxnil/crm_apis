package account

import (
	"net/http"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/account"
	accountModel "app.eirc/internal/interactor/models/accounts"
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
	Manager account.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: account.Init(db),
	}
}

// Create
// @Summary 新增帳戶
// @description 新增帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body accounts.Create true "新增帳戶"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts [post]
func (c *control) Create(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &accountModel.Create{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Create(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByList
// @Summary 取得全部帳戶
// @description 取得全部帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=accounts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &accountModel.Fields{}

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
// @Summary 取得單一帳戶
// @description 取得單一帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @success 200 object code.SuccessfulMessage{body=accounts.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts/{accountID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	accountID := ctx.Param("accountID") // 跟router對應
	input := &accountModel.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一帳戶
// @description 刪除單一帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts/{accountID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	accountID := ctx.Param("accountID")
	input := &accountModel.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Update
// @Summary 更新單一帳戶
// @description 更新單一帳戶
// @Tags account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "帳戶ID"
// @param * body accounts.Update true "更新帳戶"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts/{accountID} [patch]
func (c *control) Update(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	accountID := ctx.Param("accountID")
	input := &accountModel.Update{}
	input.AccountID = accountID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Update(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
