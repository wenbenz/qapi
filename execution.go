package qapi

import "time"

// Execution belonging to an Account
//
// Ref: http://www.questrade.com/api/documentation/rest-operations/account-calls/accounts-id-executions
type Execution struct {
	// Execution symbol.
	Symbol string `json:"symbol"`

	// Internal symbol identifier
	SymbolID int `json:"symbolId"`

	// Execution quantity.
	Quantity int `json:"quantity"`

	// Client side of the order to which execution belongs.
	Side string `json:"side"`

	// Execution price.
	Price float64 `json:"price"`

	// Internal identifier of the execution.
	ID int `json:"id"`

	// Internal identifier of the order to which the execution belongs.
	OrderID int `json:"orderId"`

	// Internal identifier of the order chain to which the execution belongs.
	OrderChainID int `json:"orderChainId"`

	// Identifier of the execution at the market where it originated.
	ExchangeExecID string `json:"exchangeExecId"`

	// Execution timestamp.
	Timestamp time.Time `json:"timestamp"`

	// Manual notes that may have been entered by Trade Desk staff
	Notes string `json:"notes"`

	// Trading venue where execution originated.
	Venue string `json:"venue"`

	// Execution cost (price x quantity).
	TotalCost float64 `json:"totalCost"`

	// Questrade commission for orders placed with Trade Desk.
	OrderPlacementCommission float64 `json:"orderPlacementCommission"`

	// Questrade commission.
	Commission float64 `json:"commission"`

	// Liquidity fee charged by execution venue.
	ExecutionFee float64 `json:"executionFee"`

	// SEC fee charged on all sales of US securities.
	SecFee float64 `json:"secFee"`

	// Additional execution fee charged by TSX (if applicable).
	CanadianExecutionFee int `json:"canadianExecutionFee"`

	// Internal identifierof the parent order.
	ParentID int `json:"parentId"`
}
