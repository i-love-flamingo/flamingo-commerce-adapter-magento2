module flamingo.me/flamingo-commerce-adapter-magento2

go 1.12

replace flamingo.me/flamingo/v3 => ../flamingo

replace flamingo.me/flamingo-commerce/v3 => ../flamingo-commerce

require (
	flamingo.me/dingo v0.1.5
	flamingo.me/flamingo-commerce-adapter-standalone v0.0.1-beta
	flamingo.me/flamingo-commerce/v3 v3.0.0-beta.1.0.20190421120452-cfdd1db5f8b8
	flamingo.me/flamingo/v3 v3.0.0-beta.2.0.20190417151243-244258f10588

	github.com/go-openapi/errors v0.19.0
	github.com/go-openapi/runtime v0.19.0
	github.com/go-openapi/strfmt v0.19.0
	github.com/go-openapi/swag v0.19.0
	github.com/go-openapi/validate v0.19.0
)
