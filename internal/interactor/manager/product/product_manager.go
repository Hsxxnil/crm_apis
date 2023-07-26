package product

import (
	"encoding/json"
	"errors"

	quoteModel "app.eirc/internal/interactor/models/quotes"
	quoteService "app.eirc/internal/interactor/service/quote"

	contractModel "app.eirc/internal/interactor/models/contracts"
	orderModel "app.eirc/internal/interactor/models/orders"
	contractService "app.eirc/internal/interactor/service/contract"

	quoteProductDB "app.eirc/internal/entity/postgresql/db/quote_products"
	productModel "app.eirc/internal/interactor/models/products"
	quoteProductModel "app.eirc/internal/interactor/models/quote_products"
	"app.eirc/internal/interactor/pkg/util"
	orderService "app.eirc/internal/interactor/service/order"
	productService "app.eirc/internal/interactor/service/product"
	quoteProductService "app.eirc/internal/interactor/service/quote_product"
	"gorm.io/gorm"

	"app.eirc/internal/interactor/pkg/util/code"
	"app.eirc/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *productModel.Create) (int, interface{})
	GetByList(input *productModel.Fields) (int, interface{})
	GetByOrderIDList(input *productModel.Fields) (int, interface{})
	GetBySingle(input *productModel.Field) (int, interface{})
	Delete(input *productModel.Field) (int, interface{})
	Update(input *productModel.Update) (int, interface{})
}

type manager struct {
	ProductService      productService.Service
	QuoteProductService quoteProductService.Service
	OrderService        orderService.Service
	ContractService     contractService.Service
	QuoteService        quoteService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		ProductService:      productService.Init(db),
		QuoteProductService: quoteProductService.Init(db),
		OrderService:        orderService.Init(db),
		ContractService:     contractService.Init(db),
		QuoteService:        quoteService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *productModel.Create) (int, interface{}) {
	defer trx.Rollback()

	// 判斷產品識別碼是否重複
	quantity, _ := m.ProductService.GetByQuantity(&productModel.Field{
		Code:      util.PointerString(input.Code),
		IsDeleted: util.PointerBool(false),
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
	input.IsDeleted = util.PointerBool(false)
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

func (m *manager) GetByOrderIDList(input *productModel.Fields) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
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

	// 透過訂單ID取得契約ID
	orderBase, _ := m.OrderService.GetBySingle(&orderModel.Field{
		OrderID:   *input.OrderID,
		IsDeleted: util.PointerBool(false),
	})

	// 透過契約ID取得商機ID
	contractBase, _ := m.ContractService.GetBySingle(&contractModel.Field{
		ContractID: *orderBase.ContractID,
		IsDeleted:  util.PointerBool(false),
	})

	// 透過商機ID取得報價ID
	quoteBase, _ := m.QuoteService.GetBySingle(&quoteModel.Field{
		OpportunityID: contractBase.OpportunityID,
		IsDeleted:     util.PointerBool(false),
	})

	// 收集所有產品ID
	productIDs := make([]string, len(productBase))
	for i, product := range productBase {
		productIDs[i] = *product.ProductID
	}

	// 透過報價ID取得所有與該報價有關的產品ID
	quoteProductBase, _ := m.QuoteProductService.GetByListNoQuantity(&quoteProductModel.Field{
		QuoteID:   quoteBase.QuoteID,
		IsDeleted: util.PointerBool(false),
	})

	// 建立產品ID的映射表
	productIDMap := make(map[string]bool)
	for _, productID := range productIDs {
		productIDMap[productID] = true
	}

	// 將相同產品ID的報價產品與產品對應
	var matchingQuoteProductBase []*quoteProductDB.Base
	for _, quoteProduct := range quoteProductBase {
		if productIDMap[*quoteProduct.ProductID] {
			matchingQuoteProductBase = append(matchingQuoteProductBase, quoteProduct)
		}
	}

	for i, products := range output.Products {
		products.CreatedBy = *productBase[i].CreatedByUsers.Name
		products.UpdatedBy = *productBase[i].UpdatedByUsers.Name
		for _, quoteProduct := range matchingQuoteProductBase {
			if *quoteProduct.ProductID == products.ProductID {
				products.QuotePrice = *quoteProduct.UnitPrice
			}
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *productModel.Field) (int, interface{}) {
	input.IsDeleted = util.PointerBool(false)
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
		IsDeleted: util.PointerBool(false),
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
		IsDeleted: util.PointerBool(false),
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
			Code:      util.PointerString(input.Code),
			IsDeleted: util.PointerBool(false),
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
