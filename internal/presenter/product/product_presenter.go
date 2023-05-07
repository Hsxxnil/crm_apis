package product

import (
	"net/http"
	"strconv"

	constant "app.eirc/internal/interactor/constants"

	"app.eirc/internal/interactor/manager/product"
	productModel "app.eirc/internal/interactor/models/products"
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
	Manager product.Manager
}

func Init(db *gorm.DB) Control {
	return &control{
		Manager: product.Init(db),
	}
}

// Create
// @Summary 新增產品
// @description 新增產品
// @Tags product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body products.Create true "新增產品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /products [post]
func (c *control) Create(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &productModel.Create{}
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
// @Summary 取得全部產品
// @description 取得全部產品
// @Tags product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=products.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /products [get]
func (c *control) GetByList(ctx *gin.Context) {
	input := &productModel.Fields{}
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	input.Limit, _ = strconv.ParseInt(limit, 10, 64)
	input.Page, _ = strconv.ParseInt(page, 10, 64)
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
// @Summary 取得單一產品
// @description 取得單一產品
// @Tags product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param productID path string true "產品ID"
// @success 200 object code.SuccessfulMessage{body=products.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /products/{productID} [get]
func (c *control) GetBySingle(ctx *gin.Context) {
	productID := ctx.Param("productID")
	input := &productModel.Field{}
	input.ProductID = productID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.GetBySingle(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一產品
// @description 刪除單一產品
// @Tags product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param productID path string true "產品ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /products/{productID} [delete]
func (c *control) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	productID := ctx.Param("productID")
	input := &productModel.Field{}
	input.ProductID = productID
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := c.Manager.Delete(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Update
// @Summary 更新單一產品
// @description 更新單一產品
// @Tags product
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param productID path string true "產品ID"
// @param * body products.Update true "更新產品"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /products/{productID} [patch]
func (c *control) Update(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	productID := ctx.Param("productID")
	input := &productModel.Update{}
	input.ProductID = productID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	//input.UpdatedBy = util.PointerString(ctx.MustGet("user_id").(string))
	codeMessage := c.Manager.Update(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
