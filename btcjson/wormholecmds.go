// Copyright (c) 2018 The copernet developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

// WhcGetInfo defines the wh_getinfo JSON-RPC command.
type WhcGetInfoCmd struct{}

// NewWhcGetInfoCmd returns a new instance which can be used to issue a
// whc_getinfo JSON-RPC command.
func NewWhcGetInfoCmd() *WhcGetInfoCmd {
	return &WhcGetInfoCmd{}
}

type WhcSetAutoCommitCmd struct {
	Flag bool
}

func NewWhcSetAutoCommitCmd(flag bool) *WhcSetAutoCommitCmd {
	return &WhcSetAutoCommitCmd{
		Flag: flag,
	}
}

type WhcGetActiveCrowdCmd struct {
	Address string
}

func NewWhcGetActiveCrowdCmd(addr string) *WhcGetActiveCrowdCmd {
	return &WhcGetActiveCrowdCmd{
		Address: addr,
	}
}

type WhcGetAllBalancesForAddressCMd struct {
	Address string
}

func NewWhcGetAllBalancesForAddressCmd(addr string) *WhcGetAllBalancesForAddressCMd {
	return &WhcGetAllBalancesForAddressCMd{
		Address: addr,
	}
}

type WhcGetAllBalanceForIdCmd struct {
	PropertyID uint64
}

func NewWhcGetAllBalanceForIdCmd(id uint64) *WhcGetAllBalanceForIdCmd {
	return &WhcGetAllBalanceForIdCmd{
		PropertyID: id,
	}
}

type WhcGetBalanceCmd struct {
	Address    string
	PropertyID uint64
}

func NewWhcGetBalanceCmd(addr string, id uint64) *WhcGetBalanceCmd {
	return &WhcGetBalanceCmd{
		Address:    addr,
		PropertyID: id,
	}
}

type WhcGetBalanceHashCmd struct {
	PropertyID uint64
}

func NewWhcGetBalanceHashCmd(id uint64) *WhcGetBalanceHashCmd {
	return &WhcGetBalanceHashCmd{
		PropertyID: id,
	}
}

type WhcGetCrowdSaleCmd struct {
	PropertyID uint64
	Verbose    *bool `jsonrpcdefault:"false"`
}

func NewWhcGetCrowdSaleCmd(id uint64, verbose *bool) *WhcGetCrowdSaleCmd {
	return &WhcGetCrowdSaleCmd{
		PropertyID: id,
		Verbose:    verbose,
	}
}

type WhcGetCurrentConsensusHashCmd struct{}

func NewWhcGetCurrentConsensusHashCmd() *WhcGetCurrentConsensusHashCmd {
	return &WhcGetCurrentConsensusHashCmd{}
}

type WhcGetGrantsCmd struct {
	PropertyID uint64
}

func NewWhcGetGrantsCmd(id uint64) *WhcGetGrantsCmd {
	return &WhcGetGrantsCmd{
		PropertyID: id,
	}
}

type WhcGetPayloadCmd struct {
	TxID string
}

func NewWhcGetPayloadCmd(txid string) *WhcGetPayloadCmd {
	return &WhcGetPayloadCmd{
		TxID: txid,
	}
}

type WhcGetPropertyCmd struct {
	PropertyID uint64
}

func NewWhcGetPropertyCmd(id uint64) *WhcGetPropertyCmd {
	return &WhcGetPropertyCmd{
		PropertyID: id,
	}
}

type WhcGetSeedBlocksCmd struct {
	StartBlock uint32
	EndBlock   uint32
}

func NewWhcGetSeedBlocksCmd(start uint32, end uint32) *WhcGetSeedBlocksCmd {
	return &WhcGetSeedBlocksCmd{
		StartBlock: start,
		EndBlock:   end,
	}
}

type WhcGetStoCmd struct {
	TxID            string
	RecipientFilter *string
}

func NewWhcGetStoCmd(txid string, filter *string) *WhcGetStoCmd {
	return &WhcGetStoCmd{
		TxID:            txid,
		RecipientFilter: filter,
	}
}

type WhcGetTransactionCmd struct {
	TxID string
}

func NewWhcGetTransactionCmd(txid string) *WhcGetTransactionCmd {
	return &WhcGetTransactionCmd{
		TxID: txid,
	}
}

type WhcListBlockTransactionsCmd struct {
	Index uint32
}

func NewWhcListBlockTransactionsCmd(index uint32) *WhcListBlockTransactionsCmd {
	return &WhcListBlockTransactionsCmd{
		Index: index,
	}
}

type WhcListPendingTransactionsCmd struct {
	Address *string
}

func NewWhcListPendingTransactionsCmd(addr *string) *WhcListPendingTransactionsCmd {
	return &WhcListPendingTransactionsCmd{
		Address: addr,
	}
}

type WhcLIstPropertiesCmd struct{}

func NewWhcLIstPropertiesCmd() *WhcLIstPropertiesCmd {
	return &WhcLIstPropertiesCmd{}
}

type WhcListTransactionsCmd struct {
	Address    *string
	Count      *int64
	Skip       *int64
	StartBlock *int64
	EndBlock   *int64
}

func NewWhcListTransactionsCmd(addr *string, count, skip, startblock, endblock *int64) *WhcListTransactionsCmd {
	return &WhcListTransactionsCmd{
		Address:    addr,
		Count:      count,
		Skip:       skip,
		StartBlock: startblock,
		EndBlock:   endblock,
	}
}

type WhcCreatePayloadBurnBCHCmd struct{}

func NewWhcCreatePayloadBurnBCHCmd() *WhcCreatePayloadBurnBCHCmd {
	return &WhcCreatePayloadBurnBCHCmd{}
}

type WhcCreatePayloadChangeIssuerCmd struct {
	PropertyID int64
}

func NewWhcCreatePayloadChangeIssuerCmd(id int64) *WhcCreatePayloadChangeIssuerCmd {
	return &WhcCreatePayloadChangeIssuerCmd{
		PropertyID: id,
	}
}

type WhcCreatePayloadCloseCrowdSaleCmd struct {
	PropertyID int64
}

func NewWhcCreatePayloadCloseCrowdSaleCmd(id int64) *WhcCreatePayloadCloseCrowdSaleCmd {
	return &WhcCreatePayloadCloseCrowdSaleCmd{
		PropertyID: id,
	}
}

type WhcCreatePayloadGrantCmd struct {
	PropertyID int64
	Amount     string
	Note       *string
}

func NewWhcCreatePayloadGrantCmd(id int64, amount string, note *string) *WhcCreatePayloadGrantCmd {
	return &WhcCreatePayloadGrantCmd{
		PropertyID: id,
		Amount:     amount,
		Note:       note,
	}
}

type WhcCreatePayloadIssuanceCrowdSaleCmd struct {
	Ecosystem         int64
	Precision         int64
	PreviousID        int64
	Category          string
	SubCategory       string
	Name              string
	Url               string
	Data              string
	PropertyIdDesired int64
	TokensPerUnit     string
	DeadLine          int64
	EarlyBonus        int64
	IssuerPercentage  int64
	Amount            string
}

func NewWhcCreatePayloadIssuanceCrowdSaleCmd(eco, precision, preId, desiredID, deadline, earlyBonus,
issuerPercentage int64, category, subcategory, name, url, data string, tokensPreUnit, amount string) *WhcCreatePayloadIssuanceCrowdSaleCmd {
	return &WhcCreatePayloadIssuanceCrowdSaleCmd{
		Ecosystem:         eco,
		Precision:         precision,
		PreviousID:        preId,
		Category:          category,
		SubCategory:       subcategory,
		Name:              name,
		Url:               url,
		Data:              data,
		PropertyIdDesired: desiredID,
		TokensPerUnit:     tokensPreUnit,
		DeadLine:          deadline,
		EarlyBonus:        earlyBonus,
		IssuerPercentage:  issuerPercentage,
		Amount:            amount,
	}
}

type WhcCreatePayloadIssuanceFixedCmd struct {
	Ecosystem   int64
	Precision   int64
	PreviousID  int64
	Category    string
	SubCategory string
	Name        string
	Url         string
	Data        string
	Amount      string
}

func NewWhcCreatePayloadIssuanceFixedCmd(eco, precision, preId int64, category, subcategory,
name, url, data, amount string) *WhcCreatePayloadIssuanceFixedCmd {
	return &WhcCreatePayloadIssuanceFixedCmd{
		Ecosystem:   eco,
		Precision:   precision,
		PreviousID:  preId,
		Category:    category,
		SubCategory: subcategory,
		Name:        name,
		Url:         url,
		Data:        data,
		Amount:      amount,
	}
}

type WhcCreatePayloadIssuanceManagedCmd struct {
	Ecosystem   int64
	Precision   int64
	PreviousID  int64
	Category    string
	SubCategory string
	Name        string
	Url         string
	Data        string
}

func NewWhcCreatePayloadIssuanceManagedCmd(eco, precision, preId int64, category, subcategory,
name, url, data string) *WhcCreatePayloadIssuanceManagedCmd {

	return &WhcCreatePayloadIssuanceManagedCmd{
		Ecosystem:   eco,
		Precision:   precision,
		PreviousID:  preId,
		Category:    category,
		SubCategory: subcategory,
		Name:        name,
		Url:         url,
		Data:        data,
	}
}

type WhcCreatePayloadPartiCrowdSaleCmd struct {
	Amount string
}

func NewWhcCreatePayloadPartiCrowdSaleCmd(amount string) *WhcCreatePayloadPartiCrowdSaleCmd {
	return &WhcCreatePayloadPartiCrowdSaleCmd{
		Amount: amount,
	}
}

type WhcCreatePayloadRevokeCmd struct {
	PropertyID int64
	Amount     string
	Note       *string
}

func NewWhcCreatePayloadRevokeCmd(id int64, amount string, note *string) *WhcCreatePayloadRevokeCmd {
	return &WhcCreatePayloadRevokeCmd{
		PropertyID: id,
		Amount:     amount,
		Note:       note,
	}
}

type WhcCreatePayloadSendAllCmd struct {
	Ecosystem int64
}

func NewWhcCreatePayloadSendAllCmd(eco int64) *WhcCreatePayloadSendAllCmd {
	return &WhcCreatePayloadSendAllCmd{
		Ecosystem: eco,
	}
}

type WhcCreatePayloadSimpleSendCmd struct {
	PropertyID int64
	Amount     string
}

func NewWhcCreatePayloadSimpleSendCmd(id int64, amount string) *WhcCreatePayloadSimpleSendCmd {
	return &WhcCreatePayloadSimpleSendCmd{
		PropertyID: id,
		Amount:     amount,
	}
}

type WhcCreatePayloadStoCmd struct {
	PropertyID int64
	Amount     string
	DisID      *int64
}

func NewWhcCreatePayloadStoCmd(fromID int64, amount string, toID *int64) *WhcCreatePayloadStoCmd {
	return &WhcCreatePayloadStoCmd{
		PropertyID: fromID,
		Amount:     amount,
		DisID:      toID,
	}
}

type PrevTx struct {
	TxID         string  `json:"txid"`
	Vout         int     `json:"vout"`
	ScriptPubKey string  `json:"scriptPubKey"`
	Value        float64 `json:"value"`
}

type WhcCreateRawTxChangeCmd struct {
	RawTX       string
	PrevTxs     []PrevTx
	Destination string
	Fee         string
	Position    *uint32
}

func NewWhcCreateRawTxChangeCmd(rawtx string, prev []PrevTx, address, fee string, position *uint32) *WhcCreateRawTxChangeCmd {
	return &WhcCreateRawTxChangeCmd{
		RawTX:       rawtx,
		PrevTxs:     prev,
		Destination: address,
		Fee:         fee,
		Position:    position,
	}
}

type WhcCreateRawTxInputCmd struct {
	RawTX string
	TxID  string
	Index int
}

func NewWhcCreateRawTxInputCmd(rawtx, txid string, index int) *WhcCreateRawTxInputCmd {
	return &WhcCreateRawTxInputCmd{
		RawTX: rawtx,
		TxID:  txid,
		Index: index,
	}
}

type WhcCreateRawTxOpReturnCmd struct {
	RawTx   string
	Payload string
}

func NewWhcCreateRawTxOpReturnCmd(rawtx, payload string) *WhcCreateRawTxOpReturnCmd {
	return &WhcCreateRawTxOpReturnCmd{
		RawTx:   rawtx,
		Payload: payload,
	}
}

type WhcCreateRawTxReferenceCmd struct {
	RawTx   string
	Address string
	Amount  *string
}

func NewWhcCreateRawTxReferenceCmd(rawtx, address string, amount *string) *WhcCreateRawTxReferenceCmd {
	return &WhcCreateRawTxReferenceCmd{
		RawTx:   rawtx,
		Address: address,
		Amount:  amount,
	}
}

type WhcDecodeTransactionCmd struct {
	RawTx   string
	PrevTxs *[]PrevTx
	Height  *int
}

func NewWhcDecodeTransactionCmd(rawtx string, prev *[]PrevTx, height *int) *WhcDecodeTransactionCmd {
	return &WhcDecodeTransactionCmd{
		RawTx:   rawtx,
		PrevTxs: prev,
		Height:  height,
	}
}

type WhcBurnBCHGetWhcCmd struct {
	Amount        string
	RedeemAddress *string
}

func NewWhcBurnBCHGetWhcCmd(amount string, redeemAddress *string) *WhcBurnBCHGetWhcCmd {
	return &WhcBurnBCHGetWhcCmd{
		Amount:        amount,
		RedeemAddress: redeemAddress,
	}
}

type WhcPartiCrowdSaleCmd struct {
	FromAddress     string
	ToAddress       string
	Amount          string
	RedeemAddress   *string
	ReferenceAmount *string
}

func NewWhcPartiCrowdSaleCmd(from, to, amount string, redeemAddress, referenceAmount *string) *WhcPartiCrowdSaleCmd {
	return &WhcPartiCrowdSaleCmd{
		FromAddress:     from,
		ToAddress:       to,
		Amount:          amount,
		RedeemAddress:   redeemAddress,
		ReferenceAmount: referenceAmount,
	}
}

type WhcSendCmd struct {
	FromAddress     string
	ToAddress       string
	PropertyID      uint64
	Amount          string
	RedeemAddress   *string
	ReferenceAmount *string
}

func NewWhcSendCmd(from, to string, id uint64, amount string, redeem, referenceAmount *string) *WhcSendCmd {
	return &WhcSendCmd{
		FromAddress:     from,
		ToAddress:       to,
		PropertyID:      id,
		Amount:          amount,
		RedeemAddress:   redeem,
		ReferenceAmount: referenceAmount,
	}
}

type WhcSendAllCmd struct {
	FromAddress     string
	ToAddress       string
	Ecosystem       int64
	RedeemAddress   *string
	ReferenceAmount *string
}

func NewWhcSendAllCmd(from, to string, eco int64, redeem, referenceAmount *string) *WhcSendAllCmd {
	return &WhcSendAllCmd{
		FromAddress:     from,
		ToAddress:       to,
		Ecosystem:       eco,
		RedeemAddress:   redeem,
		ReferenceAmount: referenceAmount,
	}
}

type WhcSendChangeIssuerCmd struct {
	FromAddress string
	ToAddress   string
	PropertyID  uint64
}

func NewWhcSendChangeIssuerCmd(from, to string, id uint64) *WhcSendChangeIssuerCmd {
	return &WhcSendChangeIssuerCmd{
		FromAddress: from,
		ToAddress:   to,
		PropertyID:  id,
	}
}

type WhcSendCloseCrowdSaleCmd struct {
	FromAddress string
	PropertyID  uint64
}

func NewWhcSendCloseCrowdSaleCmd(from string, id uint64) *WhcSendCloseCrowdSaleCmd {
	return &WhcSendCloseCrowdSaleCmd{
		FromAddress: from,
		PropertyID:  id,
	}
}

type WhcSendGrantCmd struct {
	FromAddress string
	ToAddress   string
	PropertyID  uint64
	Amount      string
	Note        *string
}

func NewWhcSendGrantCmd(from, to string, id uint64, amount string, note *string) *WhcSendGrantCmd {
	return &WhcSendGrantCmd{
		FromAddress: from,
		ToAddress:   to,
		PropertyID:  id,
		Amount:      amount,
		Note:        note,
	}
}

type WhcSendIssuanceCrowdSaleCmd struct {
	FromAddress       string
	Ecosystem         int64
	Precision         int64
	PreviousID        int64
	Category          string
	SubCategory       string
	Name              string
	Url               string
	Data              string
	PropertyIdDesired int64
	TokensPerUnit     string
	DeadLine          int64
	EarlyBonus        int64
	IssuerPercentage  int64
	Amount            string
}

func NewWhcSendIssuanceCrowdSaleCmd(from string, eco, precision, previousid int64, category,
subCategory, name, url, data string, desiredID int64, tokensPerUnit string, deadline,
earlyBonus, issuerPercentage int64, amount string) *WhcSendIssuanceCrowdSaleCmd {

	return &WhcSendIssuanceCrowdSaleCmd{
		FromAddress:       from,
		Ecosystem:         eco,
		Precision:         precision,
		PreviousID:        previousid,
		Category:          category,
		SubCategory:       subCategory,
		Name:              name,
		Url:               url,
		Data:              data,
		PropertyIdDesired: desiredID,
		TokensPerUnit:     tokensPerUnit,
		DeadLine:          deadline,
		EarlyBonus:        earlyBonus,
		IssuerPercentage:  issuerPercentage,
		Amount:            amount,
	}
}

type WhcSendIssuanceFixedCmd struct {
	FromAddress string
	Ecosystem   int64
	Precision   int64
	PreviousID  int64
	Category    string
	SubCategory string
	Name        string
	Url         string
	Data        string
	TotalNumber string
}

func NewWhcSendIssuanceFixedCmd(from string, eco, precision, previousid int64, category,
subCategory, name, url, data, total string) *WhcSendIssuanceFixedCmd {

	return &WhcSendIssuanceFixedCmd{
		FromAddress: from,
		Ecosystem:   eco,
		Precision:   precision,
		PreviousID:  previousid,
		Category:    category,
		SubCategory: subCategory,
		Name:        name,
		Url:         url,
		Data:        data,
		TotalNumber: total,
	}
}

type WhcSendIssuanceManagedCmd struct {
	FromAddress string
	Ecosystem   int64
	Precision   int64
	PreviousID  int64
	Category    string
	SubCategory string
	Name        string
	Url         string
	Data        string
}

func NewWhcSendIssuanceManagedCmd(from string, eco, precision, previousid int64, category,
subCategory, name, url, data string) *WhcSendIssuanceManagedCmd {

	return &WhcSendIssuanceManagedCmd{
		FromAddress: from,
		Ecosystem:   eco,
		Precision:   precision,
		PreviousID:  previousid,
		Category:    category,
		SubCategory: subCategory,
		Name:        name,
		Url:         url,
		Data:        data,
	}
}

type WhcSendRawTxCmd struct {
	FromAddress      string
	RawTransaction   string
	ReferenceAddress *string
	RedeemAddress    *string
	ReferenceAmount  *string
}

func NewWhcSendRawTxCmd(from, rawtx string, referenceAddress, redeem, referenceAmount *string) *WhcSendRawTxCmd {
	return &WhcSendRawTxCmd{
		FromAddress:      from,
		RawTransaction:   rawtx,
		ReferenceAmount:  referenceAmount,
		RedeemAddress:    redeem,
		ReferenceAddress: referenceAddress,
	}
}

type WhcSendRevokeCmd struct {
	FromAddress string
	PropertyID  int64
	Amount      string
	Note        *string
}

func NewWhcSendRevokeCmd(from string, id int64, amount string, note *string) *WhcSendRevokeCmd {
	return &WhcSendRevokeCmd{
		FromAddress: from,
		PropertyID:  id,
		Amount:      amount,
		Note:        note,
	}
}

type WhcSendStoCmd struct {
	FromAddress          string
	PropertyID           int64
	Amount               string
	RedeemAddress        *string
	DistributionProperty *int64
}

func NewWhcSendStoCmd(from string, id int64, amount string, redeem *string, distributionProperty *int64) *WhcSendStoCmd {
	return &WhcSendStoCmd{
		FromAddress:          from,
		PropertyID:           id,
		Amount:               amount,
		RedeemAddress:        redeem,
		DistributionProperty: distributionProperty,
	}
}

type WhcFreezeCmd struct {
	FromAddress   string
	PropertyID    int64
	Amount        string
	FrozenAddress string
}

func NewWhcSendFreezeCmd(from string, id int64, amount string, frozenAddress string) *WhcFreezeCmd {
	return &WhcFreezeCmd{
		FromAddress:   from,
		PropertyID:    id,
		Amount:        amount,
		FrozenAddress: frozenAddress,
	}
}

type WhcUnFreezeCmd struct {
	FromAddress   string
	PropertyID    int64
	Amount        string
	FrozenAddress string
}

func NewWhcSendUnFreezeCmd(from string, id int64, amount string, frozenAddress string) *WhcUnFreezeCmd {
	return &WhcUnFreezeCmd{
		FromAddress:   from,
		PropertyID:    id,
		Amount:        amount,
		FrozenAddress: frozenAddress,
	}
}


type WhcCreatePayloadFreezeCmd struct {
	ToAddress  string
	PropertyID int64
	Amount     string
}

func NewWhcCreatePayloadFreezeCmd(toAddress string, id int64, amount string) *WhcCreatePayloadFreezeCmd {
	return &WhcCreatePayloadFreezeCmd{
		ToAddress:  toAddress,
		PropertyID: id,
		Amount:     amount,
	}
}

type WhcCreatePayloadUnFreezeCmd struct {
	ToAddress  string
	PropertyID int64
	Amount     string
}

func NewWhcCreatePayloadUnFreezeCmd(toAddress string, id int64, amount string) *WhcCreatePayloadUnFreezeCmd {
	return &WhcCreatePayloadUnFreezeCmd{
		ToAddress:  toAddress,
		PropertyID: id,
		Amount:     amount,
	}
}

type GetFrozenBalanceCmd struct {
	Address  string
	PropertyID int64
}

func NewWhcGetFrozenBalanceCmd(address string, id int64) *GetFrozenBalanceCmd {
	return &GetFrozenBalanceCmd{
		Address:  address,
		PropertyID: id,
	}
}

type WhcGetFrozenBalanceForIdCmd struct {
	PropertyID int64
}

func NewWhcGetFrozenBalanceForIdCmd(id int64) *WhcGetFrozenBalanceForIdCmd {
	return &WhcGetFrozenBalanceForIdCmd{
		PropertyID: id,
	}
}


type WhcGetFrozenBalanceForAddressCmd struct {
	Address string
}

func NewWhcGetFrozenBalanceForAddressCmd(address string) *WhcGetFrozenBalanceForAddressCmd {
	return &WhcGetFrozenBalanceForAddressCmd{
		Address: address,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	MustRegisterCmd("whc_setautocommit", (*WhcSetAutoCommitCmd)(nil), flags)
	MustRegisterCmd("whc_getactivecrowd", (*WhcGetActiveCrowdCmd)(nil), flags)
	MustRegisterCmd("whc_getallbalancesforaddress", (*WhcGetAllBalancesForAddressCMd)(nil), flags)
	MustRegisterCmd("whc_getallbalancesforid", (*WhcGetAllBalanceForIdCmd)(nil), flags)
	MustRegisterCmd("whc_getbalance", (*WhcGetBalanceCmd)(nil), flags)
	MustRegisterCmd("whc_getbalanceshash", (*WhcGetBalanceHashCmd)(nil), flags)
	MustRegisterCmd("whc_getcrowdsale", (*WhcGetCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_getcurrentconsensushash", (*WhcGetCurrentConsensusHashCmd)(nil), flags)
	MustRegisterCmd("whc_getgrants", (*WhcGetGrantsCmd)(nil), flags)
	MustRegisterCmd("whc_getinfo", (*WhcGetInfoCmd)(nil), flags)
	MustRegisterCmd("whc_getpayload", (*WhcGetPayloadCmd)(nil), flags)
	MustRegisterCmd("whc_getproperty", (*WhcGetPropertyCmd)(nil), flags)
	MustRegisterCmd("whc_getseedblocks", (*WhcGetSeedBlocksCmd)(nil), flags)
	MustRegisterCmd("whc_getsto", (*WhcGetStoCmd)(nil), flags)
	MustRegisterCmd("whc_gettransaction", (*WhcGetTransactionCmd)(nil), flags)
	MustRegisterCmd("whc_listblocktransactions", (*WhcListBlockTransactionsCmd)(nil), flags)
	MustRegisterCmd("whc_listpendingtransactions", (*WhcListPendingTransactionsCmd)(nil), flags)
	MustRegisterCmd("whc_listproperties", (*WhcLIstPropertiesCmd)(nil), flags)
	MustRegisterCmd("whc_listtransactions", (*WhcListTransactionsCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_burnbch", (*WhcCreatePayloadBurnBCHCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_changeissuer", (*WhcCreatePayloadChangeIssuerCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_closecrowdsale", (*WhcCreatePayloadCloseCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_grant", (*WhcCreatePayloadGrantCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_issuancecrowdsale", (*WhcCreatePayloadIssuanceCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_issuancefixed", (*WhcCreatePayloadIssuanceFixedCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_issuancemanaged", (*WhcCreatePayloadIssuanceManagedCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_particrowdsale", (*WhcCreatePayloadPartiCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_revoke", (*WhcCreatePayloadRevokeCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_sendall", (*WhcCreatePayloadSendAllCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_simplesend", (*WhcCreatePayloadSimpleSendCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_sto", (*WhcCreatePayloadStoCmd)(nil), flags)
	MustRegisterCmd("whc_createrawtx_change", (*WhcCreateRawTxChangeCmd)(nil), flags)
	MustRegisterCmd("whc_createrawtx_input", (*WhcCreateRawTxInputCmd)(nil), flags)
	MustRegisterCmd("whc_createrawtx_opreturn", (*WhcCreateRawTxOpReturnCmd)(nil), flags)
	MustRegisterCmd("whc_createrawtx_reference", (*WhcCreateRawTxReferenceCmd)(nil), flags)
	MustRegisterCmd("whc_decodetransaction", (*WhcDecodeTransactionCmd)(nil), flags)
	MustRegisterCmd("whc_burnbchgetwhc", (*WhcBurnBCHGetWhcCmd)(nil), flags)
	MustRegisterCmd("whc_particrowsale", (*WhcPartiCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_send", (*WhcSendCmd)(nil), flags)
	MustRegisterCmd("whc_sendall", (*WhcSendAllCmd)(nil), flags)
	MustRegisterCmd("whc_sendchangeissuer", (*WhcSendChangeIssuerCmd)(nil), flags)
	MustRegisterCmd("whc_sendclosecrowdsale", (*WhcSendCloseCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_sendgrant", (*WhcSendGrantCmd)(nil), flags)
	MustRegisterCmd("whc_sendissuancecrowdsale", (*WhcSendIssuanceCrowdSaleCmd)(nil), flags)
	MustRegisterCmd("whc_sendissuancefixed", (*WhcSendIssuanceFixedCmd)(nil), flags)
	MustRegisterCmd("whc_sendissuancemanaged", (*WhcSendIssuanceManagedCmd)(nil), flags)
	MustRegisterCmd("whc_sendrawtx", (*WhcSendRawTxCmd)(nil), flags)
	MustRegisterCmd("whc_sendrevoke", (*WhcSendRevokeCmd)(nil), flags)
	MustRegisterCmd("whc_sendsto", (*WhcSendStoCmd)(nil), flags)
	MustRegisterCmd("whc_sendfreeze", (*WhcFreezeCmd)(nil), flags)
	MustRegisterCmd("whc_sendunfreeze", (*WhcUnFreezeCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_freeze", (*WhcCreatePayloadFreezeCmd)(nil), flags)
	MustRegisterCmd("whc_createpayload_unfreeze", (*WhcCreatePayloadUnFreezeCmd)(nil), flags)
	MustRegisterCmd("whc_getfrozenbalance", (*GetFrozenBalanceCmd)(nil), flags)
	MustRegisterCmd("whc_getfrozenbalanceforid", (*WhcGetFrozenBalanceForIdCmd)(nil), flags)
	MustRegisterCmd("whc_getfrozenbalanceforaddress", (*WhcGetFrozenBalanceForAddressCmd)(nil), flags)
}
