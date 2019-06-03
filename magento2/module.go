package magento2

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure"
	productSearchAdapter "flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/productSearch"
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/client"
	productSearchModule "flamingo.me/flamingo-commerce-adapter-standalone/productSearch"

	"flamingo.me/flamingo-commerce-adapter-standalone/productSearch/infrastructure/productSearch"
	"flamingo.me/flamingo/v3/framework/config"
)

type (
	Module struct {
	}
	// ProductModule load this module to get the product adapter
	ProductModule struct {
	}

	CartModule struct {

	}
)

var (
)
// Configure  - dingo module main method
func (m *Module) Configure(injector *dingo.Injector) {

	injector.Bind((*client.MagentoCommerceForB2B)(nil)).ToProvider(infrastructure.MagentoApiClientProvider)
}

// Configure  - dingo module main method
func (m *ProductModule) Configure(injector *dingo.Injector) {
	//Register Loader for productSearch
	injector.Bind((*productSearch.Loader)(nil)).To(productSearchAdapter.CsvLoader{}).In(dingo.ChildSingleton)
	//injector.Bind(new(cart.CustomerCartService)).To(cartadapter.CustomerCartServiceAdapter{})
	//injector.Bind(new(domain.ProductService)).To(product.ServiceAdapter{})
}


// Configure - dingo module main method
func (m *CartModule) Configure(injector *dingo.Injector) {
	//injector.Bind(new(cart.CustomerCartService)).To(cartAdapter.CustomerCartServiceAdapter{})
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




// Depends on other modules
func (m *ProductModule) Depends() []dingo.Module {
	return []dingo.Module{
		new(productSearchModule.Module),
	}
}