package productSearch

import (
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/client"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/client/catalog_product_repository_v1"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/client/catalog_category_repository_v1"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/models"
	priceDomain "flamingo.me/flamingo-commerce/v3/price/domain"
	"flamingo.me/flamingo-commerce/v3/product/domain"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"time"

	"flamingo.me/flamingo-commerce-adapter-standalone/productSearch/infrastructure/productSearch"
)

type (
	ApiLoader struct {
		client *client.MagentoCommerceForB2B
		logger flamingo.Logger
	}
)

func (l *ApiLoader) Inject(client *client.MagentoCommerceForB2B, logger flamingo.Logger) {
	l.client = client
	l.logger = logger.WithField(flamingo.LogKeyModule, "magento2").WithField(flamingo.LogKeyCategory, "ApiLoader")
}



func (l *ApiLoader) Load(indexer productSearch.Index) error {
	l.logger.Info("Start to load products into repository..")
	params := catalog_product_repository_v1.NewCatalogProductRepositoryV1GetListGetParams()
	pageSize := int64(1000)
	params.SearchCriteriaPageSize = &pageSize
	params.SetTimeout(time.Second * 240)
	result, err := l.client.CatalogProductRepositoryV1.CatalogProductRepositoryV1GetListGet(params)
	if err != nil {
		l.logger.Error(err)
		panic(err)
	}
	for _, item := range result.Payload.Items {
		l.logger.Debug("adding sku ", *item.Sku)
		if item.TypeID == "simple" {
			err = indexer.Add(buildSimple(item))
			if err != nil {
				l.logger.Warn(err)
			}
		}
	}
	return nil
}

func buildSimple(item *models.CatalogDataProductInterface) domain.BasicProduct {

	return domain.SimpleProduct{
		Identifier: *item.Sku,
		BasicProductData: domain.BasicProductData{
			Title:           item.Name,
			MarketPlaceCode: *item.Sku,
			Description:     getCustomAttributeString(item.CustomAttributes, "description"),
		},
		Saleable: domain.Saleable{
			ActivePrice: domain.PriceInfo{Default: priceDomain.NewFromFloat(item.Price, "€")},
		},
	}
}

func getCustomAttributeString(list []*models.FrameworkAttributeInterface, key string) string {
	for _, ca := range list {
		if *ca.AttributeCode == key {
			return ca.Value.(string)
		}
	}
	return ""
}



func (l *ApiLoader) getCategoryTeaser(id int64) *domain.CategoryTeaser {
	cparams := catalog_category_repository_v1.NewCatalogCategoryRepositoryV1GetGetParams()
	cparams.SetTimeout(time.Second * 60)
	cparams.CategoryID = id
	result, err := l.client.CatalogCategoryRepositoryV1.CatalogCategoryRepositoryV1GetGet(cparams)
	if err != nil {
		l.logger.Error(err)
		panic(err)
	}
	return &domain.CategoryTeaser{
		Code: string(id),
		Name: *result.Payload.Name,
	}
}