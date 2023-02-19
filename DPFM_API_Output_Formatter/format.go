package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-sap-billing-document/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-invoice-document-from-sap-billing-document/DPFM_API_Processing_Formatter"
)

func OutputFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_formatter.ProcessingFormatterSDC,
	osdc *Output,
) error {
	header := ConvertToHeader(*sdc, *psdc)
	item := ConvertToItem(*sdc, *psdc)
	itemPricingElement := ConvertToItemPricingElement(*sdc, *psdc)
	address := ConvertToAddress(*sdc, *psdc)
	partner := ConvertToPartner(*sdc, *psdc)

	osdc.Message = Message{
		Header:             header,
		Item:               item,
		ItemPricingElement: itemPricingElement,
		Address:            address,
		Partner:            partner,
	}

	return nil
}

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) *Header {
	dataProcessingHeader := psdc.Header
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	header := &Header{
		InvoiceDocument:                   *dataConversionProcessingHeader.ConvertedInvoiceDocument,
		CreationDate:                      dataProcessingHeader.CreationDate,
		CreationTime:                      dataProcessingHeader.CreationTime,
		LastChangeDate:                    dataProcessingHeader.LastChangeDate,
		LastChangeTime:                    dataProcessingHeader.LastChangeTime,
		BillToParty:                       dataConversionProcessingHeader.ConvertedBillToParty,
		BillFromParty:                     dataProcessingHeader.BillFromParty,
		BillFromCountry:                   dataProcessingHeader.BillFromCountry,
		Payer:                             dataConversionProcessingHeader.ConvertedPayer,
		Payee:                             dataProcessingHeader.Payee,
		InvoiceDocumentDate:               dataProcessingHeader.InvoiceDocumentDate,
		TotalNetAmount:                    dataProcessingHeader.TotalNetAmount,
		TotalTaxAmount:                    dataProcessingHeader.TotalTaxAmount,
		TotalGrossAmount:                  dataProcessingHeader.TotalGrossAmount,
		TransactionCurrency:               dataProcessingHeader.TransactionCurrency,
		Incoterms:                         dataProcessingHeader.Incoterms,
		PaymentTerms:                      dataProcessingHeader.PaymentTerms,
		PaymentMethod:                     dataProcessingHeader.PaymentMethod,
		HeaderIsCleared:                   dataProcessingHeader.HeaderIsCleared,
		HeaderPaymentBlockStatus:          dataProcessingHeader.HeaderPaymentBlockStatus,
		HeaderPaymentRequisitionIsCreated: dataProcessingHeader.HeaderPaymentRequisitionIsCreated,
		IsCancelled:                       dataProcessingHeader.IsCancelled,
	}

	return header
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Item {
	dataProcessingItem := psdc.Item
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	dataConversionProcessingItem := psdc.ConversionProcessingItem

	items := make([]*Item, 0)
	for i := range dataProcessingItem {
		item := &Item{
			InvoiceDocument:                 *dataConversionProcessingHeader.ConvertedInvoiceDocument,
			InvoiceDocumentItem:             *dataConversionProcessingItem[i].ConvertedInvoiceDocumentItem,
			InvoiceDocumentItemTextBySeller: dataProcessingItem[i].InvoiceDocumentItemTextBySeller,
			Product:                         dataConversionProcessingItem[i].ConvertedProduct,
			ProductGroup:                    dataConversionProcessingItem[i].ConvertedProductGroup,
			CreationDate:                    dataProcessingItem[i].CreationDate,
			CreationTime:                    dataProcessingItem[i].CreationTime,
			LastChangeDate:                  dataProcessingItem[i].LastChangeDate,
			LastChangeTime:                  dataProcessingItem[i].LastChangeTime,
			Buyer:                           dataConversionProcessingItem[i].ConvertedBuyer,
			Seller:                          dataProcessingItem[i].Seller,
			DeliverToParty:                  dataConversionProcessingItem[i].ConvertedDeliverToParty,
			DeliverFromParty:                dataProcessingItem[i].DeliverFromParty,
			InvoiceQuantity:                 dataProcessingItem[i].InvoiceQuantity,
			InvoiceQuantityUnit:             dataProcessingItem[i].InvoiceQuantityUnit,
			NetAmount:                       dataProcessingItem[i].NetAmount,
			TaxAmount:                       dataProcessingItem[i].TaxAmount,
			GrossAmount:                     dataProcessingItem[i].GrossAmount,
			TransactionCurrency:             dataProcessingItem[i].TransactionCurrency,
			PricingDate:                     dataProcessingItem[i].PricingDate,
			ExternalReferenceDocument:       dataProcessingItem[i].ExternalReferenceDocument,
			ExternalReferenceDocumentItem:   dataProcessingItem[i].ExternalReferenceDocumentItem,
			ItemPaymentRequisitionIsCreated: dataProcessingItem[i].ItemPaymentRequisitionIsCreated,
			ItemIsCleared:                   dataProcessingItem[i].ItemIsCleared,
			ItemPaymentBlockStatus:          dataProcessingItem[i].ItemPaymentBlockStatus,
			IsCancelled:                     dataProcessingItem[i].IsCancelled,
		}

		items = append(items, item)
	}

	return items
}

func ConvertToItemPricingElement(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*ItemPricingElement {
	dataProcessingItemPricingElement := psdc.ItemPricingElement
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	dataConversionProcessingItem := psdc.ConversionProcessingItem
	dataConversionProcessingItemPricingElement := psdc.ConversionProcessingItemPricingElement

	dataConversionProcessingItemMap := make(map[string]*dpfm_api_processing_formatter.ConversionProcessingItem, len(dataConversionProcessingItem))
	for _, v := range dataConversionProcessingItem {
		dataConversionProcessingItemMap[*v.ConvertingInvoiceDocumentItem] = v
	}

	itemPricingElements := make([]*ItemPricingElement, 0)
	for i, v := range dataProcessingItemPricingElement {
		if _, ok := dataConversionProcessingItemMap[v.ConvertingInvoiceDocumentItem]; !ok {
			continue
		}

		itemPricingElements = append(itemPricingElements, &ItemPricingElement{
			InvoiceDocument:            *dataConversionProcessingHeader.ConvertedInvoiceDocument,
			InvoiceDocumentItem:        *dataConversionProcessingItemMap[v.ConvertingInvoiceDocumentItem].ConvertedInvoiceDocumentItem,
			ConditionRecord:            dataConversionProcessingItemPricingElement[i].ConvertedConditionRecord,
			ConditionSequentialNumber:  dataConversionProcessingItemPricingElement[i].ConvertedConditionSequentialNumber,
			PricingDate:                dataProcessingItemPricingElement[i].PricingDate,
			ConditionRateValue:         dataProcessingItemPricingElement[i].ConditionRateValue,
			ConditionCurrency:          dataProcessingItemPricingElement[i].ConditionCurrency,
			ConditionQuantity:          dataProcessingItemPricingElement[i].ConditionQuantity,
			ConditionQuantityUnit:      dataProcessingItemPricingElement[i].ConditionQuantityUnit,
			TaxCode:                    dataProcessingItemPricingElement[i].TaxCode,
			ConditionAmount:            dataProcessingItemPricingElement[i].ConditionAmount,
			TransactionCurrency:        dataProcessingItemPricingElement[i].TransactionCurrency,
			ConditionIsManuallyChanged: dataProcessingItemPricingElement[i].ConditionIsManuallyChanged,
		})
	}

	return itemPricingElements
}

func ConvertToAddress(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Address {
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	addresses := make([]*Address, 0)
	addresses = append(addresses, &Address{
		InvoiceDocument: *dataConversionProcessingHeader.ConvertedInvoiceDocument,
	})

	return addresses
}

func ConvertToPartner(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Partner {
	dataProcessingPartner := psdc.Partner
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	dataConversionProcessingPartner := psdc.ConversionProcessingPartner

	partners := make([]*Partner, 0)
	for i := range dataProcessingPartner {
		partners = append(partners, &Partner{
			InvoiceDocument: *dataConversionProcessingHeader.ConvertedInvoiceDocument,
			PartnerFunction: *dataConversionProcessingPartner[i].ConvertedPartnerFunction,
			BusinessPartner: *dataConversionProcessingPartner[i].ConvertedBusinessPartner,
		})
	}

	return partners
}
