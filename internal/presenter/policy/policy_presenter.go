package policy

import (
	"net/http"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
	casbin "app.eirc/internal/router/middleware/auth"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

type Presenter interface {
	AddPolicy(ctx *gin.Context)
	GetPolicy(ctx *gin.Context)
	DeletePolicy(ctx *gin.Context)
}

// AddPolicy
// @Summary 新增策略
// @description 新增策略
// @Tags policy
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body auth.CasbinBind true "新增策略"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /policies [post]
func AddPolicy(ctx *gin.Context) {
	input := &casbin.CasbinModel{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	result, err := casbin.AddPolicy(*input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error()))
		return
	}

	if result != true {
		log.Error(err)
		ctx.JSON(http.StatusBadRequest, code.GetCodeMessage(code.BadRequest, "policy already exists."))
		return
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, "Add successful!"))
}

// GetAllPolicies
// @Summary 取得策略
// @description 取得策略
// @Tags policy
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @success 200 object code.SuccessfulMessage{body=[]auth.CasbinOutput} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /policies [get]
func GetAllPolicies(ctx *gin.Context) {

	result := casbin.GetAllPolicies()

	if result == nil {
		ctx.JSON(http.StatusNotFound, code.GetCodeMessage(code.DoesNotExist, "policy does not exist."))
		return
	}

	var output []casbin.CasbinOutput
	for _, value := range result {
		output = append(output, casbin.CasbinOutput{
			RoleName: value[0],
			Path:     value[1],
			Method:   value[2],
		})
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, output))
}

// DeletePolicy
// @Summary 刪除策略
// @description 刪除策略
// @Tags policy
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body auth.CasbinBind true "刪除策略"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /policies [delete]
func DeletePolicy(ctx *gin.Context) {
	input := &casbin.CasbinModel{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusUnsupportedMediaType, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	result, err := casbin.DeletePolicy(*input)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error()))
		return
	}

	if result != true {
		log.Error(err)
		ctx.JSON(http.StatusNotFound, code.GetCodeMessage(code.DoesNotExist, "policy does not exist."))
		return
	}

	ctx.JSON(http.StatusOK, code.GetCodeMessage(code.Successful, "Delete ok!"))
}
