package dpfm_api_processing_formatter

type ProcessingFormatterSDC struct {
	PreparingHeaderPartner                 *PreparingHeaderPartner                   `json:"PreparingHeaderPartner"`
	PreparingItemPricingElement            *PreparingItemPricingElement              `json:"PreparingItemPricingElement"`
	Header                                 *Header                                   `json:"Header"`
	ConversionProcessingHeader             *ConversionProcessingHeader               `json:"ConversionProcessingHeader"`
	Item                                   []*Item                                   `json:"Item"`
	ConversionProcessingItem               []*ConversionProcessingItem               `json:"ConversionProcessingItem"`
	ItemPricingElement                     []*ItemPricingElement                     `json:"ItemPricingElement"`
	ConversionProcessingItemPricingElement []*ConversionProcessingItemPricingElement `json:"ConversionProcessingItemPricingElement"`
	Address                                []*Address                                `json:"Address"`
	Partner                                []*Partner                                `json:"Partner"`
	ConversionProcessingPartner            []*ConversionProcessingPartner            `json:"ConversionProcessingPartner"`
}

type ConversionProcessingKey struct {
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type ConversionProcessingCommonQueryGets struct {
	CodeConversionID      int      `json:"CodeConversionID"`
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	CodeConvertToInt      *int     `json:"CodeConvertToInt"`
	CodeConvertToFloat    *float32 `json:"CodeConvertToFloat"`
	CodeConvertToString   *string  `json:"CodeConvertToString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

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

type PreparingItemPricingElement struct {
	ConditionType           *string `json:"ConditionType"`
	ConvertingConditionType *string `json:"ConvertingConditionType"`
}

type Header struct {
	ConvertingInvoiceDocument         string   `json:"ConvertingInvoiceDocument"`
	CreationDate                      *string  `json:"CreationDate"`
	CreationTime                      *string  `json:"CreationTime"`
	LastChangeDate                    *string  `json:"LastChangeDate"`
	LastChangeTime                    *string  `json:"LastChangeTime"`
	ConvertingBillToParty             *string  `json:"ConvertingBillToParty"`
	BillFromParty                     *int     `json:"BillFromParty"`
	BillFromCountry                   *string  `json:"BillFromCountry"`
	ConvertingPayer                   *string  `json:"ConvertingPayer"`
	Payee                             *int     `json:"Payee"`
	InvoiceDocumentDate               *string  `json:"InvoiceDocumentDate"`
	TotalNetAmount                    *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32 `json:"TotalGrossAmount"`
	TransactionCurrency               *string  `json:"TransactionCurrency"`
	Incoterms                         *string  `json:"Incoterms"`
	PaymentTerms                      *string  `json:"PaymentTerms"`
	PaymentMethod                     *string  `json:"PaymentMethod"`
	HeaderIsCleared                   *bool    `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool    `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool    `json:"HeaderPaymentRequisitionIsCreated"`
	IsCancelled                       *bool    `json:"IsCancelled"`
}

type ConversionProcessingHeader struct {
	ConvertingInvoiceDocument *string `json:"ConvertingInvoiceDocument"`
	ConvertedInvoiceDocument  *int    `json:"ConvertedInvoiceDocument"`
	ConvertingBillToParty     *string `json:"ConvertingBillToParty"`
	ConvertedBillToParty      *int    `json:"ConvertedBillToParty"`
	ConvertingPayer           *string `json:"ConvertingPayer"`
	ConvertedPayer            *int    `json:"ConvertedPayer"`
}

type Item struct {
	ConvertingInvoiceDocument       string   `json:"ConvertingInvoiceDocument"`
	ConvertingInvoiceDocumentItem   string   `json:"ConvertingInvoiceDocumentItem"`
	InvoiceDocumentItemTextBySeller *string  `json:"InvoiceDocumentItemTextBySeller"`
	ConvertingProduct               *string  `json:"ConvertingProduct"`
	ConvertingProductGroup          *string  `json:"ConvertingProductGroup"`
	CreationDate                    *string  `json:"CreationDate"`
	CreationTime                    *string  `json:"CreationTime"`
	LastChangeDate                  *string  `json:"LastChangeDate"`
	LastChangeTime                  *string  `json:"LastChangeTime"`
	ConvertingBuyer                 *string  `json:"ConvertingBuyer"`
	Seller                          *int     `json:"Seller"`
	ConvertingDeliverToParty        *string  `json:"ConvertingDeliverToParty"`
	DeliverFromParty                *int     `json:"DeliverFromParty"`
	InvoiceQuantity                 *float32 `json:"InvoiceQuantity"`
	InvoiceQuantityUnit             *string  `json:"InvoiceQuantityUnit"`
	NetAmount                       *float32 `json:"NetAmount"`
	TaxAmount                       *float32 `json:"TaxAmount"`
	GrossAmount                     *float32 `json:"GrossAmount"`
	TransactionCurrency             *string  `json:"TransactionCurrency"`
	PricingDate                     *string  `json:"PricingDate"`
	ExternalReferenceDocument       *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem   *string  `json:"ExternalReferenceDocumentItem"`
	ItemPaymentRequisitionIsCreated *bool    `json:"ItemPaymentRequisitionIsCreated"`
	ItemIsCleared                   *bool    `json:"ItemIsCleared"`
	ItemPaymentBlockStatus          *bool    `json:"ItemPaymentBlockStatus"`
	IsCancelled                     *bool    `json:"IsCancelled"`
}

type ConversionProcessingItem struct {
	ConvertingInvoiceDocumentItem *string `json:"ConvertingInvoiceDocumentItem"`
	ConvertedInvoiceDocumentItem  *int    `json:"ConvertedInvoiceDocumentItem"`
	ConvertingProduct             *string `json:"ConvertingProduct"`
	ConvertedProduct              *string `json:"ConvertedProduct"`
	ConvertingProductGroup        *string `json:"ConvertingProductGroup"`
	ConvertedProductGroup         *string `json:"ConvertedProductGroup"`
	ConvertingBuyer               *string `json:"ConvertingBuyer"`
	ConvertedBuyer                *int    `json:"ConvertedBuyer"`
	ConvertingDeliverToParty      *string `json:"ConvertingDeliverToParty"`
	ConvertedDeliverToParty       *int    `json:"ConvertedDeliverToParty"`
}

type ItemPricingElement struct {
	ConvertingInvoiceDocument           string   `json:"ConvertingInvoiceDocument"`
	ConvertingInvoiceDocumentItem       string   `json:"ConvertingOrderItem"`
	ConvertingConditionRecord           *string  `json:"ConvertingConditionRecord"`
	ConvertingConditionSequentialNumber *string  `json:"ConvertingConditionSequentialNumber"`
	PricingDate                         *string  `json:"PricingDate"`
	ConditionRateValue                  *float32 `json:"ConditionRateValue"`
	ConditionCurrency                   *string  `json:"ConditionCurrency"`
	ConditionQuantity                   *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit               *string  `json:"ConditionQuantityUnit"`
	TaxCode                             *string  `json:"TaxCode"`
	ConditionAmount                     *float32 `json:"ConditionAmount"`
	TransactionCurrency                 *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged          *bool    `json:"ConditionIsManuallyChanged"`
}

type ConversionProcessingItemPricingElement struct {
	ConvertingConditionRecord           *string `json:"ConvertingConditionRecord"`
	ConvertedConditionRecord            *int    `json:"ConvertedConditionRecord"`
	ConvertingConditionSequentialNumber *string `json:"ConvertingConditionSequentialNumber"`
	ConvertedConditionSequentialNumber  *int    `json:"ConvertedConditionSequentialNumber"`
}

type Address struct {
	ConvertingInvoiceDocument string `json:"ConvertingInvoiceDocument"`
}

type Partner struct {
	ConvertingInvoiceDocument string  `json:"ConvertingInvoiceDocument"`
	ConvertingPartnerFunction string  `json:"ConvertingPartnerFunction"`
	ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
	ExternalDocumentID        *string `json:"ExternalDocumentID"`
}

type ConversionProcessingPartner struct {
	ConvertingPartnerFunction *string `json:"ConvertingPartnerFunction"`
	ConvertedPartnerFunction  *string `json:"ConvertedPartnerFunction"`
	ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
	ConvertedBusinessPartner  *int    `json:"ConvertedBusinessPartner"`
}
