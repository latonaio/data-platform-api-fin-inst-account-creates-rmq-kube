package sub_func_complementer

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	APIType             string   `json:"api_type"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Header *Header `json:"Header"`
	Item   *[]Item `json:"Item"`
}

type Header struct {
	FinInstCountry             *string `json:"FinInstCountry"`
	FinInstCode                *string `json:"FinInstCode"`
	FinInstBusinessPartner     *int    `json:"FinInstBusinessPartner"`
	InternalFinInstCustomerID  *int    `json:"InternalFinInstCustomerID"`
	AccountBusinessPartner     *int    `json:"AccountBusinessPartner"`
	FinInstCustomerIDByFinInst *string `json:"FinInstCustomerIDByFinInst"`
	ValidityEndDate            *string `json:"ValidityEndDate"`
	ValidityStartDate          *string `json:"ValidityStartDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

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
