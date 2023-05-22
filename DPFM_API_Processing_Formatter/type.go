package dpfm_api_processing_formatter

type Message struct {
	HeaderUpdates *HeaderUpdates `json:"HeaderUpdates"`
}

type HeaderUpdates struct {
	FinInstCustomerIDByFinInst *string `json:"FinInstCustomerIDByFinInst"`
	ValidityEndDate            *string `json:"ValidityEndDate"`
	ValidityStartDate          *string `json:"ValidityStartDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

type ItemUpdates struct {
	ValidityEndDate     *string `json:"ValidityEndDate"`
	ValidityStartDate   *string `json:"ValidityStartDate"`
	FinInstControlKey   *string `json:"FinInstControlKey"`
	FinInstAccountName  *string `json:"FinInstAccountName"`
	FinInstAccount      *string `json:"FinInstAccount"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
