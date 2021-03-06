// Code generated by go-swagger; DO NOT EDIT.

package sales_invoice_repository_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSalesInvoiceRepositoryV1GetListGetParams creates a new SalesInvoiceRepositoryV1GetListGetParams object
// with the default values initialized.
func NewSalesInvoiceRepositoryV1GetListGetParams() *SalesInvoiceRepositoryV1GetListGetParams {
	var ()
	return &SalesInvoiceRepositoryV1GetListGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSalesInvoiceRepositoryV1GetListGetParamsWithTimeout creates a new SalesInvoiceRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSalesInvoiceRepositoryV1GetListGetParamsWithTimeout(timeout time.Duration) *SalesInvoiceRepositoryV1GetListGetParams {
	var ()
	return &SalesInvoiceRepositoryV1GetListGetParams{

		timeout: timeout,
	}
}

// NewSalesInvoiceRepositoryV1GetListGetParamsWithContext creates a new SalesInvoiceRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSalesInvoiceRepositoryV1GetListGetParamsWithContext(ctx context.Context) *SalesInvoiceRepositoryV1GetListGetParams {
	var ()
	return &SalesInvoiceRepositoryV1GetListGetParams{

		Context: ctx,
	}
}

// NewSalesInvoiceRepositoryV1GetListGetParamsWithHTTPClient creates a new SalesInvoiceRepositoryV1GetListGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSalesInvoiceRepositoryV1GetListGetParamsWithHTTPClient(client *http.Client) *SalesInvoiceRepositoryV1GetListGetParams {
	var ()
	return &SalesInvoiceRepositoryV1GetListGetParams{
		HTTPClient: client,
	}
}

/*SalesInvoiceRepositoryV1GetListGetParams contains all the parameters to send to the API endpoint
for the sales invoice repository v1 get list get operation typically these are written to a http.Request
*/
type SalesInvoiceRepositoryV1GetListGetParams struct {

	/*SearchCriteriaCurrentPage
	  Current page.

	*/
	SearchCriteriaCurrentPage *int64
	/*SearchCriteriaFilterGroupsFiltersConditionType
	  Condition type

	*/
	SearchCriteriaFilterGroupsFiltersConditionType *string
	/*SearchCriteriaFilterGroupsFiltersField
	  Field

	*/
	SearchCriteriaFilterGroupsFiltersField *string
	/*SearchCriteriaFilterGroupsFiltersValue
	  Value

	*/
	SearchCriteriaFilterGroupsFiltersValue *string
	/*SearchCriteriaPageSize
	  Page size.

	*/
	SearchCriteriaPageSize *int64
	/*SearchCriteriaSortOrdersDirection
	  Sorting direction.

	*/
	SearchCriteriaSortOrdersDirection *string
	/*SearchCriteriaSortOrdersField
	  Sorting field.

	*/
	SearchCriteriaSortOrdersField *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithTimeout(timeout time.Duration) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithContext(ctx context.Context) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithHTTPClient(client *http.Client) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearchCriteriaCurrentPage adds the searchCriteriaCurrentPage to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaCurrentPage(searchCriteriaCurrentPage *int64) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaCurrentPage(searchCriteriaCurrentPage)
	return o
}

// SetSearchCriteriaCurrentPage adds the searchCriteriaCurrentPage to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaCurrentPage(searchCriteriaCurrentPage *int64) {
	o.SearchCriteriaCurrentPage = searchCriteriaCurrentPage
}

// WithSearchCriteriaFilterGroupsFiltersConditionType adds the searchCriteriaFilterGroupsFiltersConditionType to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaFilterGroupsFiltersConditionType(searchCriteriaFilterGroupsFiltersConditionType *string) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaFilterGroupsFiltersConditionType(searchCriteriaFilterGroupsFiltersConditionType)
	return o
}

// SetSearchCriteriaFilterGroupsFiltersConditionType adds the searchCriteriaFilterGroupsFiltersConditionType to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaFilterGroupsFiltersConditionType(searchCriteriaFilterGroupsFiltersConditionType *string) {
	o.SearchCriteriaFilterGroupsFiltersConditionType = searchCriteriaFilterGroupsFiltersConditionType
}

// WithSearchCriteriaFilterGroupsFiltersField adds the searchCriteriaFilterGroupsFiltersField to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaFilterGroupsFiltersField(searchCriteriaFilterGroupsFiltersField *string) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaFilterGroupsFiltersField(searchCriteriaFilterGroupsFiltersField)
	return o
}

// SetSearchCriteriaFilterGroupsFiltersField adds the searchCriteriaFilterGroupsFiltersField to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaFilterGroupsFiltersField(searchCriteriaFilterGroupsFiltersField *string) {
	o.SearchCriteriaFilterGroupsFiltersField = searchCriteriaFilterGroupsFiltersField
}

// WithSearchCriteriaFilterGroupsFiltersValue adds the searchCriteriaFilterGroupsFiltersValue to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaFilterGroupsFiltersValue(searchCriteriaFilterGroupsFiltersValue *string) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaFilterGroupsFiltersValue(searchCriteriaFilterGroupsFiltersValue)
	return o
}

// SetSearchCriteriaFilterGroupsFiltersValue adds the searchCriteriaFilterGroupsFiltersValue to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaFilterGroupsFiltersValue(searchCriteriaFilterGroupsFiltersValue *string) {
	o.SearchCriteriaFilterGroupsFiltersValue = searchCriteriaFilterGroupsFiltersValue
}

// WithSearchCriteriaPageSize adds the searchCriteriaPageSize to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaPageSize(searchCriteriaPageSize *int64) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaPageSize(searchCriteriaPageSize)
	return o
}

// SetSearchCriteriaPageSize adds the searchCriteriaPageSize to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaPageSize(searchCriteriaPageSize *int64) {
	o.SearchCriteriaPageSize = searchCriteriaPageSize
}

// WithSearchCriteriaSortOrdersDirection adds the searchCriteriaSortOrdersDirection to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaSortOrdersDirection(searchCriteriaSortOrdersDirection *string) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaSortOrdersDirection(searchCriteriaSortOrdersDirection)
	return o
}

// SetSearchCriteriaSortOrdersDirection adds the searchCriteriaSortOrdersDirection to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaSortOrdersDirection(searchCriteriaSortOrdersDirection *string) {
	o.SearchCriteriaSortOrdersDirection = searchCriteriaSortOrdersDirection
}

// WithSearchCriteriaSortOrdersField adds the searchCriteriaSortOrdersField to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) WithSearchCriteriaSortOrdersField(searchCriteriaSortOrdersField *string) *SalesInvoiceRepositoryV1GetListGetParams {
	o.SetSearchCriteriaSortOrdersField(searchCriteriaSortOrdersField)
	return o
}

// SetSearchCriteriaSortOrdersField adds the searchCriteriaSortOrdersField to the sales invoice repository v1 get list get params
func (o *SalesInvoiceRepositoryV1GetListGetParams) SetSearchCriteriaSortOrdersField(searchCriteriaSortOrdersField *string) {
	o.SearchCriteriaSortOrdersField = searchCriteriaSortOrdersField
}

// WriteToRequest writes these params to a swagger request
func (o *SalesInvoiceRepositoryV1GetListGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SearchCriteriaCurrentPage != nil {

		// query param searchCriteria[currentPage]
		var qrSearchCriteriaCurrentPage int64
		if o.SearchCriteriaCurrentPage != nil {
			qrSearchCriteriaCurrentPage = *o.SearchCriteriaCurrentPage
		}
		qSearchCriteriaCurrentPage := swag.FormatInt64(qrSearchCriteriaCurrentPage)
		if qSearchCriteriaCurrentPage != "" {
			if err := r.SetQueryParam("searchCriteria[currentPage]", qSearchCriteriaCurrentPage); err != nil {
				return err
			}
		}

	}

	if o.SearchCriteriaFilterGroupsFiltersConditionType != nil {

		// query param searchCriteria[filterGroups][][filters][][conditionType]
		var qrSearchCriteriaFilterGroupsFiltersConditionType string
		if o.SearchCriteriaFilterGroupsFiltersConditionType != nil {
			qrSearchCriteriaFilterGroupsFiltersConditionType = *o.SearchCriteriaFilterGroupsFiltersConditionType
		}
		qSearchCriteriaFilterGroupsFiltersConditionType := qrSearchCriteriaFilterGroupsFiltersConditionType
		if qSearchCriteriaFilterGroupsFiltersConditionType != "" {
			if err := r.SetQueryParam("searchCriteria[filterGroups][][filters][][conditionType]", qSearchCriteriaFilterGroupsFiltersConditionType); err != nil {
				return err
			}
		}

	}

	if o.SearchCriteriaFilterGroupsFiltersField != nil {

		// query param searchCriteria[filterGroups][][filters][][field]
		var qrSearchCriteriaFilterGroupsFiltersField string
		if o.SearchCriteriaFilterGroupsFiltersField != nil {
			qrSearchCriteriaFilterGroupsFiltersField = *o.SearchCriteriaFilterGroupsFiltersField
		}
		qSearchCriteriaFilterGroupsFiltersField := qrSearchCriteriaFilterGroupsFiltersField
		if qSearchCriteriaFilterGroupsFiltersField != "" {
			if err := r.SetQueryParam("searchCriteria[filterGroups][][filters][][field]", qSearchCriteriaFilterGroupsFiltersField); err != nil {
				return err
			}
		}

	}

	if o.SearchCriteriaFilterGroupsFiltersValue != nil {

		// query param searchCriteria[filterGroups][][filters][][value]
		var qrSearchCriteriaFilterGroupsFiltersValue string
		if o.SearchCriteriaFilterGroupsFiltersValue != nil {
			qrSearchCriteriaFilterGroupsFiltersValue = *o.SearchCriteriaFilterGroupsFiltersValue
		}
		qSearchCriteriaFilterGroupsFiltersValue := qrSearchCriteriaFilterGroupsFiltersValue
		if qSearchCriteriaFilterGroupsFiltersValue != "" {
			if err := r.SetQueryParam("searchCriteria[filterGroups][][filters][][value]", qSearchCriteriaFilterGroupsFiltersValue); err != nil {
				return err
			}
		}

	}

	if o.SearchCriteriaPageSize != nil {

		// query param searchCriteria[pageSize]
		var qrSearchCriteriaPageSize int64
		if o.SearchCriteriaPageSize != nil {
			qrSearchCriteriaPageSize = *o.SearchCriteriaPageSize
		}
		qSearchCriteriaPageSize := swag.FormatInt64(qrSearchCriteriaPageSize)
		if qSearchCriteriaPageSize != "" {
			if err := r.SetQueryParam("searchCriteria[pageSize]", qSearchCriteriaPageSize); err != nil {
				return err
			}
		}

	}

	if o.SearchCriteriaSortOrdersDirection != nil {

		// query param searchCriteria[sortOrders][][direction]
		var qrSearchCriteriaSortOrdersDirection string
		if o.SearchCriteriaSortOrdersDirection != nil {
			qrSearchCriteriaSortOrdersDirection = *o.SearchCriteriaSortOrdersDirection
		}
		qSearchCriteriaSortOrdersDirection := qrSearchCriteriaSortOrdersDirection
		if qSearchCriteriaSortOrdersDirection != "" {
			if err := r.SetQueryParam("searchCriteria[sortOrders][][direction]", qSearchCriteriaSortOrdersDirection); err != nil {
				return err
			}
		}

	}

	if o.SearchCriteriaSortOrdersField != nil {

		// query param searchCriteria[sortOrders][][field]
		var qrSearchCriteriaSortOrdersField string
		if o.SearchCriteriaSortOrdersField != nil {
			qrSearchCriteriaSortOrdersField = *o.SearchCriteriaSortOrdersField
		}
		qSearchCriteriaSortOrdersField := qrSearchCriteriaSortOrdersField
		if qSearchCriteriaSortOrdersField != "" {
			if err := r.SetQueryParam("searchCriteria[sortOrders][][field]", qSearchCriteriaSortOrdersField); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
