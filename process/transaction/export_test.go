package transaction

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/transaction"
)

type TxProcessor *txProcessor

func (txProc *txProcessor) GetAddresses(tx *transaction.Transaction) (adrSrc, adrDst state.AddressContainer, err error) {
	return txProc.getAddresses(tx)
}

func (txProc *txProcessor) GetAccounts(adrSrc, adrDst state.AddressContainer,
) (acntSrc, acntDst state.UserAccountHandler, err error) {
	return txProc.getAccounts(adrSrc, adrDst)
}

func (txProc *txProcessor) CheckTxValues(tx *transaction.Transaction, acntSnd state.AccountHandler) error {
	return txProc.checkTxValues(tx, acntSnd)
}

func (txProc *txProcessor) MoveBalances(acntSrc, acntDst state.UserAccountHandler, value *big.Int) error {
	return txProc.moveBalances(acntSrc, acntDst, value)
}

func (txProc *txProcessor) IncreaseNonce(acntSrc state.UserAccountHandler) error {
	acntSrc.SetNonce(acntSrc.GetNonce() + 1)
	return nil
}
