package qapi

// Copied from https://github.com/paudley/qapi

// Position belonging to an account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-positions
type Position struct {
	// Position symbol.
	Symbol string `json:"symbol"`

	// Internal symbol identifier
	SymbolID int `json:"symbolId"`

	// Position quantity remaining open.
	OpenQuantity float64 `json:"openQuantity"`

	// Portion of the position that was closed today.
	ClosedQuantity float64 `json:"closedQuantity"`

	// Market value of the position (quantity x price).
	CurrentMarketValue float64 `json:"currentMarketValue"`

	// Current price of the position symbol.
	CurrentPrice float64 `json:"currentPrice"`

	// Average price paid for all executions constituting the position.
	AverageEntryPrice float64 `json:"averageEntryPrice"`

	// Realized profit/loss on this position.
	ClosedPnL float64 `json:"closedPnL"`

	// Unrealized profit/loss on this position.
	OpenPnL float64 `json:"openPnL"`

	// Total cost of the position.
	TotalCost float64 `json:"totalCost"`

	// Designates whether real-time quote was used to compute PnL.
	IsRealTime bool `json:"isRealTime"`

	// Designates whether a symbol is currently undergoing a reorg.
	IsUnderReorg bool `json:"isUnderReorg"`
}

// Positions belonging to an account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-positions
type Positions struct {
	Positions []Position `json:"positions"`
}
