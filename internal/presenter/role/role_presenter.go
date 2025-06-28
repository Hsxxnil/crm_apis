package role

import (
	"net/http"

	"crm/internal/interactor/pkg/util"

	constant "crm/internal/interactor/constants"

	"crm/internal/interactor/manager/role"
	roleModel "crm/internal/interactor/models/roles"
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
	Manager role.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: role.Init(db),
	}
}

// Create
// @Summary 新增角色
// @description 新增角色
// @Tags role
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body roles.Create true "新增角色"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /roles [post]
func (c *control) Create(ctx *gin.Context) {
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &roleModel.Create{}
	input.CompanyID = ctx.MustGet("company_id").(string)
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
// @Summary 取得全部角色
// @description 取得全部角色
// @Tags role
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=roles.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /roles [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &roleModel.Fields{}
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
// @Summary 取得單一角色
// @description 取得單一角色
// @Tags role
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param roleID path string true "角色ID"
// @success 200 object code.SuccessfulMessage{body=roles.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /roles/{roleID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	roleID := ctx.Param("roleID")
	input := &roleModel.Field{}
	input.RoleID = roleID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(httpCode, codeMessage)
}

// Delete
// @Summary 刪除單一角色
// @description 刪除單一角色
// @Tags role
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param roleID path string true "角色ID"
// @param * body roles.Update true "更新角色"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /roles/{roleID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	roleID := ctx.Param("roleID")
	input := &roleModel.Update{}
	input.RoleID = roleID
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
// @Summary 更新單一角色
// @description 更新單一角色
// @Tags role
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param roleID path string true "角色ID"
// @param * body roles.Update true "更新角色"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /roles/{roleID} [patch]
func (c *control) Update(ctx *gin.Context) {
	roleID := ctx.Param("roleID")
	input := &roleModel.Update{}
	input.RoleID = roleID
	input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	httpCode, codeMessage := c.Manager.Update(input)
	ctx.JSON(httpCode, codeMessage)
}
