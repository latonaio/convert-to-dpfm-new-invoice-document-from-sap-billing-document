package requests

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
