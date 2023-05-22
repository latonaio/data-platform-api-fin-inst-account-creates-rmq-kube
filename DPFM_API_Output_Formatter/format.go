package dpfm_api_output_formatter

import (
	dpfm_api_processing_formatter "data-platform-api-fin-inst-account-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-fin-inst-account-creates-rmq-kube/sub_func_complementer"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) *Header {
	data := subfuncSDC.Message.Header

	header := &Header{
		FinInstCountry:             data.FinInstCountry,
		FinInstCode:                data.FinInstCode,
		FinInstBusinessPartner:     data.FinInstBusinessPartner,
		InternalFinInstCustomerID:  data.InternalFinInstCustomerID,
		AccountBusinessPartner:     data.AccountBusinessPartner,
		FinInstCustomerIDByFinInst: data.FinInstCustomerIDByFinInst,
		ValidityEndDate:            data.ValidityEndDate,
		ValidityStartDate:          data.ValidityStartDate,
		IsMarkedForDeletion:        data.IsMarkedForDeletion,
	}

	return header

}

func ConvertToHeaderUpdates(headerUpdates *dpfm_api_processing_formatter.HeaderUpdates) *Header {
	data := headerUpdates

	header := &Header{
		FinInstCustomerIDByFinInst: data.FinInstCustomerIDByFinInst,
		ValidityEndDate:            data.ValidityEndDate,
		ValidityStartDate:          data.ValidityStartDate,
		IsMarkedForDeletion:        data.IsMarkedForDeletion,
	}

	return header
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) *[]Item {
	var item []Item

	for _, data := range *subfuncSDC.Message.Item {
		item = append(item, Item{
			FinInstCountry:            data.FinInstCountry,
			FinInstCode:               data.FinInstCode,
			FinInstBranchCode:         data.FinInstBranchCode,
			FinInstFullCode:           data.FinInstFullCode,
			InternalFinInstCustomerID: data.InternalFinInstCustomerID,
			InternalFinInstAccountID:  data.InternalFinInstAccountID,
			ValidityEndDate:           data.ValidityEndDate,
			ValidityStartDate:         data.ValidityStartDate,
			FinInstControlKey:         data.FinInstControlKey,
			FinInstAccountName:        data.FinInstAccountName,
			FinInstAccount:            data.FinInstAccount,
			IsMarkedForDeletion:       data.IsMarkedForDeletion,
		})
	}

	return &item
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) *[]Item {
	var item []Item

	for _, data := range *itemUpdates {
		item = append(item, Item{
			ValidityEndDate:     data.ValidityEndDate,
			ValidityStartDate:   data.ValidityStartDate,
			FinInstControlKey:   data.FinInstControlKey,
			FinInstAccountName:  data.FinInstAccountName,
			FinInstAccount:      data.FinInstAccount,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	return &item
}
