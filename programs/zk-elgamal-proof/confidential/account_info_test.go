package confidential

import (
	"errors"
	"math"
	"testing"

	token2022 "github.com/gagliardetto/solana-go/programs/token-2022"
	"github.com/gagliardetto/solana-go/programs/token-2022/zkencryption"
	zk "github.com/gagliardetto/solana-go/programs/zk-elgamal-proof"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/encryption"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/internal/zktest"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/proofdata"
)

func TestApplyPendingBalanceAccountInfo(t *testing.T) {
	kp, aesKey := generateSourceAccount(t)
	const (
		pendingLo     = uint64(65535)
		pendingHi     = uint64(3)
		available     = uint64(1000)
		creditCounter = uint64(2)
	)
	pendingCombined := pendingHi<<AmountLoBitLength + pendingLo

	applyPendingBalanceAccountInfo := NewApplyPendingBalanceAccountInfo(makeAccountState(t, kp, aesKey,
		pendingLo, pendingHi, available, creditCounter))

	if storedCreditCounter := applyPendingBalanceAccountInfo.PendingBalanceCreditCounter(); storedCreditCounter != creditCounter {
		t.Fatalf("PendingBalanceCreditCounter() = %d, want %d", storedCreditCounter, creditCounter)
	}
	if !applyPendingBalanceAccountInfo.HasPendingBalance() {
		t.Fatal("HasPendingBalance() = false with pending credits")
	}
	if storedPendingBalance, err := applyPendingBalanceAccountInfo.PendingBalance(kp); err != nil || storedPendingBalance != pendingCombined {
		t.Fatalf("PendingBalance() = (%d, %v), want (%d, nil)", storedPendingBalance, err, pendingCombined)
	}
	if storedAvailableBalance, err := applyPendingBalanceAccountInfo.AvailableBalance(aesKey); err != nil || storedAvailableBalance != available {
		t.Fatalf("AvailableBalance() = (%d, %v), want (%d, nil)", storedAvailableBalance, err, available)
	}
	if storedTotalBalance, err := applyPendingBalanceAccountInfo.TotalBalance(kp, aesKey); err != nil || storedTotalBalance != available+pendingCombined {
		t.Fatalf("TotalBalance() = (%d, %v), want (%d, nil)", storedTotalBalance, err, available+pendingCombined)
	}

	storedNewDecryptableBalance, err := applyPendingBalanceAccountInfo.NewDecryptableAvailableBalance(kp, aesKey)
	if err != nil {
		t.Fatal(err)
	}
	if storedNewDecryptedBalance, err := encryption.AeDecrypt(aesKey, storedNewDecryptableBalance); err != nil || storedNewDecryptedBalance != available+pendingCombined {
		t.Fatalf("new decryptable balance = (%d, %v), want (%d, nil)", storedNewDecryptedBalance, err, available+pendingCombined)
	}

	// Assert that account with zero pending balance reports HasPendingBalance() as false and reports TotalBalance equal to available balance
	zeroPendingBalanceAccountInfo := NewApplyPendingBalanceAccountInfo(makeAccountState(t, kp, aesKey, 0, 0, available, 0))
	if zeroPendingBalanceAccountInfo.HasPendingBalance() {
		t.Fatal("HasPendingBalance() = true with no pending credits")
	}
	if storedTotalBalance, err := zeroPendingBalanceAccountInfo.TotalBalance(kp, aesKey); err != nil || storedTotalBalance != available {
		t.Fatalf("TotalBalance() = (%d, %v), want (%d, nil)", storedTotalBalance, err, available)
	}

	// Assert that full token account cannot perform operations involving pending balance
	// i.e cannot report total balance or a new encryption of what the balance wil be after application of pending balance.
	applyPendingBalanceFullAccountInfo := NewApplyPendingBalanceAccountInfo(makeAccountState(t, kp, aesKey, 1, 0, math.MaxUint64, 1))
	if _, err := applyPendingBalanceFullAccountInfo.TotalBalance(kp, aesKey); !errors.Is(err, ErrBalanceOverflow) {
		t.Fatalf("TotalBalance() on overflowing account: got %v, want ErrBalanceOverflow", err)
	}
	if _, err := applyPendingBalanceFullAccountInfo.NewDecryptableAvailableBalance(kp, aesKey); !errors.Is(err, ErrBalanceOverflow) {
		t.Fatalf("NewDecryptableAvailableBalance() on overflowing account: got %v, want ErrBalanceOverflow", err)
	}

	// Assert wrong AE key cannot decrypt the available balance.
	wrongKey, err := encryption.NewAeKey()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := applyPendingBalanceAccountInfo.AvailableBalance(wrongKey); !errors.Is(err, zk.ErrDecryption) {
		t.Fatalf("AvailableBalance() with wrong AE key: got %v, want zk.ErrDecryption", err)
	}
}

func TestWithdrawAccountInfo(t *testing.T) {
	kp, aesKey := generateSourceAccount(t)
	const (
		available = uint64(1000)
		amount    = uint64(400)
	)
	withdrawAccountInfo := NewWithdrawAccountInfo(makeAccountState(t, kp, aesKey, 0, 0, available, 0))

	withdrawalProofData, err := withdrawAccountInfo.GenerateProofData(amount, kp, aesKey)
	if err != nil {
		t.Fatal(err)
	}
	verifyAll(t, map[string]proofdata.ProofData{
		"equality": withdrawalProofData.EqualityProofData,
		"range":    withdrawalProofData.RangeProofData,
	})

	postWithdrawalDecryptableBalance, err := withdrawAccountInfo.NewDecryptableAvailableBalance(amount, aesKey)
	if err != nil {
		t.Fatal(err)
	}
	if decryptedPostWithdrawalBalance, err := encryption.AeDecrypt(aesKey, postWithdrawalDecryptableBalance); err != nil || decryptedPostWithdrawalBalance != available-amount {
		t.Fatalf("new decryptable balance = (%d, %v), want (%d, nil)", decryptedPostWithdrawalBalance, err, available-amount)
	}

	// A withdrawal exceeding the balance is rejected.
	if _, err := withdrawAccountInfo.GenerateProofData(available+1, kp, aesKey); !errors.Is(err, ErrNotEnoughFunds) {
		t.Fatalf("GenerateProofData() exceeding balance: got %v, want ErrNotEnoughFunds", err)
	}
	if _, err := withdrawAccountInfo.NewDecryptableAvailableBalance(available+1, aesKey); !errors.Is(err, ErrNotEnoughFunds) {
		t.Fatalf("NewDecryptableAvailableBalance() exceeding balance: got %v, want ErrNotEnoughFunds", err)
	}
}

func TestTransferAccountInfo(t *testing.T) {
	sender, aesKey := generateSourceAccount(t)
	recipient := zktest.GenKeyPair(t)
	auditor := zktest.GenKeyPair(t)
	feeCollector := zktest.GenKeyPair(t)
	const (
		available = uint64(1000)
		amount    = uint64(500)
	)
	transferAccountInfo := NewTransferAccountInfo(makeAccountState(t, sender, aesKey, 0, 0, available, 0))

	transferproofData, err := transferAccountInfo.GenerateSplitTransferProofData(amount, sender, aesKey,
		recipient.Pubkey, &auditor.Pubkey)
	if err != nil {
		t.Fatal(err)
	}
	verifyAll(t, map[string]proofdata.ProofData{
		"equality": transferproofData.EqualityProofData,
		"validity": transferproofData.CiphertextValidityProofDataWithCiphertext.ProofData,
		"range":    transferproofData.RangeProofData,
	})

	feeProofData, err := transferAccountInfo.GenerateSplitTransferWithFeeProofData(amount, sender, aesKey,
		recipient.Pubkey, &auditor.Pubkey, feeCollector.Pubkey, 250, 10_000)
	if err != nil {
		t.Fatal(err)
	}
	verifyAll(t, map[string]proofdata.ProofData{
		"equality":     feeProofData.EqualityProofData,
		"validity":     feeProofData.TransferAmountCiphertextValidityProofDataWithCiphertext.ProofData,
		"percentage":   feeProofData.PercentageWithCapProofData,
		"fee validity": feeProofData.FeeCiphertextValidityProofData,
		"range":        feeProofData.RangeProofData,
	})

	postTransferDecryptableBalance, err := transferAccountInfo.NewDecryptableAvailableBalance(amount, aesKey)
	if err != nil {
		t.Fatal(err)
	}
	if decryptedPostTransferDecryptableBalance, err := encryption.AeDecrypt(aesKey, postTransferDecryptableBalance); err != nil || decryptedPostTransferDecryptableBalance != available-amount {
		t.Fatalf("new decryptable balance = (%d, %v), want (%d, nil)", decryptedPostTransferDecryptableBalance, err, available-amount)
	}

	// A transfer exceeding the balance is rejected.
	if _, err := transferAccountInfo.GenerateSplitTransferProofData(available+1, sender, aesKey,
		recipient.Pubkey, &auditor.Pubkey); !errors.Is(err, ErrNotEnoughFunds) {
		t.Fatalf("GenerateSplitTransferProofData() exceeding balance: got %v, want ErrNotEnoughFunds", err)
	}
	if _, err := transferAccountInfo.NewDecryptableAvailableBalance(available+1, aesKey); !errors.Is(err, ErrNotEnoughFunds) {
		t.Fatalf("NewDecryptableAvailableBalance() exceeding balance: got %v, want ErrNotEnoughFunds", err)
	}
}

func TestEmptyAccountInfo(t *testing.T) {
	kp, aesKey := generateSourceAccount(t)

	emptyAccountInfo := NewEmptyAccountInfo(makeAccountState(t, kp, aesKey, 0, 0, 0, 0))
	emptyAccountProofData, err := emptyAccountInfo.GenerateProofData(kp)
	if err != nil {
		t.Fatal(err)
	}
	if err := emptyAccountProofData.Verify(); err != nil {
		t.Fatalf("zero-ciphertext proof rejected: %v", err)
	}

	// Assert that account with pending balance is considered empty.
	pendingBalanceEmptyAccountInfo := NewEmptyAccountInfo(makeAccountState(t, kp, aesKey, 1, 0, 0, 0))
	pendingBalanceEmptyAccountProofData, err := pendingBalanceEmptyAccountInfo.GenerateProofData(kp)
	if err != nil {
		t.Fatal(err)
	}
	if err := pendingBalanceEmptyAccountProofData.Verify(); err != nil {
		t.Fatalf("zero-ciphertext proof rejected: %v", err)
	}

	// Assert account with funds cannot prove a zero balance.
	nonEmptyAccountInfo := NewEmptyAccountInfo(makeAccountState(t, kp, aesKey, 0, 0, 1, 0))
	if _, err := nonEmptyAccountInfo.GenerateProofData(kp); !errors.Is(err, zk.ErrProofGeneration) {
		t.Fatalf("GenerateProofData() on non-empty account: got %v, want zk.ErrProofGeneration", err)
	}

}

// makeAccountState builds a confidential transfer account state  for testing
func makeAccountState(
	t *testing.T,
	kp *encryption.ElGamalKeypair,
	aesKey zkencryption.AeKey,
	pendingLo, pendingHi, available, creditCounter uint64,
) *token2022.ConfidentialTransferAccountState {
	t.Helper()
	encrypt := func(amount uint64) [64]byte {
		ct, err := kp.Pubkey.Encrypt(amount)
		if err != nil {
			t.Fatal(err)
		}
		return ct
	}
	decryptable, err := encryption.AeEncrypt(aesKey, available)
	if err != nil {
		t.Fatal(err)
	}
	return &token2022.ConfidentialTransferAccountState{
		ElGamalPubkey:               kp.Pubkey,
		PendingBalanceLo:            encrypt(pendingLo),
		PendingBalanceHi:            encrypt(pendingHi),
		AvailableBalance:            encrypt(available),
		DecryptableAvailableBalance: decryptable,
		PendingBalanceCreditCounter: creditCounter,
	}
}
