package transaction

import (
	"errors"
	"math/big"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/data/receipt"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/transaction"
	"github.com/ElrondNetwork/elrond-go/hashing"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

// txProcessor implements TransactionProcessor interface and can modify account states according to a transaction
type txProcessor struct {
	*baseTxProcessor
	hasher           hashing.Hasher
	scProcessor      process.SmartContractProcessor
	marshalizer      marshal.Marshalizer
	txFeeHandler     process.TransactionFeeHandler
	txTypeHandler    process.TxTypeHandler
	receiptForwarder process.IntermediateTransactionHandler
	badTxForwarder   process.IntermediateTransactionHandler
}

// NewTxProcessor creates a new txProcessor engine
func NewTxProcessor(
	accounts state.AccountsAdapter,
	hasher hashing.Hasher,
	addressConv state.AddressConverter,
	marshalizer marshal.Marshalizer,
	shardCoordinator sharding.Coordinator,
	scProcessor process.SmartContractProcessor,
	txFeeHandler process.TransactionFeeHandler,
	txTypeHandler process.TxTypeHandler,
	economicsFee process.FeeHandler,
	receiptForwarder process.IntermediateTransactionHandler,
	badTxForwarder process.IntermediateTransactionHandler,
) (*txProcessor, error) {

	if check.IfNil(accounts) {
		return nil, process.ErrNilAccountsAdapter
	}
	if check.IfNil(hasher) {
		return nil, process.ErrNilHasher
	}
	if check.IfNil(addressConv) {
		return nil, process.ErrNilAddressConverter
	}
	if check.IfNil(marshalizer) {
		return nil, process.ErrNilMarshalizer
	}
	if check.IfNil(shardCoordinator) {
		return nil, process.ErrNilShardCoordinator
	}
	if check.IfNil(scProcessor) {
		return nil, process.ErrNilSmartContractProcessor
	}
	if check.IfNil(txFeeHandler) {
		return nil, process.ErrNilUnsignedTxHandler
	}
	if check.IfNil(txTypeHandler) {
		return nil, process.ErrNilTxTypeHandler
	}
	if check.IfNil(economicsFee) {
		return nil, process.ErrNilEconomicsFeeHandler
	}
	if check.IfNil(receiptForwarder) {
		return nil, process.ErrNilReceiptHandler
	}
	if check.IfNil(badTxForwarder) {
		return nil, process.ErrNilBadTxHandler
	}

	baseTxProcess := &baseTxProcessor{
		accounts:         accounts,
		shardCoordinator: shardCoordinator,
		adrConv:          addressConv,
		economicsFee:     economicsFee,
	}

	return &txProcessor{
		baseTxProcessor:  baseTxProcess,
		hasher:           hasher,
		marshalizer:      marshalizer,
		scProcessor:      scProcessor,
		txFeeHandler:     txFeeHandler,
		txTypeHandler:    txTypeHandler,
		receiptForwarder: receiptForwarder,
		badTxForwarder:   badTxForwarder,
	}, nil
}

// ProcessTransaction modifies the account states in respect with the transaction data
func (txProc *txProcessor) ProcessTransaction(tx *transaction.Transaction) error {
	if check.IfNil(tx) {
		return process.ErrNilTransaction
	}

	adrSrc, adrDst, err := txProc.getAddresses(tx)
	if err != nil {
		return err
	}

	acntSnd, err := txProc.getAccountFromAddress(adrSrc)
	if err != nil {
		return err
	}

	process.DisplayProcessTxDetails("ProcessTransaction: sender account details", acntSnd, tx)

	err = txProc.checkTxValues(tx, acntSnd)
	if err != nil {
		if errors.Is(err, process.ErrInsufficientFunds) {
			receiptErr := txProc.executingFailedTransaction(tx, acntSnd, err)
			if receiptErr != nil {
				return receiptErr
			}
		}
		return err
	}

	txType, err := txProc.txTypeHandler.ComputeTransactionType(tx)
	if err != nil {
		return err
	}

	switch txType {
	case process.MoveBalance:
		return txProc.processMoveBalance(tx, adrSrc, adrDst)
	case process.SCDeployment:
		return txProc.processSCDeployment(tx, adrSrc)
	case process.SCInvoking:
		return txProc.processSCInvoking(tx, adrSrc, adrDst)
	}

	return process.ErrWrongTransaction
}

func (txProc *txProcessor) executingFailedTransaction(
	tx *transaction.Transaction,
	acntSnd state.AccountHandler,
	txError error,
) error {
	if check.IfNil(acntSnd) {
		return nil
	}

	account, ok := acntSnd.(state.UserAccountHandler)
	if !ok {
		return process.ErrWrongTypeAssertion
	}

	txFee := txProc.economicsFee.ComputeFee(tx)

	operation := big.NewInt(0)
	account.SetBalance(operation.Sub(account.GetBalance(), txFee))
	account.SetNonce(account.GetNonce() + 1)

	err := txProc.badTxForwarder.AddIntermediateTransactions([]data.TransactionHandler{tx})
	if err != nil {
		return err
	}

	txHash, err := core.CalculateHash(txProc.marshalizer, txProc.hasher, tx)
	if err != nil {
		return err
	}

	rpt := &receipt.Receipt{
		Value:   big.NewInt(0).Set(txFee),
		SndAddr: tx.SndAddr,
		Data:    []byte(txError.Error()),
		TxHash:  txHash,
	}

	err = txProc.receiptForwarder.AddIntermediateTransactions([]data.TransactionHandler{rpt})
	if err != nil {
		return err
	}

	txProc.txFeeHandler.ProcessTransactionFee(txFee)

	err = txProc.accounts.SaveAccount(account)
	if err != nil {
		return err
	}

	return process.ErrFailedTransaction
}

func (txProc *txProcessor) createReceiptWithReturnedGas(tx *transaction.Transaction, acntSnd state.UserAccountHandler) error {
	if check.IfNil(acntSnd) {
		return nil
	}
	if core.IsSmartContractAddress(tx.RcvAddr) {
		return nil
	}

	totalProvided := big.NewInt(0)
	totalProvided.Mul(big.NewInt(0).SetUint64(tx.GasPrice), big.NewInt(0).SetUint64(tx.GasLimit))

	actualCost := txProc.economicsFee.ComputeFee(tx)
	refundValue := big.NewInt(0).Sub(totalProvided, actualCost)

	zero := big.NewInt(0)
	if refundValue.Cmp(zero) == 0 {
		return nil
	}

	txHash, err := core.CalculateHash(txProc.marshalizer, txProc.hasher, tx)
	if err != nil {
		return err
	}

	rpt := &receipt.Receipt{
		Value:   big.NewInt(0).Set(refundValue),
		SndAddr: tx.SndAddr,
		Data:    []byte("refundedGas"),
		TxHash:  txHash,
	}

	err = txProc.receiptForwarder.AddIntermediateTransactions([]data.TransactionHandler{rpt})
	if err != nil {
		return err
	}

	return nil
}

func (txProc *txProcessor) processTxFee(tx *transaction.Transaction, acntSnd state.UserAccountHandler) (*big.Int, error) {
	if acntSnd == nil {
		return big.NewInt(0), nil
	}

	cost := txProc.economicsFee.ComputeFee(tx)
	newBalance := big.NewInt(0).Add(acntSnd.GetBalance(), big.NewInt(0).Neg(cost))
	if newBalance.Cmp(big.NewInt(0)) < 0 {
		return nil, process.ErrInsufficientFunds
	}

	acntSnd.SetBalance(newBalance)

	return cost, nil
}

func (txProc *txProcessor) processMoveBalance(
	tx *transaction.Transaction,
	adrSrc, adrDst state.AddressContainer,
) error {

	// getAccounts returns acntSrc not nil if the adrSrc is in the node shard, the same, acntDst will be not nil
	// if adrDst is in the node shard. If an error occurs it will be signaled in err variable.
	acntSrc, acntDst, err := txProc.getAccounts(adrSrc, adrDst)
	if err != nil {
		return err
	}

	txFee, err := txProc.processTxFee(tx, acntSrc)
	if err != nil {
		return err
	}

	value := tx.Value

	err = txProc.moveBalances(acntSrc, acntDst, value)
	if err != nil {
		return err
	}

	// is sender address in node shard
	if acntSrc != nil {
		acntSrc.SetNonce(acntSrc.GetNonce() + 1)
	}

	err = txProc.createReceiptWithReturnedGas(tx, acntSrc)
	if err != nil {
		return err
	}

	txProc.txFeeHandler.ProcessTransactionFee(txFee)
	err = txProc.saveAccounts(acntSrc, acntDst)
	if err != nil {
		return err
	}

	return nil
}

func (txProc *txProcessor) saveAccounts(acntSnd, acntDst state.AccountHandler) error {
	if !check.IfNil(acntSnd) {
		err := txProc.accounts.SaveAccount(acntSnd)
		if err != nil {
			return err
		}
	}

	if !check.IfNil(acntDst) {
		err := txProc.accounts.SaveAccount(acntDst)
		if err != nil {
			return err
		}
	}

	return nil
}

func (txProc *txProcessor) processSCDeployment(
	tx *transaction.Transaction,
	adrSrc state.AddressContainer,
) error {
	// getAccounts returns acntSrc not nil if the adrSrc is in the node shard, the same, acntDst will be not nil
	// if adrDst is in the node shard. If an error occurs it will be signaled in err variable.
	acntSrc, err := txProc.getAccountFromAddress(adrSrc)
	if err != nil {
		return err
	}

	err = txProc.scProcessor.DeploySmartContract(tx, acntSrc)
	return err
}

func (txProc *txProcessor) processSCInvoking(
	tx *transaction.Transaction,
	adrSrc, adrDst state.AddressContainer,
) error {
	// getAccounts returns acntSrc not nil if the adrSrc is in the node shard, the same, acntDst will be not nil
	// if adrDst is in the node shard. If an error occurs it will be signaled in err variable.
	acntSrc, acntDst, err := txProc.getAccounts(adrSrc, adrDst)
	if err != nil {
		return err
	}

	err = txProc.scProcessor.ExecuteSmartContractTransaction(tx, acntSrc, acntDst)
	return err
}

func (txProc *txProcessor) moveBalances(
	acntSrc, acntDst state.UserAccountHandler,
	value *big.Int,
) error {
	// is sender address in node shard
	if !check.IfNil(acntSrc) {
		newBalance := big.NewInt(0).Add(acntSrc.GetBalance(), big.NewInt(0).Neg(value))
		if newBalance.Cmp(big.NewInt(0)) < 0 {
			return process.ErrInsufficientFunds
		}

		acntSrc.SetBalance(newBalance)
	}

	// is receiver address in node shard
	if !check.IfNil(acntDst) {
		acntDst.SetBalance(big.NewInt(0).Add(acntDst.GetBalance(), value))
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (txProc *txProcessor) IsInterfaceNil() bool {
	return txProc == nil
}
