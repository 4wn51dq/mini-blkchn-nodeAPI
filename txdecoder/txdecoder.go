package txdecoder

// input is a hexcode of transaction
// outputs are the datafields of a tx in ethereum:
/*
 * nonce: index (gets increment)
 * recipient: receiving address (160bit address or nil in case of contact creation)
 * value: amount of ETH involved in the transaction
 * yParity, r, s: digital sign components
 * init/data: calldata (contract creation or message call in smart contract)
 * gasLimit: max consumable gas units
 * type: 0 or 2
 * maxPriorityFeePerGas: validator tip
 * max fee per gas: max gas the transaction is willing to give out
 * chainId
 */

type Tx struct {
	nonce                    uint
	chaidId                  uint
	recipient                string
	value                    uint
	yParity, r, s            string
	data                     string
	tx_type                  uint8
	gas_limit                uint
	max_priority_fee_per_gas uint
	max_fee_per_gas          uint
}

/**
* raw bytes
* rlp decode
*
 */

func RlpDecode(raw []byte) {

}
