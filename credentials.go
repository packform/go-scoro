package scoro

// Credentials holds authentication parameters used to identificate Scoro customer
// and authorize requested action.
type Credentials struct {
	// ApiKey holds apiKey value that is listed in Settings > External Connections > API.
	ApiKey string `json:"apiKey"`

	// CompanyID holds company_account_id value that is listed in Settings > External Connections > API.
	CompanyID string `json:"company_account_id"`

	// API subdomain
	Subdomain string `json:"-"`
}
