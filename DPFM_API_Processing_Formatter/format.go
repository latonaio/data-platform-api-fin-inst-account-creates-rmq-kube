package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-fin-inst-account-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		FinInstCustomerIDByFinInst: data.FinInstCustomerIDByFinInst,
		ValidityEndDate:            data.ValidityEndDate,
		ValidityStartDate:          data.ValidityStartDate,
		IsMarkedForDeletion:        data.IsMarkedForDeletion,
	}
}

func ConvertToItemUpdates(item dpfm_api_input_reader.Item) *ItemUpdates {
	data := item

	return &ItemUpdates{
		ValidityEndDate:     data.ValidityEndDate,
		ValidityStartDate:   data.ValidityStartDate,
		FinInstControlKey:   data.FinInstControlKey,
		FinInstAccountName:  data.FinInstAccountName,
		FinInstAccount:      data.FinInstAccount,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}
