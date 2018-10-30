// Copyright (c) 2018 The copernet developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

import (
	"encoding/json"
	"strconv"
)

type WhcGetInfoResultResult struct {
	VersionInt     int32       `json:"wormholeversion_int"`
	Version        string      `json:"wormholeversion"`
	BitcoinVersion string      `json:"bitcoincoreversion"`
	Block          int         `json:"block"`
	BlockTime      int64       `json:"blocktime"`
	Txs            int         `json:"blocktransactions"`
	TotalTrades    int         `json:"totaltrades"`
	TotalTxs       int         `json:"totaltransactions"`
	Alerts         []AlertInfo `json:"alerts"`
}

type AlertInfo struct {
	TypeInt uint64 `json:"alerttypeint"`
	TypeStr string `json:"alerttype"`
	Expire  string `json:"alertexpiry"`
	Message string `json:"alertmessage"`
}

type WhcActiveCrowdResult struct {
	PropertyID  int     `json:"propertyid"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Subcategory string  `json:"subcategory"`
	Data        string  `json:"data"`
	Url         string  `json:"url"`
	Precision   int8    `json:"property precision"`
	Issuer      string  `json:"issuer"`
	Txid        string  `json:"creationtxid"`
	TotalTokens float64 `json:"totaltokens"`
}

type WhcGetAllBalanceForAddressResult struct {
	PropertyID uint64 `json:"propertyid"`
	Balance    string `json:"balance"`
	Reserved   string `json:"reserved"`
}

type WhcGetAllBalanceForIDResult struct {
	Address  string `json:"address"`
	Balance  string `json:"balance"`
	Reserved string `json:"reserved"`
}

type WhcGetBalanceResult struct {
	Balance  string `json:"balance"`
	Reserved string `json:"reserved"`
}

type WhcGetBalanceHashResult struct {
	Block       uint32 `json:"block"`
	BlockHash   string `json:"blockhash"`
	PropertyID  uint64 `json:"propertyid"`
	BalanceHash string `json:"balanceshash"`
}

type ParticipantTransactions struct {
	TxID              string `json:"txid"`
	Amount            string `json:"amountsent"`
	ParticipantTokens string `json:"participanttokens"`
}

type WhcGetCrowdSaleResult struct {
	PropertyId        uint64                    `json:"propertyid"`
	Name              string                    `json:"name"`
	Active            bool                      `json:"active"`
	Issuer            string                    `json:"issuer"`
	PropertyIdDesired uint64                    `json:"propertyiddesired"`
	Precision         string                    `json:"precision"`
	TokensPerUnit     string                    `json:"tokensperunit"`
	EarlyBonus        uint8                     `json:"earlybonus"`
	StartTime         int64                     `json:"starttime"`
	DeadLine          int64                     `json:"deadline"`
	AmountRaised      string                    `json:"amountraised"`
	TokensIssued      string                    `json:"tokensissued"`
	AddedIssuerTokens string                    `json:"addedissuertokens"`
	CloseDearly       *bool                     `json:"closedearly,omitempty"`
	MaxTokens         *bool                     `json:"omitempty,omniempty"`
	EndedTime         int64                     `json:"omitempty,omitempty"`
	TxidClosed        *string                   `json:"closetx,omniempty"`
	ParticipantTxs    []ParticipantTransactions `json:"participanttransactions,omniempty"`
}

type WhcGetCurrentConsensusHashResult struct {
	Block         uint32 `json:"block"`
	BlockHash     string `json:"blockhash"`
	ConsensusHash string `json:"consensushash"`
}

type WhcGetFeeShareResult struct {
	Address  string `json:"address"`
	FeeShare string `json:"feeshare"`
}

type IssuanceTxs struct {
	TxID   string `json:"txid"`
	Grant  string `json:"grant"`
	Revoke string `json:"revoke"`
}

func (i *IssuanceTxs) MarshalJSON() ([]byte, error) {
	var grant, revoke int
	var err error
	if i.Grant != "" {
		grant, err = strconv.Atoi(i.Grant)
		if err != nil {
			return nil, err
		}
	}

	if i.Revoke != "" {
		revoke, err = strconv.Atoi(i.Revoke)
		if err != nil {
			return nil, err
		}
	}

	if grant > 0 {
		grantResult := struct {
			TxID  string `json:"txid"`
			Grant string `json:"grant"`
		}{
			TxID:  i.TxID,
			Grant: i.Grant,
		}

		return json.Marshal(grantResult)
	}

	if revoke > 0 {
		revokeResult := struct {
			TxID   string `json:"txid"`
			Revoke string `json:"revoke"`
		}{
			TxID:   i.TxID,
			Revoke: i.Revoke,
		}

		return json.Marshal(revokeResult)
	}

	return nil, nil
}

type WhcGetGrantsResult struct {
	PropertyID  uint64        `json:"propertyid"`
	Name        string        `json:"name"`
	Issuer      string        `json:"issuer"`
	CreateTxID  string        `json:"creationtxid"`
	TotalTokens string        `json:"totaltokens"`
	Issuance    []IssuanceTxs `json:"issuances"`
}

type WhcGetPayloadResult struct {
	Payload     string `json:"payload"`
	PayloadSize int    `json:"payloadsize"`
}

type WhcGetPropertyResult struct {
	PropertyID      uint64 `json:"propertyid"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	Subcategory     string `json:"subcategory"`
	Data            string `json:"data"`
	Url             string `json:"url"`
	Precision       int    `json:"precision"`
	Issuer          string `json:"issuer"`
	CreateTxID      string `json:"creationtxid"`
	FixedIssuance   bool   `json:"fixedissuance"`
	ManagedIssuance bool   `json:"managedissuance"`
	FreezingEnabled *bool  `json:"freezingenabled"`
	TotalTokens     string `json:"totaltokens"`
}

type PurchasesDetails struct {
	Vout             uint64 `json:"vout"`
	AmountPaid       string `json:"amountpaid"`
	IsMine           bool   `json:"ismine"`
	ReferenceAddress string `json:"referenceaddress"`
	PropertyID       uint64 `json:"propertyid"`
	AmountBought     string `json:"amountbought"`
	Valid            bool   `json:"valid"`
}

type WhcListPropertiesResult struct {
	PropertyID  uint64 `json:"propertyid"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Data        string `json:"data"`
	Url         string `json:"url"`
	Precision   int    `json:"precision"`
}

type GenerateTransactionResult struct {
	// --------- required fields
	TxID             string `json:"txid"`
	Fee              string `json:"fee"`
	SendingAddress   string `json:"sendingaddress"`
	ReferenceAddress string `json:"referenceaddress,omitempty"`
	IsMine           bool   `json:"ismine"`
	Version          uint64 `json:"version"`
	TypeInt          uint64 `json:"type_int"`
	Type             string `json:"type,omitempty"`
	BlockHeight      uint32 `json:"block,omitempty"`
	Confirmations    int    `json:"confirmations,omitempty"`

	Valid         bool   `json:"valid,omitempty"`
	InvalidReason string `json:"invalidreason,omitempty"`
	BlockHash     string `json:"blockhash,omitempty"`
	BlockTime     int64  `json:"blocktime,omitempty"`
	PositionBlock int    `json:"positionblock,omitempty"`
	// ----------

	PropertyID int64  `json:"propertyid,omitempty"`
	Precision  string `json:"precision,omitempty"`
	Amount     string `json:"amount,omitempty"`

	TotalStoFee string      `json:"totalstofee,omitempty"`
	Recipients  []Recipient `json:"recipients,omitempty"`

	Ecosystem string    `json:"ecosystem,omitempty"`
	SubSends  []SubSend `json:"subsends,omitemtpy"`

	Mature bool   `json:"mature,omitempty"`
	Enough string `json:"Enough,omitempty"`
	Burn   string `json:"burns,omitempty"`

	ActualInvested             string `json:"actualinvested,omitempty"`
	PurchasedPropertyID        int64  `json:"purchasedpropertyid,omitempty"`
	PurchasedPropertyName      string `json:"purchasedpropertyname,omitempty"`
	PurchasedPropertyPrecision string `json:"purchasedpropertyprecision,omitempty"`
	PurchasedTokens            string `json:"purchasedtokens,omitempty"`
	IssuerTokens               string `json:"issuertokens,omitempty"`

	Category     string `json:"category,omitempty"`
	Subcategory  string `json:"subcategory,omitempty"`
	PropertyName string `json:"propertyname,omitempty"`
	Data         string `json:"data,omitempty"`
	Url          string `json:"url,omitempty"`

	PropertyIDDesired int64  `json:"propertyiddesired,omitempty"`
	TokensPerUnit     string `json:"tokensperunit,omitempty"`
	Deadline          int64  `json:"deadline,omitempty"`
	EarlyBonus        uint8  `json:"earlybonus,omitempty"`
}

type Recipient struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type SubSend struct {
	PropertyID int64  `json:"propertyid,omitempty"`
	Precision  string `json:"precision,omitempty"`
	Amount     string `json:"amount,omitempty"`
}
