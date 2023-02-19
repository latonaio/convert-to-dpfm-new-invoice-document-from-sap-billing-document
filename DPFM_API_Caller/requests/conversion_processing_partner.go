package requests

type ConversionProcessingPartner struct {
	ConvertingPartnerFunction *string `json:"ConvertingPartnerFunction"`
	ConvertedPartnerFunction  *string `json:"ConvertedPartnerFunction"`
	ConvertingBusinessPartner *string `json:"ConvertingBusinessPartner"`
	ConvertedBusinessPartner  *int    `json:"ConvertedBusinessPartner"`
}
