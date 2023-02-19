package dpfm_api_processing_formatter

import (
	"context"
	"convert-to-dpfm-invoice-document-from-sap-billing-document/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-sap-billing-document/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type ProcessingFormatter struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewProcessingFormatter(ctx context.Context, db *database.Mysql, l *logger.Logger) *ProcessingFormatter {
	return &ProcessingFormatter{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (p *ProcessingFormatter) ProcessingFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *ProcessingFormatterSDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		psdc.PreparingHeaderPartner = p.PreparingHeaderPartner(sdc)
		psdc.PreparingItemPricingElement = p.PreparingItemPricingElement(sdc)

		psdc.Header, e = p.Header(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingHeader, e = p.ConversionProcessingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.Item, e = p.Item(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingItem, e = p.ConversionProcessingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ItemPricingElement, e = p.ItemPricingElement(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingItemPricingElement, e = p.ConversionProcessingItemPricingElement(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.Address, e = p.Address(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.Partner, e = p.Partner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingPartner, e = p.ConversionProcessingPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	p.l.Info(psdc)

	return nil
}

func (p *ProcessingFormatter) PreparingHeaderPartner(sdc *dpfm_api_input_reader.SDC) *PreparingHeaderPartner {
	pm := &requests.PreparingHeaderPartner{}

	for _, v := range sdc.Header.HeaderPartner {
		partnerFunction := v.PartnerFunction
		switch partnerFunction {
		case "PY":
			pm.ConvertingPayer = v.Customer
			pm.ConvertingCustomer = v.Customer
		case "SH":
			pm.ConvertingDeliverToParty = v.Customer
			pm.ConvertingSupplier = v.Supplier
		case "BP":
			pm.ConvertingBillToParty = v.Customer
			pm.ConvertingCustomer = v.Customer
		case "SP":
			pm.ConvertingSupplier = v.Supplier
		}

		if partnerFunction == "SP" || partnerFunction == "SH" || partnerFunction == "BP" || partnerFunction == "PY" {
			pm.ConvertingPartnerFunction = append(pm.ConvertingPartnerFunction, v.PartnerFunction)
		}
	}

	data := pm
	res := PreparingHeaderPartner{
		ConvertingBillToParty:     data.ConvertingBillToParty,
		ConvertingPayer:           data.ConvertingPayer,
		ConvertingDeliverToParty:  data.ConvertingDeliverToParty,
		ConvertingPartnerFunction: data.ConvertingPartnerFunction,
		ConvertingCustomer:        data.ConvertingCustomer,
		ConvertingSupplier:        data.ConvertingSupplier,
	}

	return &res
}

func (psdc *ProcessingFormatter) PreparingItemPricingElement(sdc *dpfm_api_input_reader.SDC) *PreparingItemPricingElement {
	pm := &requests.PreparingItemPricingElement{}

	for _, item := range sdc.Header.Item {
		for _, v := range item.ItemPricingElement {
			conditionType := v.ConditionType
			if conditionType != nil {
				if *conditionType == "PR00" || *conditionType == "MWST" {
					pm.ConditionType = v.ConditionType
				}
			}
		}
	}

	data := pm
	res := PreparingItemPricingElement{
		ConvertingConditionType: data.ConvertingConditionType,
	}

	return &res
}

func (p *ProcessingFormatter) Header(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*Header, error) {
	data := sdc.Header
	dataPeparingHeaderPartner := psdc.PreparingHeaderPartner

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	totalNetAmount, err := parseFloat32Ptr(data.TotalNetAmount)
	if err != nil {
		return nil, xerrors.Errorf("Parse Error: %w", err)
	}
	taxAmount, err := parseFloat32Ptr(data.TaxAmount)
	if err != nil {
		return nil, xerrors.Errorf("Parse Error: %w", err)
	}
	totalGrossAmount, err := parseFloat32Ptr(data.TotalGrossAmount)
	if err != nil {
		return nil, xerrors.Errorf("Parse Error: %w", err)
	}

	res := Header{
		ConvertingInvoiceDocument:         data.BillingDocument,
		CreationDate:                      systemDate,
		CreationTime:                      systemTime,
		LastChangeDate:                    systemDate,
		LastChangeTime:                    systemTime,
		ConvertingBillToParty:             dataPeparingHeaderPartner.ConvertingBillToParty,
		BillFromParty:                     sdc.BusinessPartnerID,
		BillFromCountry:                   data.Country,
		ConvertingPayer:                   dataPeparingHeaderPartner.ConvertingPayer,
		Payee:                             sdc.BusinessPartnerID,
		InvoiceDocumentDate:               data.BillingDocumentDate,
		TotalNetAmount:                    totalNetAmount,
		TotalTaxAmount:                    taxAmount,
		TotalGrossAmount:                  totalGrossAmount,
		TransactionCurrency:               data.TransactionCurrency,
		Incoterms:                         data.IncotermsClassification,
		PaymentTerms:                      data.CustomerPaymentTerms,
		PaymentMethod:                     data.PaymentMethod,
		HeaderIsCleared:                   getBoolPtr(false),
		HeaderPaymentBlockStatus:          getBoolPtr(false),
		HeaderPaymentRequisitionIsCreated: getBoolPtr(false),
		IsCancelled:                       getBoolPtr(false),
	}

	return &res, nil
}

func (p *ProcessingFormatter) ConversionProcessingHeader(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*ConversionProcessingHeader, error) {
	dataKey := make([]*ConversionProcessingKey, 0)

	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "BillingDocument", "InvoiceDocument", psdc.Header.ConvertingInvoiceDocument))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Customer", "BillToParty", psdc.Header.ConvertingBillToParty))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Customer", "Payer", psdc.Header.ConvertingPayer))

	dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
	if err != nil {
		return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
	}

	data, err := p.ConvertToConversionProcessingHeader(dataKey, dataQueryGets)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
	}

	return data, nil
}

func (psdc *ProcessingFormatter) ConvertToConversionProcessingHeader(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingHeader, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingHeader{}

	pm.ConvertingInvoiceDocument = data["InvoiceDocument"].CodeConvertFromString
	pm.ConvertedInvoiceDocument = data["InvoiceDocument"].CodeConvertToInt
	pm.ConvertingBillToParty = data["BillToParty"].CodeConvertFromString
	pm.ConvertedBillToParty = data["BillToParty"].CodeConvertToInt
	pm.ConvertingPayer = data["Payer"].CodeConvertFromString
	pm.ConvertedPayer = data["Payer"].CodeConvertToInt

	res := &ConversionProcessingHeader{
		ConvertingInvoiceDocument: pm.ConvertingInvoiceDocument,
		ConvertedInvoiceDocument:  pm.ConvertedInvoiceDocument,
		ConvertingBillToParty:     pm.ConvertingBillToParty,
		ConvertedBillToParty:      pm.ConvertedBillToParty,
		ConvertingPayer:           pm.ConvertingPayer,
		ConvertedPayer:            pm.ConvertedPayer,
	}

	return res, nil
}

func (p *ProcessingFormatter) Item(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Item, error) {
	res := make([]*Item, 0)
	dataProcessingFormatterHeader := psdc.Header
	dataInputReaderHeader := sdc.Header
	dataPreparingHeaderPartner := psdc.PreparingHeaderPartner
	data := sdc.Header.Item

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	for _, data := range data {
		billingQuantity, err := parseFloat32Ptr(data.BillingQuantity)
		if err != nil {
			return nil, err
		}
		netAmount, err := parseFloat32Ptr(data.NetAmount)
		if err != nil {
			return nil, err
		}
		taxAmount, err := parseFloat32Ptr(data.TaxAmount)
		if err != nil {
			return nil, err
		}
		grossAmount, err := parseFloat32Ptr(data.GrossAmount)
		if err != nil {
			return nil, err
		}

		res = append(res, &Item{
			ConvertingInvoiceDocument:       dataProcessingFormatterHeader.ConvertingInvoiceDocument,
			ConvertingInvoiceDocumentItem:   data.BillingDocumentItem,
			InvoiceDocumentItemTextBySeller: data.BillingDocumentItemText,
			ConvertingProduct:               data.Material,
			ConvertingProductGroup:          data.MaterialGroup,
			CreationDate:                    systemDate,
			CreationTime:                    systemTime,
			LastChangeDate:                  systemDate,
			LastChangeTime:                  systemTime,
			ConvertingBuyer:                 dataInputReaderHeader.SoldToParty,
			Seller:                          sdc.BusinessPartnerID,
			ConvertingDeliverToParty:        dataPreparingHeaderPartner.ConvertingDeliverToParty,
			DeliverFromParty:                sdc.BusinessPartnerID,
			InvoiceQuantity:                 billingQuantity,
			InvoiceQuantityUnit:             data.BillingQuantityUnit,
			NetAmount:                       netAmount,
			TaxAmount:                       taxAmount,
			GrossAmount:                     grossAmount,
			TransactionCurrency:             data.TransactionCurrency,
			PricingDate:                     data.PricingDate,
			ExternalReferenceDocument:       data.ExternalReferenceDocument,
			ExternalReferenceDocumentItem:   data.ExternalReferenceDocumentItem,
			ItemPaymentRequisitionIsCreated: getBoolPtr(false),
			ItemIsCleared:                   getBoolPtr(false),
			ItemPaymentBlockStatus:          getBoolPtr(false),
			IsCancelled:                     getBoolPtr(false),
		})
	}

	return res, nil
}

func (p *ProcessingFormatter) ConversionProcessingItem(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItem, error) {
	data := make([]*ConversionProcessingItem, 0)

	for _, item := range psdc.Item {
		dataKey := make([]*ConversionProcessingKey, 0)

		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "BillingDocumentItem", "InvoiceDocumentItem", item.ConvertingInvoiceDocumentItem))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Material", "Product", item.ConvertingProduct))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "MaterialGroup", "ProductGroup", item.ConvertingProductGroup))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "SoldToParty", "Buyer", item.ConvertingBuyer))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Customer", "DeliverToParty", item.ConvertingDeliverToParty))

		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
		if err != nil {
			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
		}

		datum, err := p.ConvertToConversionProcessingItem(dataKey, dataQueryGets)
		if err != nil {
			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
		}

		data = append(data, datum)
	}

	return data, nil
}

func (p *ProcessingFormatter) ConvertToConversionProcessingItem(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItem, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingItem{}

	pm.ConvertingInvoiceDocumentItem = data["InvoiceDocumentItem"].CodeConvertFromString
	pm.ConvertedInvoiceDocumentItem = data["InvoiceDocumentItem"].CodeConvertToInt
	pm.ConvertingProduct = data["Product"].CodeConvertFromString
	pm.ConvertedProduct = data["Product"].CodeConvertFromString
	pm.ConvertingProductGroup = data["ProductGroup"].CodeConvertFromString
	pm.ConvertedProductGroup = data["ProductGroup"].CodeConvertFromString
	pm.ConvertingBuyer = data["Buyer"].CodeConvertFromString
	pm.ConvertedBuyer = data["Buyer"].CodeConvertToInt
	pm.ConvertingDeliverToParty = data["DeliverToParty"].CodeConvertFromString
	pm.ConvertedDeliverToParty = data["DeliverToParty"].CodeConvertToInt

	res := &ConversionProcessingItem{
		ConvertingInvoiceDocumentItem: pm.ConvertingInvoiceDocumentItem,
		ConvertedInvoiceDocumentItem:  pm.ConvertedInvoiceDocumentItem,
		ConvertingProduct:             pm.ConvertingProduct,
		ConvertedProduct:              pm.ConvertedProduct,
		ConvertingProductGroup:        pm.ConvertingProductGroup,
		ConvertedProductGroup:         pm.ConvertedProductGroup,
		ConvertingBuyer:               pm.ConvertingBuyer,
		ConvertedBuyer:                pm.ConvertedBuyer,
		ConvertingDeliverToParty:      pm.ConvertingDeliverToParty,
		ConvertedDeliverToParty:       pm.ConvertedDeliverToParty,
	}

	return res, nil
}

func (p *ProcessingFormatter) ItemPricingElement(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ItemPricingElement, error) {
	res := make([]*ItemPricingElement, 0)
	dataHeader := psdc.Header
	dataItem := psdc.Item

	for i, dataItem := range dataItem {
		data := sdc.Header.Item[i].ItemPricingElement
		for _, data := range data {
			conditionRateValue, err := parseFloat32Ptr(data.ConditionRateValue)
			if err != nil {
				return nil, err
			}
			conditionQuantity, err := parseFloat32Ptr(data.ConditionQuantity)
			if err != nil {
				return nil, err
			}
			conditionAmount, err := parseFloat32Ptr(data.ConditionAmount)
			if err != nil {
				return nil, err
			}

			res = append(res, &ItemPricingElement{
				ConvertingInvoiceDocument:           dataHeader.ConvertingInvoiceDocument,
				ConvertingInvoiceDocumentItem:       dataItem.ConvertingInvoiceDocumentItem,
				ConvertingConditionRecord:           data.ConditionRecord,
				ConvertingConditionSequentialNumber: data.ConditionSequentialNumber,
				PricingDate:                         dataItem.PricingDate,
				ConditionRateValue:                  conditionRateValue,
				ConditionCurrency:                   data.ConditionCurrency,
				ConditionQuantity:                   conditionQuantity,
				ConditionQuantityUnit:               data.ConditionQuantityUnit,
				TaxCode:                             data.TaxCode,
				ConditionAmount:                     conditionAmount,
				TransactionCurrency:                 data.TransactionCurrency,
				ConditionIsManuallyChanged:          getBoolPtr(true),
			})
		}
	}

	return res, nil
}

func (p *ProcessingFormatter) ConversionProcessingItemPricingElement(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItemPricingElement, error) {
	data := make([]*ConversionProcessingItemPricingElement, 0)

	for _, itemPricingElement := range psdc.ItemPricingElement {
		dataKey := make([]*ConversionProcessingKey, 0)

		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ConditionRecord", "ConditionRecord", itemPricingElement.ConvertingConditionRecord))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ConditionSequentialNumber", "ConditionSequentialNumber", itemPricingElement.ConvertingConditionSequentialNumber))

		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
		if err != nil {
			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
		}

		datum, err := p.ConvertToConversionProcessingItemPricingElement(dataKey, dataQueryGets)
		if err != nil {
			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
		}

		data = append(data, datum)
	}

	return data, nil
}

func (p *ProcessingFormatter) ConvertToConversionProcessingItemPricingElement(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItemPricingElement, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingItemPricingElement{}

	pm.ConvertingConditionRecord = data["ConditionRecord"].CodeConvertFromString
	pm.ConvertedConditionRecord = data["ConditionRecord"].CodeConvertToInt
	pm.ConvertingConditionSequentialNumber = data["ConditionSequentialNumber"].CodeConvertFromString
	pm.ConvertedConditionSequentialNumber = data["ConditionSequentialNumber"].CodeConvertToInt

	res := &ConversionProcessingItemPricingElement{
		ConvertingConditionRecord:           pm.ConvertingConditionRecord,
		ConvertedConditionRecord:            pm.ConvertedConditionRecord,
		ConvertingConditionSequentialNumber: pm.ConvertingConditionSequentialNumber,
		ConvertedConditionSequentialNumber:  pm.ConvertedConditionSequentialNumber,
	}

	return res, nil
}

func (p *ProcessingFormatter) Address(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Address, error) {
	res := make([]*Address, 0)
	dataHeader := psdc.Header

	res = append(res, &Address{
		ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
	})

	return res, nil
}

func (p *ProcessingFormatter) Partner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Partner, error) {
	res := make([]*Partner, 0)
	dataHeader := psdc.Header
	dataPreparingHeaderPartner := psdc.PreparingHeaderPartner

	for _, convertingPartnerFunction := range dataPreparingHeaderPartner.ConvertingPartnerFunction {
		res = append(res, &Partner{
			ConvertingInvoiceDocument: dataHeader.ConvertingInvoiceDocument,
			ConvertingPartnerFunction: convertingPartnerFunction,
			ConvertingBusinessPartner: dataPreparingHeaderPartner.ConvertingCustomer,
		})
	}
	return res, nil
}

func (p *ProcessingFormatter) ConversionProcessingPartner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingPartner, error) {
	data := make([]*ConversionProcessingPartner, 0)

	for _, partner := range psdc.Partner {
		dataKey := make([]*ConversionProcessingKey, 0)

		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "PartnerFunction", "PartnerFunction", partner.ConvertingPartnerFunction))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Customer", "BusinessPartner", partner.ConvertingBusinessPartner))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "Supplier", "BusinessPartner", partner.ConvertingBusinessPartner))

		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
		if err != nil {
			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
		}

		datum, err := p.ConvertToConversionProcessingPartner(dataKey, dataQueryGets)
		if err != nil {
			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
		}

		data = append(data, datum)
	}

	return data, nil
}

func (p *ProcessingFormatter) ConvertToConversionProcessingPartner(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingPartner, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingPartner{}

	pm.ConvertingPartnerFunction = data["PartnerFunction"].CodeConvertFromString
	pm.ConvertedPartnerFunction = data["PartnerFunction"].CodeConvertToString
	pm.ConvertingBusinessPartner = data["BusinessPartner"].CodeConvertFromString
	pm.ConvertedBusinessPartner = data["BusinessPartner"].CodeConvertToInt

	res := &ConversionProcessingPartner{
		ConvertingPartnerFunction: pm.ConvertingPartnerFunction,
		ConvertedPartnerFunction:  pm.ConvertedPartnerFunction,
		ConvertingBusinessPartner: pm.ConvertingBusinessPartner,
		ConvertedBusinessPartner:  pm.ConvertedBusinessPartner,
	}

	return res, nil
}
