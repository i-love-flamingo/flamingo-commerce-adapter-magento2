// Code generated by go-swagger; DO NOT EDIT.

package catalog_inventory_stock_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams creates a new CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams object
// with the default values initialized.
func NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams() *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	var ()
	return &CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParamsWithTimeout creates a new CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParamsWithTimeout(timeout time.Duration) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	var ()
	return &CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams{

		timeout: timeout,
	}
}

// NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParamsWithContext creates a new CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParamsWithContext(ctx context.Context) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	var ()
	return &CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams{

		Context: ctx,
	}
}

// NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParamsWithHTTPClient creates a new CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParamsWithHTTPClient(client *http.Client) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	var ()
	return &CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams{
		HTTPClient: client,
	}
}

/*CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams contains all the parameters to send to the API endpoint
for the catalog inventory stock registry v1 update stock item by sku put operation typically these are written to a http.Request
*/
type CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams struct {

	/*CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody*/
	CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody
	/*ItemID*/
	ItemID string
	/*ProductSku*/
	ProductSku string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WithTimeout(timeout time.Duration) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WithContext(ctx context.Context) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WithHTTPClient(client *http.Client) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody adds the catalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WithCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody(catalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	o.SetCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody(catalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody)
	return o
}

// SetCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody adds the catalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) SetCatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody(catalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody) {
	o.CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody = catalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody
}

// WithItemID adds the itemID to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WithItemID(itemID string) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) SetItemID(itemID string) {
	o.ItemID = itemID
}

// WithProductSku adds the productSku to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WithProductSku(productSku string) *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams {
	o.SetProductSku(productSku)
	return o
}

// SetProductSku adds the productSku to the catalog inventory stock registry v1 update stock item by sku put params
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) SetProductSku(productSku string) {
	o.ProductSku = productSku
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CatalogInventoryStockRegistryV1UpdateStockItemBySkuPutBody); err != nil {
		return err
	}

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID); err != nil {
		return err
	}

	// path param productSku
	if err := r.SetPathParam("productSku", o.ProductSku); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
