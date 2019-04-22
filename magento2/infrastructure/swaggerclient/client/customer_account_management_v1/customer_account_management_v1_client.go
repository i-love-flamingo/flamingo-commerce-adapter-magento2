// Code generated by go-swagger; DO NOT EDIT.

package customer_account_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new customer account management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer account management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomerAccountManagementV1ActivateByIDPut Activate a customer account using a key that was sent in a confirmation email.
*/
func (a *Client) CustomerAccountManagementV1ActivateByIDPut(params *CustomerAccountManagementV1ActivateByIDPutParams) (*CustomerAccountManagementV1ActivateByIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ActivateByIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ActivateByIdPut",
		Method:             "PUT",
		PathPattern:        "/V1/customers/me/activate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ActivateByIDPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ActivateByIDPutOK), nil

}

/*
CustomerAccountManagementV1ActivatePut Activate a customer account using a key that was sent in a confirmation email.
*/
func (a *Client) CustomerAccountManagementV1ActivatePut(params *CustomerAccountManagementV1ActivatePutParams) (*CustomerAccountManagementV1ActivatePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ActivatePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ActivatePut",
		Method:             "PUT",
		PathPattern:        "/V1/customers/{email}/activate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ActivatePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ActivatePutOK), nil

}

/*
CustomerAccountManagementV1ChangePasswordByIDPut Change customer password.
*/
func (a *Client) CustomerAccountManagementV1ChangePasswordByIDPut(params *CustomerAccountManagementV1ChangePasswordByIDPutParams) (*CustomerAccountManagementV1ChangePasswordByIDPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ChangePasswordByIDPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ChangePasswordByIdPut",
		Method:             "PUT",
		PathPattern:        "/V1/customers/me/password",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ChangePasswordByIDPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ChangePasswordByIDPutOK), nil

}

/*
CustomerAccountManagementV1CreateAccountPost Create customer account. Perform necessary business operations like sending email.
*/
func (a *Client) CustomerAccountManagementV1CreateAccountPost(params *CustomerAccountManagementV1CreateAccountPostParams) (*CustomerAccountManagementV1CreateAccountPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1CreateAccountPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1CreateAccountPost",
		Method:             "POST",
		PathPattern:        "/V1/customers",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1CreateAccountPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1CreateAccountPostOK), nil

}

/*
CustomerAccountManagementV1GetConfirmationStatusGet Gets the account confirmation status.
*/
func (a *Client) CustomerAccountManagementV1GetConfirmationStatusGet(params *CustomerAccountManagementV1GetConfirmationStatusGetParams) (*CustomerAccountManagementV1GetConfirmationStatusGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1GetConfirmationStatusGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1GetConfirmationStatusGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/{customerId}/confirm",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1GetConfirmationStatusGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1GetConfirmationStatusGetOK), nil

}

/*
CustomerAccountManagementV1GetDefaultBillingAddressGet Retrieve default billing address for the given customerId.
*/
func (a *Client) CustomerAccountManagementV1GetDefaultBillingAddressGet(params *CustomerAccountManagementV1GetDefaultBillingAddressGetParams) (*CustomerAccountManagementV1GetDefaultBillingAddressGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1GetDefaultBillingAddressGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1GetDefaultBillingAddressGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/{customerId}/billingAddress",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1GetDefaultBillingAddressGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1GetDefaultBillingAddressGetOK), nil

}

/*
CustomerAccountManagementV1GetDefaultBillingAddressGetMe Retrieve default billing address for the given customerId.
*/
func (a *Client) CustomerAccountManagementV1GetDefaultBillingAddressGetMe(params *CustomerAccountManagementV1GetDefaultBillingAddressGetMeParams) (*CustomerAccountManagementV1GetDefaultBillingAddressGetMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1GetDefaultBillingAddressGetMeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1GetDefaultBillingAddressGetMe",
		Method:             "GET",
		PathPattern:        "/V1/customers/me/billingAddress",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1GetDefaultBillingAddressGetMeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1GetDefaultBillingAddressGetMeOK), nil

}

/*
CustomerAccountManagementV1GetDefaultShippingAddressGet Retrieve default shipping address for the given customerId.
*/
func (a *Client) CustomerAccountManagementV1GetDefaultShippingAddressGet(params *CustomerAccountManagementV1GetDefaultShippingAddressGetParams) (*CustomerAccountManagementV1GetDefaultShippingAddressGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1GetDefaultShippingAddressGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1GetDefaultShippingAddressGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/{customerId}/shippingAddress",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1GetDefaultShippingAddressGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1GetDefaultShippingAddressGetOK), nil

}

/*
CustomerAccountManagementV1GetDefaultShippingAddressGetMe Retrieve default shipping address for the given customerId.
*/
func (a *Client) CustomerAccountManagementV1GetDefaultShippingAddressGetMe(params *CustomerAccountManagementV1GetDefaultShippingAddressGetMeParams) (*CustomerAccountManagementV1GetDefaultShippingAddressGetMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1GetDefaultShippingAddressGetMeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1GetDefaultShippingAddressGetMe",
		Method:             "GET",
		PathPattern:        "/V1/customers/me/shippingAddress",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1GetDefaultShippingAddressGetMeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1GetDefaultShippingAddressGetMeOK), nil

}

/*
CustomerAccountManagementV1InitiatePasswordResetPut Send an email to the customer with a password reset link.
*/
func (a *Client) CustomerAccountManagementV1InitiatePasswordResetPut(params *CustomerAccountManagementV1InitiatePasswordResetPutParams) (*CustomerAccountManagementV1InitiatePasswordResetPutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1InitiatePasswordResetPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1InitiatePasswordResetPut",
		Method:             "PUT",
		PathPattern:        "/V1/customers/password",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1InitiatePasswordResetPutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1InitiatePasswordResetPutOK), nil

}

/*
CustomerAccountManagementV1IsEmailAvailablePost Check if given email is associated with a customer account in given website.
*/
func (a *Client) CustomerAccountManagementV1IsEmailAvailablePost(params *CustomerAccountManagementV1IsEmailAvailablePostParams) (*CustomerAccountManagementV1IsEmailAvailablePostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1IsEmailAvailablePostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1IsEmailAvailablePost",
		Method:             "POST",
		PathPattern:        "/V1/customers/isEmailAvailable",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1IsEmailAvailablePostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1IsEmailAvailablePostOK), nil

}

/*
CustomerAccountManagementV1IsReadonlyGet Check if customer can be deleted.
*/
func (a *Client) CustomerAccountManagementV1IsReadonlyGet(params *CustomerAccountManagementV1IsReadonlyGetParams) (*CustomerAccountManagementV1IsReadonlyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1IsReadonlyGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1IsReadonlyGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/{customerId}/permissions/readonly",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1IsReadonlyGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1IsReadonlyGetOK), nil

}

/*
CustomerAccountManagementV1ResendConfirmationPost Resend confirmation email.
*/
func (a *Client) CustomerAccountManagementV1ResendConfirmationPost(params *CustomerAccountManagementV1ResendConfirmationPostParams) (*CustomerAccountManagementV1ResendConfirmationPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ResendConfirmationPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ResendConfirmationPost",
		Method:             "POST",
		PathPattern:        "/V1/customers/confirm",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ResendConfirmationPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ResendConfirmationPostOK), nil

}

/*
CustomerAccountManagementV1ResetPasswordPost Reset customer password.
*/
func (a *Client) CustomerAccountManagementV1ResetPasswordPost(params *CustomerAccountManagementV1ResetPasswordPostParams) (*CustomerAccountManagementV1ResetPasswordPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ResetPasswordPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ResetPasswordPost",
		Method:             "POST",
		PathPattern:        "/V1/customers/resetPassword",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ResetPasswordPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ResetPasswordPostOK), nil

}

/*
CustomerAccountManagementV1ValidatePut Validate customer data.
*/
func (a *Client) CustomerAccountManagementV1ValidatePut(params *CustomerAccountManagementV1ValidatePutParams) (*CustomerAccountManagementV1ValidatePutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ValidatePutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ValidatePut",
		Method:             "PUT",
		PathPattern:        "/V1/customers/validate",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ValidatePutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ValidatePutOK), nil

}

/*
CustomerAccountManagementV1ValidateResetPasswordLinkTokenGet Check if password reset token is valid.
*/
func (a *Client) CustomerAccountManagementV1ValidateResetPasswordLinkTokenGet(params *CustomerAccountManagementV1ValidateResetPasswordLinkTokenGetParams) (*CustomerAccountManagementV1ValidateResetPasswordLinkTokenGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerAccountManagementV1ValidateResetPasswordLinkTokenGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "customerAccountManagementV1ValidateResetPasswordLinkTokenGet",
		Method:             "GET",
		PathPattern:        "/V1/customers/{customerId}/password/resetLinkToken/{resetPasswordLinkToken}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomerAccountManagementV1ValidateResetPasswordLinkTokenGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomerAccountManagementV1ValidateResetPasswordLinkTokenGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}