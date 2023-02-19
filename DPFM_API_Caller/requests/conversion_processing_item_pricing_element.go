package requests

type ConversionProcessingItemPricingElement struct {
	ConvertingConditionRecord           *string `json:"ConvertingConditionRecord"`
	ConvertedConditionRecord            *int    `json:"ConvertedConditionRecord"`
	ConvertingConditionSequentialNumber *string `json:"ConvertingConditionSequentialNumber"`
	ConvertedConditionSequentialNumber  *int    `json:"ConvertedConditionSequentialNumber"`
}
