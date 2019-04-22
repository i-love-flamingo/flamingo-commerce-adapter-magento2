package infrastructure

import (
	"flamingo.me/flamingo-commerce-adapter-magento2/magento2/infrastructure/swaggerclient/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)



//MagentoApiClientProvider
func MagentoApiClientProvider(config *struct {
	AccessToken string `inject:"config:flamingo-commerce-adapter-magento2.magento2.accessToken"`
	Host string `inject:"config:flamingo-commerce-adapter-magento2.magento2.host"`
	Path string `inject:"config:flamingo-commerce-adapter-magento2.magento2.path,optional"`
	Scheme string `inject:"config:flamingo-commerce-adapter-magento2.magento2.scheme"`
}) *client.MagentoCommerceForB2B {
	// create the API client

	transport := httptransport.New(config.Host, config.Path, []string{config.Scheme})
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
		return req.SetHeaderParam("Authorization", "Bearer "+config.AccessToken)
	})
	return client.New(transport, strfmt.Default)
}

