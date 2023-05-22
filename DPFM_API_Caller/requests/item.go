package requests

type Item struct {
	FinInstCountry            *string `json:"FinInstCountry"`
	FinInstCode               *string `json:"FinInstCode"`
	FinInstBranchCode         *string `json:"FinInstBranchCode"`
	FinInstFullCode           *string `json:"FinInstFullCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	ValidityEndDate           *string `json:"ValidityEndDate"`
	ValidityStartDate         *string `json:"ValidityStartDate"`
	FinInstControlKey         *string `json:"FinInstControlKey"`
	FinInstAccountName        *string `json:"FinInstAccountName"`
	FinInstAccount            *string `json:"FinInstAccount"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
