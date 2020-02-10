package stripe

import "encoding/json"

// InvoiceItemParams is the set of parameters that can be used when creating or updating an invoice item.
// For more details see https://stripe.com/docs/api#create_invoiceitem and https://stripe.com/docs/api#update_invoiceitem.
type InvoiceItemParams struct {
	Params            `form:"*" json:"*"`
	Amount            *int64                   `form:"amount" json:"amount"`
	Currency          *string                  `form:"currency" json:"currency"`
	Customer          *string                  `form:"customer" json:"customer"`
	Description       *string                  `form:"description" json:"description"`
	Discountable      *bool                    `form:"discountable" json:"discountable"`
	Invoice           *string                  `form:"invoice" json:"invoice"`
	Period            *InvoiceItemPeriodParams `form:"period" json:"period"`
	Quantity          *int64                   `form:"quantity" json:"quantity"`
	Subscription      *string                  `form:"subscription" json:"subscription"`
	TaxRates          []*string                `form:"tax_rates" json:"tax_rates"`
	UnitAmount        *int64                   `form:"unit_amount" json:"unit_amount"`
	UnitAmountDecimal *float64                 `form:"unit_amount_decimal,high_precision" json:"unit_amount_decimal,high_precision"`
}

// InvoiceItemPeriodParams represents the period associated with that invoice item.
type InvoiceItemPeriodParams struct {
	End   *int64 `form:"end" json:"end"`
	Start *int64 `form:"start" json:"start"`
}

// InvoiceItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type InvoiceItemListParams struct {
	ListParams   `form:"*" json:"*"`
	Created      *int64            `form:"created" json:"created"`
	CreatedRange *RangeQueryParams `form:"created" json:"created"`
	Customer     *string           `form:"customer" json:"customer"`
	Invoice      *string           `form:"invoice" json:"invoice"`
	Pending      *bool             `form:"pending" json:"pending"`
}

// InvoiceItem is the resource represneting a Stripe invoice item.
// For more details see https://stripe.com/docs/api#invoiceitems.
type InvoiceItem struct {
	Amount            int64             `json:"amount"`
	Currency          Currency          `json:"currency"`
	Customer          *Customer         `json:"customer"`
	Date              int64             `json:"date"`
	Deleted           bool              `json:"deleted"`
	Description       string            `json:"description"`
	Discountable      bool              `json:"discountable"`
	ID                string            `json:"id"`
	Invoice           *Invoice          `json:"invoice"`
	Livemode          bool              `json:"livemode"`
	Metadata          map[string]string `json:"metadata"`
	Period            *Period           `json:"period"`
	Plan              *Plan             `json:"plan"`
	Proration         bool              `json:"proration"`
	Quantity          int64             `json:"quantity"`
	Subscription      *Subscription     `json:"subscription"`
	TaxRates          []*TaxRate        `json:"tax_rates"`
	UnitAmount        int64             `json:"unit_amount"`
	UnitAmountDecimal float64           `json:"unit_amount_decimal,string"`
}

// InvoiceItemList is a list of invoice items as retrieved from a list endpoint.
type InvoiceItemList struct {
	ListMeta
	Data []*InvoiceItem `json:"data"`
}

// UnmarshalJSON handles deserialization of an InvoiceItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoiceItem InvoiceItem
	var v invoiceItem
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = InvoiceItem(v)
	return nil
}
