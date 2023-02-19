package requests

type ConversionProcessingHeader struct {
	ConvertingInvoiceDocument *string `json:"ConvertingInvoiceDocument"`
	ConvertedInvoiceDocument  *int    `json:"ConvertedInvoiceDocument"`
	ConvertingBillToParty     *string `json:"ConvertingBillToParty"`
	ConvertedBillToParty      *int    `json:"ConvertedBillToParty"`
	ConvertingPayer           *string `json:"ConvertingPayer"`
	ConvertedPayer            *int    `json:"ConvertedPayer"`
}
