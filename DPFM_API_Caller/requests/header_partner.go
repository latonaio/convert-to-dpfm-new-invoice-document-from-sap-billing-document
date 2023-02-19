package requests

type HeaderPartner struct {
	BillingDocument string  `json:"BillingDocument"`
	PartnerFunction string  `json:"PartnerFunction"`
	Customer        *string `json:"Customer"`
	Supplier        *string `json:"Supplier"`
}
