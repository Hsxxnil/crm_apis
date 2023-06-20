package product

import (
	"encoding/json"
	"errors"

	"app.eirc/internal/interactor/pkg/util"

	productModel "app.eirc/internal/interactor/models/products"
	productService "app.eirc/internal/interactor/service/product"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *productModel.Create) (int, interface{})
	GetByList(input *productModel.Fields) (int, interface{})
	GetBySingle(input *productModel.Field) (int, interface{})
	Delete(input *productModel.Field) (int, interface{})
	Update(input *productModel.Update) (int, interface{})
}

type manager struct {
	ProductService productService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ProductService: productService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *productModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 判斷產品識別碼是否重複
	quantity, _ := m.ProductService.GetByQuantity(&productModel.Field{
		Code: util.PointerString(input.Code),
	})

	if quantity > 0 {
		log.Info("Code already exists. Code: ", input.Code)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Code already exists.")
	}

	productBase, err := m.ProductService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, productBase.ProductID)
}

func (m *manager) GetByList(input *productModel.Fields) (int, interface{}) {
	output := &productModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, productBase, err := m.ProductService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Total.Total = quantity
	productByte, err := json.Marshal(productBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(productByte, &output.Products)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, products := range output.Products {
		products.CreatedBy = *productBase[i].CreatedByUsers.Name
		products.UpdatedBy = *productBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *productModel.Field) (int, interface{}) {
	productBase, err := m.ProductService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &productModel.Single{}
	productByte, _ := json.Marshal(productBase)
	err = json.Unmarshal(productByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *productBase.CreatedByUsers.Name
	output.UpdatedBy = *productBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *productModel.Field) (int, interface{}) {
	_, err := m.ProductService.GetBySingle(&productModel.Field{
		ProductID: input.ProductID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.ProductService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *productModel.Update) (int, interface{}) {
	productBase, err := m.ProductService.GetBySingle(&productModel.Field{
		ProductID: input.ProductID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 判斷產品識別碼是否重複
	if *productBase.Code != input.Code {
		quantity, _ := m.ProductService.GetByQuantity(&productModel.Field{
			Code: util.PointerString(input.Code),
		})
		if quantity > 0 {
			log.Info("Code already exists. Code: ", input.Code)
			return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Code already exists.")
		}
	}

	err = m.ProductService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, productBase.ProductID)
}
