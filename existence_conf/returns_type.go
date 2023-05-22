package existence_conf

type Returns struct {
	ConnectionKey       string               `json:"connection_key"`
	Result              bool                 `json:"result"`
	RedisKey            string               `json:"redis_key"`
	RuntimeSessionID    string               `json:"runtime_session_id"`
	BusinessPartnerID   *int                 `json:"business_partner"`
	Filepath            string               `json:"filepath"`
	ServiceLabel        string               `json:"service_label"`
	BPGeneralReturn     BPGeneralReturn      `json:"BusinessPartnerGeneral"`
	FinInstMasterReturn *FinInstMasterReturn `json:"FinInstMasterReturn,omitempty"`
	CountryReturn       CountryReturn        `json:"Country,omitempty"`
	APISchema           string               `json:"api_schema"`
	Accepter            []string             `json:"accepter"`
	Deleted             bool                 `json:"deleted"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type FinInstMasterReturn struct {
	FinInstCountry string `json:"FinInstCountry"`
	FinInstCode    string `json:"FinInstCode"`
}

type CountryReturn struct {
	Country string `json:"Country"`
}
