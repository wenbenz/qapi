package qapi

// Balance belonging to an Account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-balances
type Balance struct {

	// Currency of the balance figure(e.g., "USD" or "CAD").
	Currency string `json:"currency"`

	// Balance amount.
	Cash float32 `json:"cash"`

	// Market value of all securities in the account in a given currency.
	MarketValue float32 `json:"marketValue"`

	// Equity as a difference between cash and marketValue properties.
	TotalEquity float32 `json:"totalEquity"`

	// Buying power for that particular currency side of the account.
	BuyingPower float32 `json:"buyingPower"`

	// Maintenance excess for that particular side of the account.
	MaintenanceExcess float32 `json:"maintenanceExcess"`

	// Whether real-time data was used to calculate the above values.
	IsRealTime bool `json:"isRealTime"`
}

// AccountBalances represents per-currency and combined balances for a specified account.
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-balances
type AccountBalances struct {
	PerCurrencyBalances    []Balance `json:"perCurrencyBalances"`
	CombinedBalances       []Balance `json:"combinedBalances"`
	SODPerCurrencyBalances []Balance `json:"sodPerCurrencyBalances"`
	SODCombinedBalances    []Balance `json:"sodCombinedBalances"`
}
