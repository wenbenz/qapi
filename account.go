package qapi

// Account associated with the user to whom the API client is authorized.
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts
type Account struct {
	// Type of the account (e.g., "Cash", "Margin").
	Type string `json:"type"`

	// Eight-digit account number (e.g., "26598145")
	// Stored as a string, it's used for making account-related API calls
	Number string `json:"number"`

	// Status of the account (e.g., Active).
	Status string `json:"status"`

	// Whether this is a primary account for the holder.
	IsPrimary bool `json:"isPrimary"`

	// Whether this account is one that gets billed for various expenses such as inactivity fees, market data, etc.
	IsBilling bool `json:"isBilling"`

	// Type of client holding the account (e.g., "Individual").
	ClientAccountType string `json:"clientAccountType"`
}
