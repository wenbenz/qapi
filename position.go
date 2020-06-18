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
	OpenQuantity float32 `json:"openQuantity"`

	// Portion of the position that was closed today.
	ClosedQuantity float32 `json:"closedQuantity"`

	// Market value of the position (quantity x price).
	CurrentMarketValue float32 `json:"currentMarketValue"`

	// Current price of the position symbol.
	CurrentPrice float32 `json:"currentPrice"`

	// Average price paid for all executions constituting the position.
	AverageEntryPrice float32 `json:"averageEntryPrice"`

	// Realized profit/loss on this position.
	ClosedPnL float32 `json:"closedPnL"`

	// Unrealized profit/loss on this position.
	OpenPnL float32 `json:"openPnL"`

	// Total cost of the position.
	TotalCost float32 `json:"totalCost"`

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
