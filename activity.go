package qapi

import (
	"time"
)

// Copied from https://github.com/paudley/qapi

// Activities in an Account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-activities
type Activities struct {
	List []Activity `json:"activities"`
}

// Activity belonging to an Account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-activities
type Activity struct {
	// Trade timestamp.
	TradeDate time.Time `json:"tradeDate"`

	// Transaction timestamp.
	TransactionDate time.Time `json:"transactionDate"`

	// Expected settlement timestamp.
	SettlementDate time.Time `json:"settlementDate"`

	// Action code (ie: DEP, CON, BUY, SELL, TFI).
	Action string `json:"action"`

	// Activity symbol (if activity is related to a security).
	Symbol string `json:"symbol"`

	// Internal symbol identifier.
	SymbolID int `json:"symbolId"`

	// Notes field describing the activity.
	Description string `json:"description"`

	// Currency of activity (ie: CAD, USD).
	Currency string `json:"currency"`

	// Activity quantity (number of shares, etc.).
	Quantity float64 `json:"quantity"`

	// Price paid.
	Price float64 `json:"price"`

	// Gross amount.
	GrossAmount float64 `json:"grossAmount"`

	// Total of any commission paid.
	Commission float64 `json:"commission"`

	// Net amount.
	NetAmount float64 `json:"netAmount"`

	// Text description of the action (ie: "Transfers", "Interest").
	Type string `json:"type"`
}
