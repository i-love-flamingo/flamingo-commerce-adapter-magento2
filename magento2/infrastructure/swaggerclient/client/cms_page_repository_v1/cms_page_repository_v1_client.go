// Code generated by go-swagger; DO NOT EDIT.

package cms_page_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cms page repository v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cms page repository v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CmsPageRepositoryV1DeleteByIDDelete Delete page by ID.
*/
func (a *Client) CmsPageRepositoryV1DeleteByIDDelete(params *CmsPageRepositoryV1DeleteByIDDeleteParams) (*CmsPageRepositoryV1DeleteByIDDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmsPageRepositoryV1DeleteByIDDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmsPageRepositoryV1DeleteByIdDelete",
		Method:             "DELETE",
		PathPattern:        "/V1/cmsPage/{pageId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CmsPageRepositoryV1DeleteByIDDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CmsPageRepositoryV1DeleteByIDDeleteOK), nil

}

/*
CmsPageRepositoryV1GetByIDGet Retrieve page.
*/
func (a *Client) CmsPageRepositoryV1GetByIDGet(params *CmsPageRepositoryV1GetByIDGetParams) (*CmsPageRepositoryV1GetByIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmsPageRepositoryV1GetByIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmsPageRepositoryV1GetByIdGet",
		Method:             "GET",
		PathPattern:        "/V1/cmsPage/{pageId}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CmsPageRepositoryV1GetByIDGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CmsPageRepositoryV1GetByIDGetOK), nil

}

/*
CmsPageRepositoryV1GetListGet Retrieve pages matching the specified criteria.
*/
func (a *Client) CmsPageRepositoryV1GetListGet(params *CmsPageRepositoryV1GetListGetParams) (*CmsPageRepositoryV1GetListGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmsPageRepositoryV1GetListGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmsPageRepositoryV1GetListGet",
		Method:             "GET",
		PathPattern:        "/V1/cmsPage/search",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CmsPageRepositoryV1GetListGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CmsPageRepositoryV1GetListGetOK), nil

}

/*
CmsPageRepositoryV1SavePost Save page.
*/
func (a *Client) CmsPageRepositoryV1SavePost(params *CmsPageRepositoryV1SavePostParams) (*CmsPageRepositoryV1SavePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmsPageRepositoryV1SavePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmsPageRepositoryV1SavePost",
		Method:             "POST",
		PathPattern:        "/V1/cmsPage",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CmsPageRepositoryV1SavePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CmsPageRepositoryV1SavePostOK), nil

}

/*
CmsPageRepositoryV1SavePut Save page.
*/
func (a *Client) CmsPageRepositoryV1SavePut(params *CmsPageRepositoryV1SavePutParams) (*CmsPageRepositoryV1SavePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCmsPageRepositoryV1SavePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cmsPageRepositoryV1SavePut",
		Method:             "PUT",
		PathPattern:        "/V1/cmsPage/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CmsPageRepositoryV1SavePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CmsPageRepositoryV1SavePutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
