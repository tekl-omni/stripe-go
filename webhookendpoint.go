package stripe

import "encoding/json"

// WebhookEndpointParams is the set of parameters that can be used when creating a webhook endpoint.
// For more details see https://stripe.com/docs/api#create_webhook_endpoint.
type WebhookEndpointParams struct {
	Params        `form:"*" json:"*"`
	Connect       *bool     `form:"connect" json:"connect"`
	Disabled      *bool     `form:"disabled" json:"disabled"`
	EnabledEvents []*string `form:"enabled_events" json:"enabled_events"`
	URL           *string   `form:"url" json:"url"`

	// This parameter is only available on creation.
	// We recommend setting the API version that the library is pinned to. See apiversion in stripe.go
	APIVersion *string `form:"api_version" json:"api_version"`
}

// WebhookEndpointListParams is the set of parameters that can be used when listing webhook endpoints.
// For more detail see https://stripe.com/docs/api#list_webhook_endpoints.
type WebhookEndpointListParams struct {
	ListParams   `form:"*" json:"*"`
	Created      *int64            `form:"created" json:"created"`
	CreatedRange *RangeQueryParams `form:"created" json:"created"`
}

// WebhookEndpoint is the resource representing a Stripe webhook endpoint.
// For more details see https://stripe.com/docs/api#webhook_endpoints.
type WebhookEndpoint struct {
	APIVersion    string   `json:"api_version"`
	Application   string   `json:"application"`
	Connect       bool     `json:"connect"`
	Created       int64    `json:"created"`
	Deleted       bool     `json:"deleted"`
	EnabledEvents []string `json:"enabled_events"`
	ID            string   `json:"id"`
	Livemode      bool     `json:"livemode"`
	Object        string   `json:"object"`
	Secret        string   `json:"secret"`
	Status        string   `json:"status"`
	URL           string   `json:"url"`
}

// WebhookEndpointList is a list of webhook endpoints as retrieved from a list endpoint.
type WebhookEndpointList struct {
	ListMeta
	Data []*WebhookEndpoint `json:"data"`
}

// UnmarshalJSON handles deserialization of a WebhookEndpoint.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *WebhookEndpoint) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type endpoint WebhookEndpoint
	var v endpoint
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = WebhookEndpoint(v)
	return nil
}
