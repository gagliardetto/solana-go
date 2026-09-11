package token2022

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token-2022/zkencryption"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/confidential"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/encryption"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/proofdata"
	"github.com/gagliardetto/solana-go/programs/zk-elgamal-proof/zkprogram"
)

func TestConfidentialTransferClientInlineProofs(t *testing.T) {
	t.Parallel()
	for _, op := range ctcOperations(t) {
		t.Run(op.name, func(t *testing.T) {
			t.Parallel()
			instructions, err := op.build(true)
			assertNoErr(t, err)
			if got, want := len(instructions), 1+len(op.verifiers); got != want {
				t.Fatalf("got %d instructions, want %d", got, want)
			}
			subInstruction, data := ctcExtensionInstruction(t, instructions[0])
			if subInstruction != op.subInstruction {
				t.Errorf("sub-instruction = %d, want %d", subInstruction, op.subInstruction)
			}
			for i, offset := range op.offsets(data) {
				if offset != int8(i+1) {
					t.Errorf("proof offset %d = %d, want %d for a proof inlined in order", i, offset, i+1)
				}
			}
			for i, verifier := range op.verifiers {
				verifyInstruction := instructions[1+i]
				if got := verifyInstruction.ProgramID(); !got.Equals(zkprogram.ProgramID) {
					t.Fatalf("program ID = %s, want %s", got, zkprogram.ProgramID)
				}
				if got := len(verifyInstruction.Accounts()); got != 0 {
					t.Errorf("%s takes %d accounts, want none for an inlined proof", verifier, got)
				}
				verifyData, err := verifyInstruction.Data()
				assertNoErr(t, err)
				if got := zkprogram.ProofInstruction(verifyData[0]); got != verifier {
					t.Fatalf("proof instruction = %s, want %s", got, verifier)
				}
				proof := proofdata.NewProofData(proofdata.ProofType(verifier))
				if err := proof.UnmarshalBinary(verifyData[1:]); err != nil {
					t.Fatalf("%s proof data: %v", verifier, err)
				}
				if err := proof.Verify(); err != nil {
					t.Fatalf("%s proof rejected: %v", verifier, err)
				}
			}
			if op.checkInline != nil {
				op.checkInline(t, data)
			}
		})
	}
}

func TestConfidentialTransferClientContextStateAccounts(t *testing.T) {
	t.Parallel()
	for _, op := range ctcOperations(t) {
		t.Run(op.name, func(t *testing.T) {
			t.Parallel()
			instructions, err := op.build(false)
			assertNoErr(t, err)
			if len(instructions) != 1 {
				t.Fatalf("got %d instructions, want 1: no proof is generated when every proof account is set", len(instructions))
			}
			_, data := ctcExtensionInstruction(t, instructions[0])
			for i, offset := range op.offsets(data) {
				if offset != 0 {
					t.Errorf("proof offset %d = %d, want 0 for a context state account", i, offset)
				}
			}
			for _, want := range op.accounts {
				if !ctcHasAccount(instructions[0], want) {
					t.Errorf("context state account %s missing from the instruction accounts", want)
				}
			}
			if ctcHasAccount(instructions[0], solana.SysVarInstructionsPubkey) {
				t.Error("instructions sysvar present, want it only when a proof is inlined")
			}
			if op.checkAccounts != nil {
				op.checkAccounts(t, data)
			}
		})
	}
}

// TestConfidentialTransferClientMixedProofLocations checks the limitation that an operation cannot
// take one proof from a context state account and inline another.
func TestConfidentialTransferClientMixedProofLocations(t *testing.T) {
	t.Parallel()
	kp, aesKey := ctcKeys(t)
	state := ctcState(t, kp, aesKey, ctcAvailableBalance, 0)

	_, err := ConfidentialTransferWithdraw(
		ctTokenAccount, ctMint, ctAuthority, nil, &ctContextEquality, nil,
		ctcTransferAmount, ctDecimals, NewWithdrawAccountInfo(state), kp, aesKey)
	if err == nil || !strings.Contains(err.Error(), "proof instruction offset") {
		t.Fatalf("got %v, want a proof instruction offset error", err)
	}
}

// TestConfidentialTransferCreateContextStateAccountRejectsMissingProof checks
// a proof is required, rather than panicking on the nil.
func TestConfidentialTransferCreateContextStateAccountRejectsMissingProof(t *testing.T) {
	t.Parallel()
	for _, testCase := range []struct {
		name  string
		proof proofdata.ProofData
	}{
		{"nil interface", nil},
		{"typed nil", (*proofdata.PubkeyValidityProofData)(nil)},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			if _, err := ConfidentialTransferCreateContextStateAccount(
				ctContextSingle, ctAuthority, ctPayer, testCase.proof, 1); err == nil {
				t.Fatal("got nil error, want a missing proof data error")
			}
		})
	}
}

// TestConfidentialTransferCreateContextStateAccountFromRecord checks the
// variant reading the proof out of a record program account
func TestConfidentialTransferCreateContextStateAccountFromRecord(t *testing.T) {
	t.Parallel()
	const rentLamports = uint64(7_654_321)

	instructions, err := ConfidentialTransferCreateContextStateAccountFromRecord(
		ctContextRange, ctAuthority, ctPayer, ctRecord,
		proofdata.ProofTypeBatchedRangeProofU128, rentLamports)
	assertNoErr(t, err)
	if len(instructions) != 2 {
		t.Fatalf("got %d instructions, want 2", len(instructions))
	}

	accountCreationInstructionData, err := instructions[0].Data()
	assertNoErr(t, err)
	decodedAccountCreationInstructionData, err := system.DecodeInstruction(instructions[0].Accounts(), accountCreationInstructionData)
	assertNoErr(t, err)
	accountCreationData, ok := decodedAccountCreationInstructionData.Impl.(*system.CreateAccount)
	if !ok {
		t.Fatalf("first instruction is %T, want system.CreateAccount", decodedAccountCreationInstructionData.Impl)
	}
	space, err := zkprogram.ProofContextSize(proofdata.ProofTypeBatchedRangeProofU128)
	assertNoErr(t, err)
	if *accountCreationData.Lamports != rentLamports || *accountCreationData.Space != space || !accountCreationData.Owner.Equals(zkprogram.ProgramID) {
		t.Errorf("create account lamports/space/owner = (%d, %d, %s), want (%d, %d, %s)",
			*accountCreationData.Lamports, *accountCreationData.Space, accountCreationData.Owner, rentLamports, space, zkprogram.ProgramID)
	}
	if decodedPayer := accountCreationData.GetFundingAccount().PublicKey; !decodedPayer.Equals(ctPayer) {
		t.Errorf("funding account = %s, want %s", decodedPayer, ctPayer)
	}
	if decodedContextAccount := accountCreationData.GetNewAccount().PublicKey; !decodedContextAccount.Equals(ctContextRange) {
		t.Errorf("new account = %s, want %s", decodedContextAccount, ctContextRange)
	}

	if decodedProgramID := instructions[1].ProgramID(); !decodedProgramID.Equals(zkprogram.ProgramID) {
		t.Fatalf("verify program ID = %s, want %s", decodedProgramID, zkprogram.ProgramID)
	}
	verifyInstructionData, err := instructions[1].Data()
	assertNoErr(t, err)
	if got := zkprogram.ProofInstruction(verifyInstructionData[0]); got != zkprogram.VerifyBatchedRangeProofU128 {
		t.Errorf("proof instruction = %s, want %s", got, zkprogram.VerifyBatchedRangeProofU128)
	}
	wantAccounts := []solana.PublicKey{ctRecord, ctContextRange, ctAuthority}
	verifyInstructionAccounts := instructions[1].Accounts()
	if len(verifyInstructionAccounts) != len(wantAccounts) {
		t.Fatalf("verify accounts = %v, want %v", verifyInstructionAccounts, wantAccounts)
	}
	for i, account := range wantAccounts {
		if !verifyInstructionAccounts[i].PublicKey.Equals(account) {
			t.Errorf("verify account %d = %s, want %s", i, verifyInstructionAccounts[i].PublicKey, account)
		}
	}
	if !verifyInstructionAccounts[1].IsWritable {
		t.Error("context state account is not writable")
	}
	if len(verifyInstructionData) != 5 || binary.LittleEndian.Uint32(verifyInstructionData[1:]) != recordAccountDataOffset {
		t.Errorf("verify instruction data = % x, want a %d proof account offset", verifyInstructionData, recordAccountDataOffset)
	}
}

// ctcOperation describes one proof-carrying operation to the table tests
type ctcOperation struct {
	name           string
	subInstruction uint8
	verifiers      []zkprogram.ProofInstruction
	accounts       []solana.PublicKey
	build          func(inline bool) ([]solana.Instruction, error)
	offsets        func(ConfidentialTransferSubInstructionData) []int8
	checkInline    func(t *testing.T, data ConfidentialTransferSubInstructionData)
	checkAccounts  func(t *testing.T, data ConfidentialTransferSubInstructionData)
}

// ctcOperations builds the shared table of proof-carrying operations over one
// freshly generated set of keys and account states.
func ctcOperations(t *testing.T) []ctcOperation {
	t.Helper()
	kp, aesKey := ctcKeys(t)
	spender := ctcState(t, kp, aesKey, ctcAvailableBalance, 0)
	empty := ctcState(t, kp, aesKey, 0, 0)
	destination, _ := ctcKeys(t)
	auditor, _ := ctcKeys(t)
	withheld, _ := ctcKeys(t)

	// account yields the context state account of one proof, or nil when the
	// proof is to be generated and inlined instead.
	account := func(inline bool, context *solana.PublicKey) *solana.PublicKey {
		if inline {
			return nil
		}
		return context
	}
	validity := func(inline bool) *ProofAccountWithCiphertext {
		if inline {
			return nil
		}
		return &ProofAccountWithCiphertext{
			ContextStateAccount: ctContextValidity,
			CiphertextLo:        ctCiphertextLo,
			CiphertextHi:        ctCiphertextHi,
		}
	}

	return []ctcOperation{
		{
			name:           "ConfigureAccount",
			subInstruction: ConfidentialTransfer_ConfigureAccount,
			verifiers:      []zkprogram.ProofInstruction{zkprogram.VerifyPubkeyValidity},
			accounts:       []solana.PublicKey{ctContextSingle},
			build: func(inline bool) ([]solana.Instruction, error) {
				return ConfidentialTransferConfigureTokenAccount(
					ctTokenAccount, ctMint, ctAuthority, nil, account(inline, &ctContextSingle), nil, kp, aesKey)
			},
			offsets: func(data ConfidentialTransferSubInstructionData) []int8 {
				return []int8{data.(*ConfidentialTransferConfigureAccountData).ProofInstructionOffset}
			},
			checkInline: func(t *testing.T, data ConfidentialTransferSubInstructionData) {
				configure := data.(*ConfidentialTransferConfigureAccountData)
				ctcDecryptable(t, aesKey, configure.DecryptableZeroBalance, 0, "decryptable zero balance")
			},
		},
		{
			name:           "EmptyAccount",
			subInstruction: ConfidentialTransfer_EmptyAccount,
			verifiers:      []zkprogram.ProofInstruction{zkprogram.VerifyZeroCiphertext},
			accounts:       []solana.PublicKey{ctContextSingle},
			build: func(inline bool) ([]solana.Instruction, error) {
				return ConfidentialTransferEmptyAccount(
					ctTokenAccount, ctAuthority, nil, account(inline, &ctContextSingle), NewEmptyAccountInfo(empty), kp)
			},
			offsets: func(data ConfidentialTransferSubInstructionData) []int8 {
				return []int8{data.(*ConfidentialTransferEmptyAccountData).ProofInstructionOffset}
			},
		},
		{
			name:           "Withdraw",
			subInstruction: ConfidentialTransfer_Withdraw,
			verifiers: []zkprogram.ProofInstruction{
				zkprogram.VerifyCiphertextCommitmentEquality,
				zkprogram.VerifyBatchedRangeProofU64,
			},
			accounts: []solana.PublicKey{ctContextEquality, ctContextRange},
			build: func(inline bool) ([]solana.Instruction, error) {
				return ConfidentialTransferWithdraw(
					ctTokenAccount, ctMint, ctAuthority, nil,
					account(inline, &ctContextEquality), account(inline, &ctContextRange),
					ctcTransferAmount, ctDecimals, NewWithdrawAccountInfo(spender), kp, aesKey)
			},
			offsets: func(data ConfidentialTransferSubInstructionData) []int8 {
				withdraw := data.(*ConfidentialTransferWithdrawData)
				return []int8{withdraw.EqualityProofInstructionOffset, withdraw.RangeProofInstructionOffset}
			},
			checkInline: func(t *testing.T, data ConfidentialTransferSubInstructionData) {
				withdraw := data.(*ConfidentialTransferWithdrawData)
				if withdraw.Amount != ctcTransferAmount || withdraw.Decimals != ctDecimals {
					t.Errorf("amount/decimals = (%d, %d), want (%d, %d)",
						withdraw.Amount, withdraw.Decimals, ctcTransferAmount, ctDecimals)
				}
				ctcDecryptable(t, aesKey, withdraw.NewDecryptableAvailableBalance,
					ctcAvailableBalance-ctcTransferAmount, "new decryptable available balance")
			},
		},
		{
			name:           "Transfer",
			subInstruction: ConfidentialTransfer_Transfer,
			verifiers: []zkprogram.ProofInstruction{
				zkprogram.VerifyCiphertextCommitmentEquality,
				zkprogram.VerifyBatchedGroupedCiphertext3HandlesValidity,
				zkprogram.VerifyBatchedRangeProofU128,
			},
			accounts: []solana.PublicKey{ctContextEquality, ctContextValidity, ctContextRange},
			build: func(inline bool) ([]solana.Instruction, error) {
				return ConfidentialTransferTransfer(
					ctTokenAccount, ctDestination, ctMint, ctAuthority, nil,
					account(inline, &ctContextEquality), validity(inline), account(inline, &ctContextRange),
					ctcTransferAmount, NewTransferAccountInfo(spender), kp, aesKey,
					destination.Pubkey, &auditor.Pubkey)
			},
			offsets: func(data ConfidentialTransferSubInstructionData) []int8 {
				transfer := data.(*ConfidentialTransferTransferData)
				return []int8{
					transfer.EqualityProofInstructionOffset,
					transfer.CiphertextValidityProofInstructionOffset,
					transfer.RangeProofInstructionOffset,
				}
			},
			checkInline: func(t *testing.T, data ConfidentialTransferSubInstructionData) {
				transfer := data.(*ConfidentialTransferTransferData)
				ctcDecryptable(t, aesKey, transfer.NewSourceDecryptableAvailableBalance,
					ctcAvailableBalance-ctcTransferAmount, "new source decryptable available balance")
				lo, err := auditor.DecryptU32(transfer.TransferAmountAuditorCiphertextLo)
				assertNoErr(t, err)
				hi, err := auditor.DecryptU32(transfer.TransferAmountAuditorCiphertextHi)
				assertNoErr(t, err)
				if got := lo + hi<<confidential.AmountLoBitLength; got != ctcTransferAmount {
					t.Errorf("auditor ciphertexts decrypt to %d, want %d", got, ctcTransferAmount)
				}
			},
			checkAccounts: func(t *testing.T, data ConfidentialTransferSubInstructionData) {
				transfer := data.(*ConfidentialTransferTransferData)
				if transfer.TransferAmountAuditorCiphertextLo != ctCiphertextLo ||
					transfer.TransferAmountAuditorCiphertextHi != ctCiphertextHi {
					t.Error("auditor ciphertexts are not the ones the proof account was given")
				}
			},
		},
		{
			name:           "TransferWithFee",
			subInstruction: ConfidentialTransfer_TransferWithFee,
			verifiers: []zkprogram.ProofInstruction{
				zkprogram.VerifyCiphertextCommitmentEquality,
				zkprogram.VerifyBatchedGroupedCiphertext3HandlesValidity,
				zkprogram.VerifyPercentageWithCap,
				zkprogram.VerifyBatchedGroupedCiphertext2HandlesValidity,
				zkprogram.VerifyBatchedRangeProofU256,
			},
			accounts: []solana.PublicKey{
				ctContextEquality, ctContextValidity, ctContextFeeSigma,
				ctContextFeeValidity, ctContextRange,
			},
			build: func(inline bool) ([]solana.Instruction, error) {
				return ConfidentialTransferTransferWithFee(
					ctTokenAccount, ctDestination, ctMint, ctAuthority, nil,
					account(inline, &ctContextEquality), validity(inline),
					account(inline, &ctContextFeeSigma), account(inline, &ctContextFeeValidity),
					account(inline, &ctContextRange),
					ctcTransferAmount, NewTransferAccountInfo(spender), kp, aesKey,
					destination.Pubkey, &auditor.Pubkey, withheld.Pubkey, ctcFeeRate, ctcMaximumFee)
			},
			offsets: func(data ConfidentialTransferSubInstructionData) []int8 {
				transfer := data.(*ConfidentialTransferTransferWithFeeData)
				return []int8{
					transfer.EqualityProofInstructionOffset,
					transfer.TransferAmountCiphertextValidityProofInstructionOffset,
					transfer.FeeSigmaProofInstructionOffset,
					transfer.FeeCiphertextValidityProofInstructionOffset,
					transfer.RangeProofInstructionOffset,
				}
			},
		},
	}
}

const (
	ctcAvailableBalance = uint64(1_000_000)
	ctcTransferAmount   = uint64(4_321)
	ctcFeeRate          = uint16(250)
	ctcMaximumFee       = uint64(50)
)

// assertNoErr fails the test on err.
func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// ctcKeys generates the ElGamal keypair and AE key of one confidential account.
func ctcKeys(t *testing.T) (*encryption.ElGamalKeypair, zkencryption.AeKey) {
	t.Helper()
	kp, err := encryption.NewElGamalKeypair()
	assertNoErr(t, err)
	aesKey, err := encryption.NewAeKey()
	assertNoErr(t, err)
	return kp, aesKey
}

// ctcState builds the confidential transfer account state of an account
// holding available as its available balance and pending as its pending one.
func ctcState(
	t *testing.T, kp *encryption.ElGamalKeypair, aesKey zkencryption.AeKey, available, pending uint64,
) *ConfidentialTransferAccountState {
	t.Helper()
	encrypt := func(amount uint64) encryption.ElGamalCiphertext {
		ciphertext, err := kp.Pubkey.Encrypt(amount)
		assertNoErr(t, err)
		return ciphertext
	}
	decryptable, err := encryption.AeEncrypt(aesKey, available)
	assertNoErr(t, err)
	var pendingCredits uint64
	if pending > 0 {
		pendingCredits = 1
	}
	return &ConfidentialTransferAccountState{
		Approved:                    true,
		ElGamalPubkey:               kp.Pubkey,
		AvailableBalance:            encrypt(available),
		DecryptableAvailableBalance: decryptable,
		PendingBalanceLo:            encrypt(pending & (1<<confidential.AmountLoBitLength - 1)),
		PendingBalanceHi:            encrypt(pending >> confidential.AmountLoBitLength),
		PendingBalanceCreditCounter: pendingCredits,
	}
}

// ctcExtensionInstruction decodes a ConfidentialTransfer extension instruction
// back into its sub-instruction ID and typed data.
func ctcExtensionInstruction(t *testing.T, instruction solana.Instruction) (uint8, ConfidentialTransferSubInstructionData) {
	t.Helper()
	if got := instruction.ProgramID(); !got.Equals(ProgramID) {
		t.Fatalf("program ID = %s, want %s", got, ProgramID)
	}
	data, err := instruction.Data()
	assertNoErr(t, err)
	if len(data) < 2 || data[0] != Instruction_ConfidentialTransferExtension {
		t.Fatalf("instruction data % x is not a ConfidentialTransfer extension instruction", data)
	}
	wrapper := ConfidentialTransferExtension{SubInstruction: data[1], RawData: data[2:]}
	subInstructionData, err := wrapper.DecodeSubInstructionData()
	assertNoErr(t, err)
	return wrapper.SubInstruction, subInstructionData
}

// ctcDecryptable fails unless ct decrypts to want under aesKey.
func ctcDecryptable(t *testing.T, aesKey zkencryption.AeKey, ct encryption.AeCiphertext, want uint64, what string) {
	t.Helper()
	got, err := encryption.AeDecrypt(aesKey, ct)
	assertNoErr(t, err)
	if got != want {
		t.Errorf("%s decrypts to %d, want %d", what, got, want)
	}
}

func ctcHasAccount(instruction solana.Instruction, want solana.PublicKey) bool {
	for _, account := range instruction.Accounts() {
		if account.PublicKey.Equals(want) {
			return true
		}
	}
	return false
}
