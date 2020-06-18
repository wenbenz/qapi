package qapi

// Balance belonging to an Account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-balances
type Balance struct {

	// Currency of the balance figure(e.g., "USD" or "CAD").
	Currency string `json:"currency"`

	// Balance amount.
	Cash float64 `json:"cash"`

	// Market value of all securities in the account in a given currency.
	MarketValue float64 `json:"marketValue"`

	// Equity as a difference between cash and marketValue properties.
	TotalEquity float64 `json:"totalEquity"`

	// Buying power for that particular currency side of the account.
	BuyingPower float64 `json:"buyingPower"`

	// Maintenance excess for that particular side of the account.
	MaintenanceExcess float64 `json:"maintenanceExcess"`

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
