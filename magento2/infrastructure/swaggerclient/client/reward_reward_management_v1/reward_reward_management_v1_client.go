// Code generated by go-swagger; DO NOT EDIT.

package reward_reward_management_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new reward reward management v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for reward reward management v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RewardRewardManagementV1SetPost Set reward points to quote
*/
func (a *Client) RewardRewardManagementV1SetPost(params *RewardRewardManagementV1SetPostParams) (*RewardRewardManagementV1SetPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRewardRewardManagementV1SetPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rewardRewardManagementV1SetPost",
		Method:             "POST",
		PathPattern:        "/V1/reward/mine/use-reward",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RewardRewardManagementV1SetPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RewardRewardManagementV1SetPostOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}