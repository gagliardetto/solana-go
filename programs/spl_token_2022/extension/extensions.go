package extension

import (
	"errors"
	. "github.com/donutnomad/solana-web3/spl_token_2022"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/cpi_guard"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/default_account_state"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/group_member_pointer"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/group_pointer"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/interest_bearing_mint"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/memo_transfer"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/metadata_pointer"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/token_group"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/transfer_fee"
	"github.com/donutnomad/solana-web3/spl_token_2022/extension/transfer_hook"
	"github.com/donutnomad/solana-web3/token_metadata"
	"github.com/donutnomad/solana-web3/web3"
	"reflect"
)

type GroupMemberPointerParams struct {
	Initialize *group_member_pointer.Initialize
	Update     *group_member_pointer.Update
}

func NewGroupMemberPointerParamsUpdate(
	memberAddress web3.PublicKey,
	mint web3.PublicKey,
	authority web3.PublicKey) GroupMemberPointerParams {
	update := group_member_pointer.NewUpdateInstruction(memberAddress, mint, authority)
	return GroupMemberPointerParams{Update: update}
}

func NewGroupMemberPointerParamsInitialize(authority web3.PublicKey,
	memberAddress web3.PublicKey,
	mint web3.PublicKey) GroupMemberPointerParams {
	initialize := group_member_pointer.NewInitializeInstruction(authority, memberAddress, mint)
	return GroupMemberPointerParams{Initialize: initialize}
}

func (m GroupMemberPointerParams) ExtensionType() ExtensionType {
	return ExtensionTypeGroupMemberPointer
}

type GroupPointerParams struct {
	Initialize *group_pointer.Initialize
	Update     *group_pointer.Update
}

func NewGroupPointerParamsUpdate(groupAddress web3.PublicKey,
	mint web3.PublicKey,
	authority web3.PublicKey) GroupPointerParams {
	update := group_pointer.NewUpdateInstruction(groupAddress, mint, authority)
	return GroupPointerParams{Update: update}
}

func NewGroupPointerParamsInitialize(authority web3.PublicKey,
	groupAddress web3.PublicKey,
	mint web3.PublicKey) GroupPointerParams {
	initialize := group_pointer.NewInitializeInstruction(authority, groupAddress, mint)
	return GroupPointerParams{Initialize: initialize}
}

func (m GroupPointerParams) ExtensionType() ExtensionType {
	return ExtensionTypeGroupPointer
}

type CpiGuardParams struct {
	Enable  bool
	Account web3.PublicKey
	Owner   web3.PublicKey
}

func NewCpiGuardParams(enable bool, account web3.PublicKey, owner web3.PublicKey) CpiGuardParams {
	return CpiGuardParams{Enable: enable, Account: account, Owner: owner}
}

func (m CpiGuardParams) ExtensionType() ExtensionType {
	return ExtensionTypeCpiGuard
}

type NonTransferableParams struct {
	Mint web3.PublicKey
}

func NewNonTransferableParams(mint web3.PublicKey) NonTransferableParams {
	return NonTransferableParams{Mint: mint}
}

func (m NonTransferableParams) ExtensionType() ExtensionType {
	return ExtensionTypeNonTransferable
}

type ImmutableOwnerParams struct {
	Account web3.PublicKey
}

func NewImmutableOwnerParams(account web3.PublicKey) ImmutableOwnerParams {
	return ImmutableOwnerParams{Account: account}
}

func (m ImmutableOwnerParams) ExtensionType() ExtensionType {
	return ExtensionTypeImmutableOwner
}

type MemoTransferParams struct {
	Enable       bool
	Account      web3.PublicKey
	AccountOwner web3.PublicKey
	MultiSigners []web3.PublicKey
}

func NewMemoTransferParams(enable bool, account web3.PublicKey, accountOwner web3.PublicKey, multiSigners ...web3.PublicKey) MemoTransferParams {
	return MemoTransferParams{Enable: enable, Account: account, AccountOwner: accountOwner, MultiSigners: multiSigners}
}

func (m MemoTransferParams) ExtensionType() ExtensionType {
	return ExtensionTypeMemoTransfer
}

type DefaultAccountStateParams struct {
	State AccountState
}

func (p DefaultAccountStateParams) ExtensionType() ExtensionType {
	return ExtensionTypeDefaultAccountState
}

func NewDefaultAccountStateParams(state AccountState) DefaultAccountStateParams {
	return DefaultAccountStateParams{
		State: state,
	}
}

type MintCloseAuthorityParams struct {
	CloseAuthority *web3.PublicKey
}

func (p MintCloseAuthorityParams) ExtensionType() ExtensionType {
	return ExtensionTypeMintCloseAuthority
}

func NewMintCloseAuthorityParams(closeAuthority *web3.PublicKey) MintCloseAuthorityParams {
	return MintCloseAuthorityParams{
		CloseAuthority: closeAuthority,
	}
}

type TransferFeeConfigParams struct {
	ExtensionTyp               ExtensionType
	TransferFeeConfigAuthority *web3.PublicKey
	WithdrawWithheldAuthority  *web3.PublicKey
	TransferFeeBasisPoints     uint16
	MaximumFee                 uint64
}

func (p TransferFeeConfigParams) ExtensionType() ExtensionType {
	return ExtensionTypeTransferFeeConfig
}

func NewTransferFeeConfigParams(
	transferFeeBasisPoints uint16,
	maximumFee uint64,
	transferFeeConfigAuthority, withdrawWithheldAuthority *web3.PublicKey,
) TransferFeeConfigParams {
	return TransferFeeConfigParams{
		TransferFeeConfigAuthority: transferFeeConfigAuthority,
		WithdrawWithheldAuthority:  withdrawWithheldAuthority,
		TransferFeeBasisPoints:     transferFeeBasisPoints,
		MaximumFee:                 maximumFee,
	}
}

type InterestBearingConfigParams struct {
	RateAuthority web3.PublicKey
	Rate          int16
}

func (p InterestBearingConfigParams) ExtensionType() ExtensionType {
	return ExtensionTypeInterestBearingConfig
}

func NewInterestBearingConfigParams(rateAuthority *web3.PublicKey, rate int16) InterestBearingConfigParams {
	o := InterestBearingConfigParams{}
	if rateAuthority != nil {
		o.RateAuthority = *rateAuthority
	}
	o.Rate = rate
	return o
}

type PermanentDelegateParams struct {
	Delegate web3.PublicKey
}

func (p PermanentDelegateParams) ExtensionType() ExtensionType {
	return ExtensionTypePermanentDelegate
}

func NewPermanentDelegateParams(delegate web3.PublicKey) PermanentDelegateParams {
	return PermanentDelegateParams{
		Delegate: delegate,
	}
}

type TransferHookParams struct {
	Initialize *transfer_hook.Initialize
	Update     *transfer_hook.Update
}

func NewTransferHookParamsUpdate(programId web3.PublicKey,
	mint web3.PublicKey,
	authority web3.PublicKey) TransferHookParams {
	update := transfer_hook.NewUpdateInstruction(programId, mint, authority)
	return TransferHookParams{Update: update}
}

func NewTransferHookParamsInitialize(authority web3.PublicKey,
	programId web3.PublicKey,
	mint web3.PublicKey) TransferHookParams {
	initialize := transfer_hook.NewInitializeInstruction(authority, programId, mint)
	return TransferHookParams{Initialize: initialize}
}

func (p TransferHookParams) ExtensionType() ExtensionType {
	return ExtensionTypeTransferHookAccount
}

type MetadataPointerParams struct {
	Initialize *metadata_pointer.Initialize
	Update     *metadata_pointer.Update
}

func NewMetadataPointerParamsUpdate(
	metadataAddress web3.PublicKey,
	mint web3.PublicKey,
	owner web3.PublicKey,
) MetadataPointerParams {
	update := metadata_pointer.NewUpdateInstruction(metadataAddress, mint, owner)
	return MetadataPointerParams{Update: update}
}

func NewMetadataPointerParamsInitialize(
	authority *web3.PublicKey,
	metadataAddress *web3.PublicKey,
	mint web3.PublicKey) MetadataPointerParams {
	initialize := metadata_pointer.NewInitializeInstruction(nilDef(authority), nilDef(metadataAddress), mint)
	return MetadataPointerParams{Initialize: initialize}
}

func nilDef[T any](a *T) T {
	if a == nil {
		var ret T
		return ret
	}
	return *a
}

func (p MetadataPointerParams) ExtensionType() ExtensionType {
	return ExtensionTypeMetadataPointer
}

type ExtensionInitializationParams interface {
	ExtensionType() ExtensionType
}

func Nested[T interface {
	SetProgramId(key *web3.PublicKey) T
	SetAccounts(accounts []*web3.AccountMeta) error
}, T2 interface {
	Validate() error
}](programId *web3.PublicKey, innerBuilder T2, next func(data []byte) T) (web3.Instruction, error) {
	inner, err := adapterInstructionTo(innerBuilder)
	if err != nil {
		return nil, err
	}
	data, err := inner.Data()
	if err != nil {
		return nil, err
	}
	ret := next(data)
	err = ret.SetProgramId(programId).
		SetAccounts(inner.Accounts())
	if err != nil {
		return nil, err
	}
	return adapterInstructionTo(ret)
}

var InvalidParams = errors.New("invalid params")

func ExtensionInitializationParamsToInstruction(
	extension ExtensionInitializationParams,
	mint web3.PublicKey,
	programId web3.PublicKey,
) (web3.Instruction, error) {
	switch extension.ExtensionType() {
	case ExtensionTypeTransferFeeConfig:
		p1 := extension.(TransferFeeConfigParams)
		return Nested(
			&programId,
			transfer_fee.NewInitializeTransferFeeConfigInstruction(
				p1.TransferFeeConfigAuthority.Ref(),
				p1.WithdrawWithheldAuthority.Ref(),
				p1.TransferFeeBasisPoints,
				p1.MaximumFee,
				mint,
			),
			NewTransferFeeExtensionInstruction,
		)
	case ExtensionTypeMintCloseAuthority:
		p1 := extension.(MintCloseAuthorityParams)
		return NewInitializeMintCloseAuthorityInstruction(p1.CloseAuthority, mint).SetProgramId(&programId).ValidateAndBuild()
	case ExtensionTypeDefaultAccountState:
		p1 := extension.(DefaultAccountStateParams)
		return Nested(
			&programId,
			default_account_state.NewInitializeInstruction(p1.State, mint),
			NewDefaultAccountStateExtensionInstruction,
		)
	case ExtensionTypeImmutableOwner:
		p1 := extension.(ImmutableOwnerParams)
		return NewInitializeImmutableOwnerInstruction(p1.Account).SetProgramId(&programId).ValidateAndBuild()
	case ExtensionTypeMemoTransfer:
		p1 := extension.(MemoTransferParams)
		if p1.Enable {
			return Nested(
				&programId,
				memo_transfer.NewEnableInstruction(p1.Account, p1.AccountOwner).SetAccountOwnerAccount(p1.AccountOwner, p1.MultiSigners...),
				NewMemoTransferExtensionInstruction,
			)
		} else {
			return Nested(
				&programId,
				memo_transfer.NewDisableInstruction(p1.Account, p1.AccountOwner).SetAccountOwnerAccount(p1.AccountOwner, p1.MultiSigners...),
				NewMemoTransferExtensionInstruction,
			)
		}
	case ExtensionTypeNonTransferable:
		p1 := extension.(NonTransferableParams)
		return NewInitializeNonTransferableMintInstruction(p1.Mint).SetProgramId(&programId).ValidateAndBuild()
	case ExtensionTypeInterestBearingConfig:
		p1 := extension.(InterestBearingConfigParams)
		return Nested(
			&programId,
			interest_bearing_mint.NewInitializeInstruction(
				p1.RateAuthority,
				p1.Rate,
				mint,
			),
			NewInterestBearingMintExtensionInstruction,
		)
	case ExtensionTypeCpiGuard:
		p1 := extension.(CpiGuardParams)
		if p1.Enable {
			return Nested(
				&programId,
				cpi_guard.NewEnableInstruction(p1.Account, p1.Owner),
				NewCpiGuardExtensionInstruction,
			)
		} else {
			return Nested(
				&programId,
				cpi_guard.NewDisableInstruction(p1.Account, p1.Owner),
				NewCpiGuardExtensionInstruction,
			)
		}
	case ExtensionTypePermanentDelegate:
		p1 := extension.(PermanentDelegateParams)
		return NewInitializePermanentDelegateInstruction(p1.Delegate, mint).SetProgramId(&programId).ValidateAndBuild()
	case ExtensionTypeTransferHook:
		p1 := extension.(TransferHookParams)
		if p1.Initialize != nil {
			return Nested(&programId, p1.Initialize, NewTransferHookExtensionInstruction)
		} else if p1.Update != nil {
			return Nested(&programId, p1.Update, NewTransferHookExtensionInstruction)
		} else {
			return nil, InvalidParams
		}
	case ExtensionTypeTransferFeeAmount:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeNonTransferableAccount:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeTransferHookAccount:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeConfidentialTransferMint:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeConfidentialTransferAccount:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeConfidentialTransferFeeConfig:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeConfidentialTransferFeeAmount:
		// this extension is not a instruction
		return nil, nil
	case ExtensionTypeGroupPointer:
		p1 := extension.(GroupPointerParams)
		if p1.Initialize != nil {
			return Nested(&programId, p1.Initialize, NewGroupPointerExtensionInstruction)
		} else if p1.Update != nil {
			return Nested(&programId, p1.Update, NewGroupPointerExtensionInstruction)
		} else {
			return nil, InvalidParams
		}
	case ExtensionTypeGroupMemberPointer:
		p1 := extension.(GroupMemberPointerParams)
		if p1.Initialize != nil {
			return Nested(&programId, p1.Initialize, NewGroupMemberPointerExtensionInstruction)
		} else if p1.Update != nil {
			return Nested(&programId, p1.Update, NewGroupMemberPointerExtensionInstruction)
		} else {
			return nil, InvalidParams
		}
	case ExtensionTypeMetadataPointer:
		p1 := extension.(MetadataPointerParams)
		if p1.Initialize != nil {
			return Nested(&programId, p1.Initialize, NewMetadataPointerExtensionInstruction)
		} else if p1.Update != nil {
			return Nested(&programId, p1.Update, NewMetadataPointerExtensionInstruction)
		} else {
			return nil, InvalidParams
		}
	case ExtensionTypeTokenMetadata:
		// ignore, use alone --> see
		_ = token_metadata.Initialize{}
		_ = token_metadata.UpdateField{}
		_ = token_metadata.RemoveKey{}
		_ = token_metadata.UpdateAuthority{}
		_ = token_metadata.Emit{}
		return nil, nil
	case ExtensionTypeTokenGroupMember:
		// ignore, use alone --> see
		_ = token_group.InitializeMember{}
		return nil, nil
	case ExtensionTypeTokenGroup:
		// ignore, use alone --> see
		_ = token_group.Initialize{}
		_ = token_group.UpdateGroupMaxSize{}
		_ = token_group.UpdateGroupAuthority{}
		return nil, nil
	default:
		return nil, nil
	}
}

func ExistsExtensionType(array []ExtensionInitializationParams, extensionType ExtensionType) bool {
	for _, item := range array {
		if item.ExtensionType() == extensionType {
			return true
		}
	}
	return false
}

func validateAndBuild[T any](target any) (*T, error) {
	var invalidMethod = errors.New("method ValidateAndBuild is invalid")
	method := reflect.ValueOf(target).MethodByName("ValidateAndBuild")
	if !method.IsValid() {
		return nil, invalidMethod
	}
	results := method.Call(nil)
	if len(results) != 2 {
		return nil, invalidMethod
	}
	errValue := results[1]
	if !errValue.IsNil() {
		err, ok := errValue.Interface().(error)
		if !ok {
			err = invalidMethod
		}
		return nil, err
	}
	instruction, ok := results[0].Interface().(T)
	if !ok {
		return nil, invalidMethod
	}
	return &instruction, nil
}

func adapterInstructionTo(target any) (web3.Instruction, error) {
	ins, err := validateAndBuild[web3.Instruction](target)
	if err == nil {
		return *ins, nil
	}
	return nil, err
}
