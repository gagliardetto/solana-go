package confidential

import (
	"math"

	token2022 "github.com/gagliardetto/solana-go/programs/token-2022"
	"github.com/gagliardetto/solana-go/programs/token-2022/zkencryption"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/encryption"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/proofdata"
)

// --- ApplyPendingBalance ---

// ApplyPendingBalanceAccountInfo is the account state an ApplyPendingBalance instruction is built from.
type ApplyPendingBalanceAccountInfo struct {
	pendingBalanceCreditCounter uint64
	pendingBalanceLo            encryption.ElGamalCiphertext
	pendingBalanceHi            encryption.ElGamalCiphertext
	decryptableAvailableBalance encryption.AeCiphertext
}

// NewApplyPendingBalanceAccountInfo extracts the state an ApplyPendingBalance
// instruction needs from a confidential transfer account state.
func NewApplyPendingBalanceAccountInfo(s *token2022.ConfidentialTransferAccountState) ApplyPendingBalanceAccountInfo {
	return ApplyPendingBalanceAccountInfo{
		pendingBalanceCreditCounter: s.PendingBalanceCreditCounter,
		pendingBalanceLo:            encryption.ElGamalCiphertext(s.PendingBalanceLo),
		pendingBalanceHi:            encryption.ElGamalCiphertext(s.PendingBalanceHi),
		decryptableAvailableBalance: encryption.AeCiphertext(s.DecryptableAvailableBalance),
	}
}

// PendingBalanceCreditCounter is the number of credits received since last pending credit application.
func (i ApplyPendingBalanceAccountInfo) PendingBalanceCreditCounter() uint64 {
	return i.pendingBalanceCreditCounter
}

// HasPendingBalance reports whether any deposit or transfer has been credited
// to the pending balance since the last ApplyPendingBalance.
func (i ApplyPendingBalanceAccountInfo) HasPendingBalance() bool {
	return i.pendingBalanceCreditCounter > 0
}

// PendingBalance decrypts both halves of the pending balance and recombines them.
func (i ApplyPendingBalanceAccountInfo) PendingBalance(kp *encryption.ElGamalKeypair) (uint64, error) {
	lo, err := kp.DecryptU32(i.pendingBalanceLo)
	if err != nil {
		return 0, err
	}
	hi, err := kp.DecryptU32(i.pendingBalanceHi)
	if err != nil {
		return 0, err
	}
	// hi stores the summed high part of account credits,
	// which must be recovered by left-shifting by `AmountLoBitLength`.
	return hi<<AmountLoBitLength + lo, nil
}

// AvailableBalance decrypts the account's decryptable available balance.
func (i ApplyPendingBalanceAccountInfo) AvailableBalance(aesKey zkencryption.AeKey) (uint64, error) {
	return encryption.AeDecrypt(aesKey, i.decryptableAvailableBalance)
}

// TotalBalance is the available balance plus the pending balance: what the
// account will hold once ApplyPendingBalance is applied.
func (i ApplyPendingBalanceAccountInfo) TotalBalance(kp *encryption.ElGamalKeypair, aesKey zkencryption.AeKey) (uint64, error) {
	available, err := i.AvailableBalance(aesKey)
	if err != nil {
		return 0, err
	}
	pending, err := i.PendingBalance(kp)
	if err != nil {
		return 0, err
	}
	if pending > math.MaxUint64-available {
		return 0, ErrBalanceOverflow
	}
	return available + pending, nil
}

// NewDecryptableAvailableBalance is the AE ciphertext the ApplyPendingBalance
// instruction carries: the account's total balance encrypted with aesKey.
func (i ApplyPendingBalanceAccountInfo) NewDecryptableAvailableBalance(
	kp *encryption.ElGamalKeypair, aesKey zkencryption.AeKey,
) (encryption.AeCiphertext, error) {
	total, err := i.TotalBalance(kp, aesKey)
	if err != nil {
		return encryption.AeCiphertext{}, err
	}
	return encryption.AeEncrypt(aesKey, total)
}

// --- Withdraw ---

// WithdrawInfo is the account state a confidential Withdraw instruction is built from.
type WithdrawInfo struct {
	availableBalance            encryption.ElGamalCiphertext
	decryptableAvailableBalance encryption.AeCiphertext
}

// NewWithdrawInfo extracts the state a Withdraw instruction needs from a confidential transfer account.
func NewWithdrawInfo(s *token2022.ConfidentialTransferAccountState) WithdrawInfo {
	return WithdrawInfo{
		availableBalance:            encryption.ElGamalCiphertext(s.AvailableBalance),
		decryptableAvailableBalance: encryption.AeCiphertext(s.DecryptableAvailableBalance),
	}
}

// GenerateProofData builds the proofs the Withdraw instruction carries.
func (i WithdrawInfo) GenerateProofData(
	amount uint64, kp *encryption.ElGamalKeypair, aesKey zkencryption.AeKey,
) (*WithdrawProofData, error) {
	balance, err := encryption.AeDecrypt(aesKey, i.decryptableAvailableBalance)
	if err != nil {
		return nil, err
	}
	return NewWithdrawProofData(i.availableBalance, balance, amount, kp)
}

// NewDecryptableAvailableBalance is the AE ciphertext the Withdraw instruction
// carries: the available balance less the withdrawn amount, encrypted with aesKey.
func (i WithdrawInfo) NewDecryptableAvailableBalance(
	amount uint64, aesKey zkencryption.AeKey,
) (encryption.AeCiphertext, error) {
	return decryptableBalanceAfterSpend(aesKey, i.decryptableAvailableBalance, amount)
}

// --- Transfer ---

// TransferInfo is the account state a confidential Transfer or TransferWithFee instruction is built from.
type TransferInfo struct {
	availableBalance            encryption.ElGamalCiphertext
	decryptableAvailableBalance encryption.AeCiphertext
}

// NewTransferInfo extracts the TransferInfo state from a confidential transfer account.
func NewTransferInfo(s *token2022.ConfidentialTransferAccountState) TransferInfo {
	return TransferInfo{
		availableBalance:            encryption.ElGamalCiphertext(s.AvailableBalance),
		decryptableAvailableBalance: encryption.AeCiphertext(s.DecryptableAvailableBalance),
	}
}

// GenerateSplitTransferProofData builds the three proofs a confidential transfer requires.
func (i TransferInfo) GenerateSplitTransferProofData(
	amount uint64,
	kp *encryption.ElGamalKeypair,
	aesKey zkencryption.AeKey,
	destinationPubkey encryption.ElGamalPubkey,
	auditorPubkey *encryption.ElGamalPubkey,
) (*TransferProofData, error) {
	return TransferSplitProofData(i.availableBalance, i.decryptableAvailableBalance,
		amount, kp, aesKey, destinationPubkey, auditorPubkey)
}

// GenerateSplitTransferWithFeeProofData builds the five proofs a confidential transfer on a fee-extended mint requires.
func (i TransferInfo) GenerateSplitTransferWithFeeProofData(
	amount uint64,
	kp *encryption.ElGamalKeypair,
	aesKey zkencryption.AeKey,
	destinationPubkey encryption.ElGamalPubkey,
	auditorPubkey *encryption.ElGamalPubkey,
	withdrawWithheldAuthorityPubkey encryption.ElGamalPubkey,
	feeRateBasisPoints uint16,
	maximumFee uint64,
) (*TransferWithFeeProofData, error) {
	return TransferWithFeeSplitProofData(i.availableBalance, i.decryptableAvailableBalance,
		amount, kp, aesKey, destinationPubkey, auditorPubkey,
		withdrawWithheldAuthorityPubkey, feeRateBasisPoints, maximumFee)
}

// NewDecryptableAvailableBalance is the AE ciphertext a Transfer instruction
// carries: the available balance less the transferred amount, encrypted with aesKey.
func (i TransferInfo) NewDecryptableAvailableBalance(
	amount uint64, aesKey zkencryption.AeKey,
) (encryption.AeCiphertext, error) {
	return decryptableBalanceAfterSpend(aesKey, i.decryptableAvailableBalance, amount)
}

// decryptableBalanceAfterSpend decrypts an available balance, subtracts the
// amount spent, and encrypts the remainder afresh under the same AE key.
func decryptableBalanceAfterSpend(
	aesKey zkencryption.AeKey, ct encryption.AeCiphertext, amount uint64,
) (encryption.AeCiphertext, error) {
	balance, err := encryption.AeDecrypt(aesKey, ct)
	if err != nil {
		return encryption.AeCiphertext{}, err
	}
	if amount > balance {
		return encryption.AeCiphertext{}, ErrNotEnoughFunds
	}
	return encryption.AeEncrypt(aesKey, balance-amount)
}

// --- EmptyAccount ---

// EmptyAccountInfo is the account state an EmptyAccount instruction is built from.
type EmptyAccountInfo struct {
	availableBalance encryption.ElGamalCiphertext
}

// NewEmptyAccountInfo extracts the state an EmptyAccount instruction needs from a confidential transfer account.
func NewEmptyAccountInfo(s *token2022.ConfidentialTransferAccountState) EmptyAccountInfo {
	return EmptyAccountInfo{availableBalance: encryption.ElGamalCiphertext(s.AvailableBalance)}
}

// GenerateProofData builds the  proof the EmptyAccount instruction carries.
func (i EmptyAccountInfo) GenerateProofData(kp *encryption.ElGamalKeypair) (*proofdata.ZeroCiphertextProofData, error) {
	return proofdata.NewZeroCiphertextProofData(kp, i.availableBalance)
}
