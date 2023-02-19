package requests

type PreparingHeaderPartner struct {
	PartnerFunction           string   `json:"PartnerFunction"`
	Customer                  *string  `json:"Customer"`
	Supplier                  *string  `json:"Supplier"`
	ConvertingBillToParty     *string  `json:"ConvertingBillToParty"`
	ConvertingPayer           *string  `json:"ConvertingPayer"`
	ConvertingDeliverToParty  *string  `json:"ConvertingDeliverToParty"`
	ConvertingPartnerFunction []string `json:"ConvertingPartnerFunction"`
	ConvertingCustomer        *string  `json:"ConvertingCustomer"`
	ConvertingSupplier        *string  `json:"ConvertingSupplier"`
}
