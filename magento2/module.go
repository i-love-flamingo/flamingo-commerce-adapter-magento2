package magento2

import (
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure"
	cartAdapter "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/cart"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/product"
	"flamingo.me/dingo"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/client"
	"flamingo.me/flamingo-commerce/v3/cart/domain/cart"
	"flamingo.me/flamingo-commerce/v3/product/domain"
	"flamingo.me/flamingo/v3/framework/config"
	inMemoryProductSearchInfrastructure "flamingo.me/flamingo-commerce-adapter-standalone/inMemoryProductSearch/infrastructure"
)

type (
	// Module is the (dingo) module for this package - loaded for default configs
	Module struct {
	}

	// ProductModule load this module to get the product adapter
	ProductModule struct {
	}

	// ProductSearchModule load this module to get the product adapter
	ProductSearchModule struct {
	}

	// CartModule load this module to get the cart adapters
	CartModule struct {
	}
)
var (
	//check interface compatibility
	_ product.Repository = &inMemoryProductSearchInfrastructure.InMemoryProductRepository{}
)
// Configure  - dingo module main method
func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind((*product.Repository)(nil)).ToProvider(
		func(loader *product.RepositoryLoader) product.Repository {
			rep := &inMemoryProductSearchInfrastructure.InMemoryProductRepository{}
			loader.AddFromMagento2(rep)
			return rep
		}).In(dingo.ChildSingleton)

	injector.Bind((*client.MagentoCommerceForB2B)(nil)).ToProvider(infrastructure.MagentoApiClientProvider)
}

// Configure  - dingo module main method
func (m *ProductModule) Configure(injector *dingo.Injector) {
	//injector.Bind(new(cart.CustomerCartService)).To(cartadapter.CustomerCartServiceAdapter{})
	injector.Bind(new(domain.ProductService)).To(product.ServiceAdapter{})
}


// Configure  - dingo module main method
func (m *ProductSearchModule) Configure(injector *dingo.Injector) {
	//injector.Bind(new(cart.CustomerCartService)).To(cartadapter.CustomerCartServiceAdapter{})
	injector.Bind(new(domain.SearchService)).To(product.ServiceAdapter{})
}

// Configure - dingo module main method
func (m *CartModule) Configure(injector *dingo.Injector) {
	injector.Bind(new(cart.CustomerCartService)).To(cartAdapter.CustomerCartServiceAdapter{})
}

// DefaultConfig for magento adapter module
func (m *Module) DefaultConfig() config.Map {
	return config.Map{
		"flamingo-commerce-adapter-magento2.magento2": config.Map{
			"accessToken":        "",
			"storeCode":          "default",
			"apiCallTimeout":     "5s",
			"apiCallTimeoutSlow": "15s",
		},
	}
}
