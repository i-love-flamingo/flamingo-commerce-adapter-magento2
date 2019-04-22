package product

import (
	"context"
	"flamingo.me/flamingo-commerce/v3/product/domain"
	searchDomain "flamingo.me/flamingo-commerce/v3/search/domain"
)

type (
	//ProductServiceAdapter the Adapter
	ServiceAdapter struct {
		repository Repository
	}

	//ProductRepository - interface required by the adapter
	Repository interface {
		Add(product domain.BasicProduct) error
		FindByMarketplaceCode(marketplaceCode string) (domain.BasicProduct, error)
		Find(filters ...searchDomain.Filter) (*domain.SearchResult, error)
	}
)


// Inject - dingo compatible Inject to get the dependencies passed
func (ps *ServiceAdapter) Inject(repository Repository) {
	ps.repository = repository
}



// Get returns a product struct (implements ProductService)
func (ps *ServiceAdapter) Get(ctx context.Context, marketplaceCode string) (domain.BasicProduct, error) {
	return ps.repository.FindByMarketplaceCode(marketplaceCode)
}


// Search returns a Search Result for the given context and supplied filters
func (p *ServiceAdapter) Search(ctx context.Context, filter ...searchDomain.Filter) (*domain.SearchResult, error) {
	return p.repository.Find(filter...)
}

// SearchBy returns Products prefiltered by the given attribute (also based on additional given Filters) e.g. SearchBy(ctx,"brandCode","apple")
func (p *ServiceAdapter) SearchBy(ctx context.Context, attribute string, values []string, filter ...searchDomain.Filter) (*domain.SearchResult, error) {
	return p.Search(ctx, nil)
}
