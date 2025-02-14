// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_vault

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type InitVaultArgs struct {
	AllowFurtherShareCreation bool
}

func (obj InitVaultArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AllowFurtherShareCreation` param:
	err = encoder.Encode(obj.AllowFurtherShareCreation)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitVaultArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AllowFurtherShareCreation`:
	err = decoder.Decode(&obj.AllowFurtherShareCreation)
	if err != nil {
		return err
	}
	return nil
}

type AmountArgs struct {
	Amount uint64
}

func (obj AmountArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmountArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type NumberOfShareArgs struct {
	NumberOfShares uint64
}

func (obj NumberOfShareArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NumberOfShares` param:
	err = encoder.Encode(obj.NumberOfShares)
	if err != nil {
		return err
	}
	return nil
}

func (obj *NumberOfShareArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NumberOfShares`:
	err = decoder.Decode(&obj.NumberOfShares)
	if err != nil {
		return err
	}
	return nil
}

type MintEditionProxyArgs struct {
	Edition uint64
}

func (obj MintEditionProxyArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Edition` param:
	err = encoder.Encode(obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MintEditionProxyArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Edition`:
	err = decoder.Decode(&obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

type ExternalPriceAccount struct {
	Key           Key
	PricePerShare uint64

	// Mint of the currency we are pricing the shares against, should be same as redeem_treasury.
	// Most likely will be USDC mint most of the time.
	PriceMint ag_solanago.PublicKey

	// Whether or not combination has been allowed for this vault.
	AllowedToCombine bool
}

func (obj ExternalPriceAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	// Serialize `PricePerShare` param:
	err = encoder.Encode(obj.PricePerShare)
	if err != nil {
		return err
	}
	// Serialize `PriceMint` param:
	err = encoder.Encode(obj.PriceMint)
	if err != nil {
		return err
	}
	// Serialize `AllowedToCombine` param:
	err = encoder.Encode(obj.AllowedToCombine)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ExternalPriceAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	// Deserialize `PricePerShare`:
	err = decoder.Decode(&obj.PricePerShare)
	if err != nil {
		return err
	}
	// Deserialize `PriceMint`:
	err = decoder.Decode(&obj.PriceMint)
	if err != nil {
		return err
	}
	// Deserialize `AllowedToCombine`:
	err = decoder.Decode(&obj.AllowedToCombine)
	if err != nil {
		return err
	}
	return nil
}

type Key ag_binary.BorshEnum

const (
	Uninitialized_Key Key = iota
	SafetyDepositBoxV1_Key
	ExternalAccountKeyV1_Key
	VaultV1_Key
)

func (value Key) String() string {
	switch value {
	case Uninitialized_Key:
		return "Uninitialized"
	case SafetyDepositBoxV1_Key:
		return "SafetyDepositBoxV1"
	case ExternalAccountKeyV1_Key:
		return "ExternalAccountKeyV1"
	case VaultV1_Key:
		return "VaultV1"
	default:
		return ""
	}
}

type VaultState ag_binary.BorshEnum

const (
	Inactive_VaultState VaultState = iota
	Active_VaultState
	Combined_VaultState
	Deactivated_VaultState
)

func (value VaultState) String() string {
	switch value {
	case Inactive_VaultState:
		return "Inactive"
	case Active_VaultState:
		return "Active"
	case Combined_VaultState:
		return "Combined"
	case Deactivated_VaultState:
		return "Deactivated"
	default:
		return ""
	}
}
