package api

type Account struct {
	Address string
	Balance uint
}

type TX struct {
	ID     string
	from   string
	to     string
	amount uint
}

/**
 * @notice these below are the parameters the API will take
 */

type HealthResponse struct {
	Status string
}

type SubmitTxRequest struct {
	Tx TX
}

type SubmitTxResponse struct {
	Status string
}
