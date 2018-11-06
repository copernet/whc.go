// Copyright (c) 2018 The copernet developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpcclient

import (
	"encoding/json"

	"github.com/copernet/whc.go/btcjson"
)

type FutureWhcGetInfoResult chan *response

func (r FutureWhcGetInfoResult) Receive() (*btcjson.WhcGetInfoResultResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var whcInfoResult btcjson.WhcGetInfoResultResult
	err = json.Unmarshal(res, &whcInfoResult)
	if err != nil {
		return nil, err
	}

	return &whcInfoResult, nil
}

func (c *Client) WhcGetInfoAsync() FutureWhcGetInfoResult {
	cmd := btcjson.NewWhcGetInfoCmd()
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetInfo() (*btcjson.WhcGetInfoResultResult, error) {
	return c.WhcGetInfoAsync().Receive()
}

type FutureWhcSetAutoCommitResult chan *response

func (r FutureWhcSetAutoCommitResult) Receive() (bool, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return false, err
	}

	var result bool
	err = json.Unmarshal(res, &result)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (c *Client) WhcSetAutoCommitAsync(flag bool) FutureWhcSetAutoCommitResult {
	cmd := btcjson.NewWhcSetAutoCommitCmd(flag)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSetAutoCommit(flag bool) (bool, error) {
	return c.WhcSetAutoCommitAsync(flag).Receive()
}

type FutureWhcGetActiveCrowdResult chan *response

func (r FutureWhcGetActiveCrowdResult) Receive() (*btcjson.WhcActiveCrowdResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var activeCrowdResult btcjson.WhcActiveCrowdResult
	err = json.Unmarshal(res, &activeCrowdResult)
	if err != nil {
		return nil, err
	}

	return &activeCrowdResult, nil
}

func (c *Client) WhcGetActiveCrowdAsync(addr string) FutureWhcGetActiveCrowdResult {
	cmd := btcjson.NewWhcGetActiveCrowdCmd(addr)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetActiveCrowd(addr string) (*btcjson.WhcActiveCrowdResult, error) {
	return c.WhcGetActiveCrowdAsync(addr).Receive()
}

type FutureWhcGetAllBalancesForAddressResult chan *response

func (r FutureWhcGetAllBalancesForAddressResult) Receive() ([]btcjson.WhcGetAllBalanceForAddressResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var balances []btcjson.WhcGetAllBalanceForAddressResult
	err = json.Unmarshal(res, &balances)
	if err != nil {
		return nil, err
	}

	return balances, nil
}

func (c *Client) WhcGetAllBalancesForAddressAsync(addr string) FutureWhcGetAllBalancesForAddressResult {
	cmd := btcjson.NewWhcGetAllBalancesForAddressCmd(addr)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetAllBalancesForAddress(addr string) ([]btcjson.WhcGetAllBalanceForAddressResult, error) {
	return c.WhcGetAllBalancesForAddressAsync(addr).Receive()
}

type FutureWhcGetAllBalancesForIdResult chan *response

func (r FutureWhcGetAllBalancesForIdResult) Receive() ([]btcjson.WhcGetAllBalanceForIDResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var balances []btcjson.WhcGetAllBalanceForIDResult
	err = json.Unmarshal(res, &balances)
	if err != nil {
		return nil, err
	}

	return balances, nil
}

func (c *Client) WhcGetAllBalancesForIdAsync(id uint64) FutureWhcGetAllBalancesForIdResult {
	cmd := btcjson.NewWhcGetAllBalanceForIdCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetAllBalancesForID(id uint64) ([]btcjson.WhcGetAllBalanceForIDResult, error) {
	return c.WhcGetAllBalancesForIdAsync(id).Receive()
}

type FutureWhcGetBalancesResult chan *response

func (r FutureWhcGetBalancesResult) Receive() (*btcjson.WhcGetBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var balances btcjson.WhcGetBalanceResult
	err = json.Unmarshal(res, &balances)
	if err != nil {
		return nil, err
	}

	return &balances, nil
}

func (c *Client) WhcGetAllBalancesAsync(addr string, id uint64) FutureWhcGetBalancesResult {
	cmd := btcjson.NewWhcGetBalanceCmd(addr, id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetBalances(addr string, id uint64) (*btcjson.WhcGetBalanceResult, error) {
	return c.WhcGetAllBalancesAsync(addr, id).Receive()
}

type FutureWhcGetBalancesHashResult chan *response

func (r FutureWhcGetBalancesHashResult) Receive() (*btcjson.WhcGetBalanceHashResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var balances btcjson.WhcGetBalanceHashResult
	err = json.Unmarshal(res, &balances)
	if err != nil {
		return nil, err
	}

	return &balances, nil
}

func (c *Client) WhcGetAllBalanceHashAsync(id uint64) FutureWhcGetBalancesHashResult {
	cmd := btcjson.NewWhcGetBalanceHashCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetBalanceHash(id uint64) (*btcjson.WhcGetBalanceHashResult, error) {
	return c.WhcGetAllBalanceHashAsync(id).Receive()
}

type FutureWhcGetCrowdSaleResult chan *response

func (r FutureWhcGetCrowdSaleResult) Receive() (*btcjson.WhcGetCrowdSaleResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.WhcGetCrowdSaleResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetCrowdSaleAsync(id uint64, v *bool) FutureWhcGetCrowdSaleResult {
	cmd := btcjson.NewWhcGetCrowdSaleCmd(id, v)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetCrowdSale(id uint64, v *bool) (*btcjson.WhcGetCrowdSaleResult, error) {
	return c.WhcGetCrowdSaleAsync(id, v).Receive()
}

type FutureWhcGetCurrentConsensusHashResult chan *response

func (r FutureWhcGetCurrentConsensusHashResult) Receive() (*btcjson.WhcGetCurrentConsensusHashResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.WhcGetCurrentConsensusHashResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetCurrentConsensusHashAsync() FutureWhcGetCurrentConsensusHashResult {
	cmd := btcjson.NewWhcGetCurrentConsensusHashCmd()
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetCurrentConsensusHash() (*btcjson.WhcGetCurrentConsensusHashResult, error) {
	return c.WhcGetCurrentConsensusHashAsync().Receive()
}

type FutureWhcGetGrantsResult chan *response

func (r FutureWhcGetGrantsResult) Receive() (*btcjson.WhcGetGrantsResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.WhcGetGrantsResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetGrantsAsync(id uint64) FutureWhcGetGrantsResult {
	cmd := btcjson.NewWhcGetGrantsCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetGrants(id uint64) (*btcjson.WhcGetGrantsResult, error) {
	return c.WhcGetGrantsAsync(id).Receive()
}

type FutureWhcGetPayloadResult chan *response

func (r FutureWhcGetPayloadResult) Receive() (*btcjson.WhcGetPayloadResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.WhcGetPayloadResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetPayloadAsync(txid string) FutureWhcGetPayloadResult {
	cmd := btcjson.NewWhcGetPayloadCmd(txid)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetPayload(txid string) (*btcjson.WhcGetPayloadResult, error) {
	return c.WhcGetPayloadAsync(txid).Receive()
}

type FutureWhcGetPropertyResult chan *response

func (r FutureWhcGetPropertyResult) Receive() (*btcjson.WhcGetPropertyResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.WhcGetPropertyResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetPropertyAsync(id uint64) FutureWhcGetPropertyResult {
	cmd := btcjson.NewWhcGetPropertyCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetProperty(id uint64) (*btcjson.WhcGetPropertyResult, error) {
	return c.WhcGetPropertyAsync(id).Receive()
}

type FutureWhcGetSeedBlocksResult chan *response

func (r FutureWhcGetSeedBlocksResult) Receive() ([]uint32, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []uint32
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) WhcGetSeedBlocksAsync(start uint32, end uint32) FutureWhcGetSeedBlocksResult {
	cmd := btcjson.NewWhcGetSeedBlocksCmd(start, end)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetSeedBlocks(start uint32, end uint32) ([]uint32, error) {
	return c.WhcGetSeedBlocksAsync(start, end).Receive()
}

type FutureWhcGetStoResult chan *response

func (r FutureWhcGetStoResult) Receive() (*btcjson.GenerateTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.GenerateTransactionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetStoAsync(txid string, filter *string) FutureWhcGetStoResult {
	cmd := btcjson.NewWhcGetStoCmd(txid, filter)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetSto(txid string, filter *string) (*btcjson.GenerateTransactionResult, error) {
	return c.WhcGetStoAsync(txid, filter).Receive()
}

type FutureWhcGetTransactionResult chan *response

func (r FutureWhcGetTransactionResult) Receive() (*btcjson.GenerateTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.GenerateTransactionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetTransactionAsync(txid string) FutureWhcGetTransactionResult {
	cmd := btcjson.NewWhcGetTransactionCmd(txid)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetTransaction(txid string) (*btcjson.GenerateTransactionResult, error) {
	return c.WhcGetTransactionAsync(txid).Receive()
}

type FutureWhcListBlockTransactionsResult chan *response

func (r FutureWhcListBlockTransactionsResult) Receive() ([]string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) WhcListBlockTransactionsAsync(index uint32) FutureWhcListBlockTransactionsResult {
	cmd := btcjson.NewWhcListBlockTransactionsCmd(index)
	return c.sendCmd(cmd)
}

func (c *Client) WhcListBlockTransactions(index uint32) ([]string, error) {
	return c.WhcListBlockTransactionsAsync(index).Receive()
}

type FutureWhcListPendingTransactionsResult chan *response

func (r FutureWhcListPendingTransactionsResult) Receive() ([]btcjson.GenerateTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []btcjson.GenerateTransactionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) WhcListPendingTransactionsAsync(addr *string) FutureWhcListPendingTransactionsResult {
	cmd := btcjson.NewWhcListPendingTransactionsCmd(addr)
	return c.sendCmd(cmd)
}

func (c *Client) WhcListPendingTransactions(addr *string) ([]btcjson.GenerateTransactionResult, error) {
	return c.WhcListPendingTransactionsAsync(addr).Receive()
}

type FutureWhcListPropertiesResult chan *response

func (r FutureWhcListPropertiesResult) Receive() ([]btcjson.WhcListPropertiesResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []btcjson.WhcListPropertiesResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) WhcListPropertiesAsync() FutureWhcListPropertiesResult {
	cmd := btcjson.NewWhcLIstPropertiesCmd()
	return c.sendCmd(cmd)
}

func (c *Client) WhcListProperties() ([]btcjson.WhcListPropertiesResult, error) {
	return c.WhcListPropertiesAsync().Receive()
}

type FutureWhcListTransactionsResult chan *response

func (r FutureWhcListTransactionsResult) Receive() ([]btcjson.GenerateTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []btcjson.GenerateTransactionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) WhcListTransactionsAsync(addr *string, count, skip, startblock,
endblock *int64) FutureWhcListTransactionsResult {

	cmd := btcjson.NewWhcListTransactionsCmd(addr, count, skip, startblock, endblock)
	return c.sendCmd(cmd)
}

func (c *Client) WhcListTransactions(addr *string, count, skip, startblock,
endblock *int64) ([]btcjson.GenerateTransactionResult, error) {

	return c.WhcListTransactionsAsync(addr, count, skip, startblock, endblock).Receive()
}

type FutureWhcCreatePayloadBurnBCHResult chan *response

func (r FutureWhcCreatePayloadBurnBCHResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadBurnBCHAsync() FutureWhcCreatePayloadBurnBCHResult {
	cmd := btcjson.NewWhcCreatePayloadBurnBCHCmd()
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadBurnBCH() (string, error) {
	return c.WhcCreatePayloadBurnBCHAsync().Receive()
}

type FutureWhcCreatePayloadChangeIssuerResult chan *response

func (r FutureWhcCreatePayloadChangeIssuerResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadChangeIssuerAsync(id int64) FutureWhcCreatePayloadChangeIssuerResult {
	cmd := btcjson.NewWhcCreatePayloadChangeIssuerCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadChangeIssuer(id int64) (string, error) {
	return c.WhcCreatePayloadChangeIssuerAsync(id).Receive()
}

type FutureWhcCreatePayloadCloseCrowdSaleResult chan *response

func (r FutureWhcCreatePayloadCloseCrowdSaleResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadCloseCrowdSaleAsync(id int64) FutureWhcCreatePayloadCloseCrowdSaleResult {
	cmd := btcjson.NewWhcCreatePayloadCloseCrowdSaleCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadCloseCrowdSale(id int64) (string, error) {
	return c.WhcCreatePayloadCloseCrowdSaleAsync(id).Receive()
}

type FutureWhcCreatePayloadGrantResult chan *response

func (r FutureWhcCreatePayloadGrantResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadGrantAsync(id int64, amount string,
	note *string) FutureWhcCreatePayloadGrantResult {

	cmd := btcjson.NewWhcCreatePayloadGrantCmd(id, amount, note)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadGrant(id int64, amount string, note *string) (string, error) {
	return c.WhcCreatePayloadGrantAsync(id, amount, note).Receive()
}

type FutureWhcCreatePayloadIssuanceCrowdSaleResult chan *response

func (r FutureWhcCreatePayloadIssuanceCrowdSaleResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadIssuanceCrowdSaleAsync(eco, precision, preId, desiredID, deadline,
earlyBonus, issuerPercentage int64, category, subcategory, name, url, data string, tokensPreUnit,
amount string) FutureWhcCreatePayloadIssuanceCrowdSaleResult {

	cmd := btcjson.NewWhcCreatePayloadIssuanceCrowdSaleCmd(eco, precision, preId, desiredID, deadline, earlyBonus,
		issuerPercentage, category, subcategory, name, url, data, tokensPreUnit, amount)

	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadIssuanceCrowdSale(eco, precision, preId, desiredID, deadline, earlyBonus,
issuerPercentage int64, category, subcategory, name, url, data string, tokensPreUnit,
amount string) (string, error) {

	return c.WhcCreatePayloadIssuanceCrowdSaleAsync(eco, precision, preId, desiredID, deadline, earlyBonus,
		issuerPercentage, category, subcategory, name, url, data, tokensPreUnit, amount).Receive()
}

type FutureWhcCreatePayloadIssuanceFixedResult chan *response

func (r FutureWhcCreatePayloadIssuanceFixedResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadIssuanceFixedAsync(eco, precision, preId int64, category, subcategory,
name, url, data, amount string) FutureWhcCreatePayloadIssuanceFixedResult {

	cmd := btcjson.NewWhcCreatePayloadIssuanceFixedCmd(eco, precision, preId, category, subcategory,
		name, url, data, amount)

	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadIssuanceFixed(eco, precision, preId int64, category, subcategory,
name, url, data, amount string) (string, error) {

	return c.WhcCreatePayloadIssuanceFixedAsync(eco, precision, preId, category, subcategory, name,
		url, data, amount).Receive()
}

type FutureWhcCreatePayloadIssuanceManagedResult chan *response

func (r FutureWhcCreatePayloadIssuanceManagedResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadIssuanceManagedAsync(eco, precision, preId int64, category, subcategory,
name, url, data string) FutureWhcCreatePayloadIssuanceManagedResult {

	cmd := btcjson.NewWhcCreatePayloadIssuanceManagedCmd(eco, precision, preId, category, subcategory,
		name, url, data)

	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadIssuanceManaged(eco, precision, preId int64, category, subcategory,
name, url, data string) (string, error) {

	return c.WhcCreatePayloadIssuanceManagedAsync(eco, precision, preId, category, subcategory,
		name, url, data).Receive()
}

type FutureWhcCreatePayloadPartiCrowdSaleResult chan *response

func (r FutureWhcCreatePayloadPartiCrowdSaleResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadPartiCrowdSaleAsync(amount string) FutureWhcCreatePayloadPartiCrowdSaleResult {
	cmd := btcjson.NewWhcCreatePayloadPartiCrowdSaleCmd(amount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadPartiCrowdSale(amount string) (string, error) {
	return c.WhcCreatePayloadPartiCrowdSaleAsync(amount).Receive()
}

type FutureWhcCreatePayloadRevokeResult chan *response

func (r FutureWhcCreatePayloadRevokeResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadRevokeAsync(id int64, amount string,
	note *string) FutureWhcCreatePayloadRevokeResult {

	cmd := btcjson.NewWhcCreatePayloadRevokeCmd(id, amount, note)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadRevoke(id int64, amount string, note *string) (string, error) {
	return c.WhcCreatePayloadRevokeAsync(id, amount, note).Receive()
}

type FutureWhcCreatePayloadSendALlResult chan *response

func (r FutureWhcCreatePayloadSendALlResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadSendAllAsync(eco int64) FutureWhcCreatePayloadSendALlResult {
	cmd := btcjson.NewWhcCreatePayloadSendAllCmd(eco)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadSendAll(eco int64) (string, error) {
	return c.WhcCreatePayloadSendAllAsync(eco).Receive()
}

type FutureWhcCreatePayloadSimpleSendResult chan *response

func (r FutureWhcCreatePayloadSimpleSendResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadSimpleSendAsync(id int64,
	amount string) FutureWhcCreatePayloadSimpleSendResult {

	cmd := btcjson.NewWhcCreatePayloadSimpleSendCmd(id, amount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadSimpleSend(id int64, amount string) (string, error) {
	return c.WhcCreatePayloadSimpleSendAsync(id, amount).Receive()
}

type FutureWhcCreatePayloadStoResult chan *response

func (r FutureWhcCreatePayloadStoResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadStoAsync(fromId int64, amount string,
	toID *int64) FutureWhcCreatePayloadStoResult {

	cmd := btcjson.NewWhcCreatePayloadStoCmd(fromId, amount, toID)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadSto(fromId int64, amount string, toID *int64) (string, error) {
	return c.WhcCreatePayloadStoAsync(fromId, amount, toID).Receive()
}

type FutureWhcCreateRawTxChangeResult chan *response

func (r FutureWhcCreateRawTxChangeResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreateRawTxChangeAsync(rawtx string, prev []btcjson.PrevTx, address, fee string, position *uint32) FutureWhcCreateRawTxChangeResult {

	cmd := btcjson.NewWhcCreateRawTxChangeCmd(rawtx, prev, address, fee, position)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreateRawTxChange(rawtx string, prev []btcjson.PrevTx, address, fee string, position *uint32) (string, error) {

	return c.WhcCreateRawTxChangeAsync(rawtx, prev, address, fee, position).Receive()
}

type FutureWhcCreateRawTxInputResult chan *response

func (r FutureWhcCreateRawTxInputResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreateRawTxInputAsync(rawtx, txid string, index int) FutureWhcCreateRawTxInputResult {
	cmd := btcjson.NewWhcCreateRawTxInputCmd(rawtx, txid, index)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreateRawTxInput(rawtx, txid string, index int) (string, error) {
	return c.WhcCreateRawTxInputAsync(rawtx, txid, index).Receive()
}

type FutureWhcCreateRawTxOpReturnResult chan *response

func (r FutureWhcCreateRawTxOpReturnResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreateRawTxOpReturnAsync(rawtx, payload string) FutureWhcCreateRawTxOpReturnResult {
	cmd := btcjson.NewWhcCreateRawTxOpReturnCmd(rawtx, payload)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreateRawTxOpReturn(rawtx, payload string) (string, error) {
	return c.WhcCreateRawTxOpReturnAsync(rawtx, payload).Receive()
}

type FutureWhcCreateRawTxReferenceResult chan *response

func (r FutureWhcCreateRawTxReferenceResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreateRawTxReferenceAsync(rawtx, address string,
	amount *string) FutureWhcCreateRawTxReferenceResult {

	cmd := btcjson.NewWhcCreateRawTxReferenceCmd(rawtx, address, amount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreateRawTxReference(rawtx, address string, amount *string) (string, error) {
	return c.WhcCreateRawTxReferenceAsync(rawtx, address, amount).Receive()
}

type FutureWhcDecodeTransactionResult chan *response

func (r FutureWhcDecodeTransactionResult) Receive() (*btcjson.GenerateTransactionResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.GenerateTransactionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcDecodeTransactionAsync(rawtx string, prev *[]btcjson.PrevTx,
	height *int) FutureWhcDecodeTransactionResult {

	cmd := btcjson.NewWhcDecodeTransactionCmd(rawtx, prev, height)
	return c.sendCmd(cmd)
}

func (c *Client) WhcDecodeTransaction(rawtx string, prev *[]btcjson.PrevTx,
	height *int) (*btcjson.GenerateTransactionResult, error) {

	return c.WhcDecodeTransactionAsync(rawtx, prev, height).Receive()
}

type FutureWhcBurnBCHGetWhcResult chan *response

func (r FutureWhcBurnBCHGetWhcResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcBurnBCHGetWhcAsync(amount string, redeemAddress *string) FutureWhcBurnBCHGetWhcResult {
	cmd := btcjson.NewWhcBurnBCHGetWhcCmd(amount, redeemAddress)
	return c.sendCmd(cmd)
}

func (c *Client) WhcBurnBCHGetWhc(amount string, redeemAddress *string) (string, error) {
	return c.WhcBurnBCHGetWhcAsync(amount, redeemAddress).Receive()
}

type FutureWhcPartiCrowdSaleResult chan *response

func (r FutureWhcPartiCrowdSaleResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcPartiCrowdSaleAsync(from, to, amount string, redeemAddress,
referenceAmount *string) FutureWhcPartiCrowdSaleResult {

	cmd := btcjson.NewWhcPartiCrowdSaleCmd(from, to, amount, redeemAddress, referenceAmount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcPartiCrowdSale(from, to, amount string, redeemAddress,
referenceAmount *string) (string, error) {

	return c.WhcPartiCrowdSaleAsync(from, to, amount, redeemAddress, referenceAmount).Receive()
}

type FutureWhcSendResult chan *response

func (r FutureWhcSendResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendAsync(from, to string, id uint64, amount string, redeem,
referenceAmount *string) FutureWhcSendResult {

	cmd := btcjson.NewWhcSendCmd(from, to, id, amount, redeem, referenceAmount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSend(from, to string, id uint64, amount string, redeem,
referenceAmount *string) (string, error) {

	return c.WhcSendAsync(from, to, id, amount, redeem, referenceAmount).Receive()
}

type FutureWhcSendAllResult chan *response

func (r FutureWhcSendAllResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendAllAsync(from, to string, eco int64, redeem,
referenceAmount *string) FutureWhcSendAllResult {

	cmd := btcjson.NewWhcSendAllCmd(from, to, eco, redeem, referenceAmount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendAll(from, to string, eco int64, redeem,
referenceAmount *string) (string, error) {

	return c.WhcSendAllAsync(from, to, eco, redeem, referenceAmount).Receive()
}

type FutureWhcSendChangeIssuerResult chan *response

func (r FutureWhcSendChangeIssuerResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendChangeIssuerAsync(from, to string, id uint64) FutureWhcSendChangeIssuerResult {
	cmd := btcjson.NewWhcSendChangeIssuerCmd(from, to, id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendChangeIssuer(from, to string, id uint64) (string, error) {
	return c.WhcSendChangeIssuerAsync(from, to, id).Receive()
}

type FutureWhcSendCloseCrowdSaleResult chan *response

func (r FutureWhcSendCloseCrowdSaleResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendCloseCrowdSaleAsync(from string, id uint64) FutureWhcSendCloseCrowdSaleResult {
	cmd := btcjson.NewWhcSendCloseCrowdSaleCmd(from, id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendCloseCrowdSale(from string, id uint64) (string, error) {
	return c.WhcSendCloseCrowdSaleAsync(from, id).Receive()
}

type FutureWhcSendGrantResult chan *response

func (r FutureWhcSendGrantResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendGrantAsync(from, to string, id uint64,
	amount string, note *string) FutureWhcSendGrantResult {

	cmd := btcjson.NewWhcSendGrantCmd(from, to, id, amount, note)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendGrant(from, to string, id uint64, amount string,
	note *string) (string, error) {

	return c.WhcSendGrantAsync(from, to, id, amount, note).Receive()
}

type FutureWhcSendIssuanceCrowdSaleResult chan *response

func (r FutureWhcSendIssuanceCrowdSaleResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendIssuanceCrowdSaleAsync(from string, eco, precision, previousid int64, category,
subCategory, name, url, data string, desiredID int64, tokensPerUnit string, deadline,
earlyBonus, issuerPercentage int64, amount string) FutureWhcSendIssuanceCrowdSaleResult {

	cmd := btcjson.NewWhcSendIssuanceCrowdSaleCmd(from, eco, precision, previousid, category, subCategory,
		name, url, data, deadline, tokensPerUnit, deadline, earlyBonus, issuerPercentage, amount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendIssuanceCrowdSale(from string, eco, precision, previousid int64, category,
subCategory, name, url, data string, desiredID int64, tokensPerUnit string, deadline,
earlyBonus, issuerPercentage int64, amount string) (string, error) {

	return c.WhcSendIssuanceCrowdSaleAsync(from, eco, precision, previousid, category, subCategory,
		name, url, data, deadline, tokensPerUnit, deadline, earlyBonus, issuerPercentage, amount).Receive()
}

type FutureWhcSendIssuanceFixedResult chan *response

func (r FutureWhcSendIssuanceFixedResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendIssuanceFixedAsync(from string, eco, precision, previousid int64, category,
subCategory, name, url, data, total string) FutureWhcSendIssuanceFixedResult {

	cmd := btcjson.NewWhcSendIssuanceFixedCmd(from, eco, precision, previousid, category, subCategory,
		name, url, data, total)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendIssuanceFixed(from string, eco, precision, previousid int64, category,
subCategory, name, url, data, total string) (string, error) {

	return c.WhcSendIssuanceFixedAsync(from, eco, precision, previousid, category, subCategory,
		name, url, data, total).Receive()
}

type FutureWhcSendIssuanceManagedResult chan *response

func (r FutureWhcSendIssuanceManagedResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendIssuanceManagedAsync(from string, eco, precision, previousid int64, category,
subCategory, name, url, data string) FutureWhcSendIssuanceManagedResult {

	cmd := btcjson.NewWhcSendIssuanceManagedCmd(from, eco, precision, previousid, category, subCategory,
		name, url, data)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendIssuanceManaged(from string, eco, precision, previousid int64, category,
subCategory, name, url, data string) (string, error) {

	return c.WhcSendIssuanceManagedAsync(from, eco, precision, previousid, category, subCategory,
		name, url, data).Receive()
}

type FutureWhcSendRawTxResult chan *response

func (r FutureWhcSendRawTxResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendRawTxAsync(from, rawtx string, referenceAddress, redeem,
referenceAmount *string) FutureWhcSendRawTxResult {

	cmd := btcjson.NewWhcSendRawTxCmd(from, rawtx, referenceAddress, redeem, referenceAmount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendRawTx(from, rawtx string, referenceAddress, redeem,
referenceAmount *string) (string, error) {

	return c.WhcSendRawTxAsync(from, rawtx, referenceAddress, redeem, referenceAmount).Receive()
}

type FutureWhcSendRevokeResult chan *response

func (r FutureWhcSendRevokeResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendRevokeAsync(from string, id int64, amount string,
	note *string) FutureWhcSendRevokeResult {

	cmd := btcjson.NewWhcSendRevokeCmd(from, id, amount, note)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendRevoke(from string, id int64, amount string, note *string) (string, error) {
	return c.WhcSendRevokeAsync(from, id, amount, note).Receive()
}

type FutureWhcSendStoResult chan *response

func (r FutureWhcSendStoResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

func (c *Client) WhcSendStoAsync(from string, id int64, amount string, redeem *string,
	distributionProperty *int64) FutureWhcSendStoResult {

	cmd := btcjson.NewWhcSendStoCmd(from, id, amount, redeem, distributionProperty)
	return c.sendCmd(cmd)
}

func (c *Client) WhcSendSto(from string, id int64, amount string, redeem *string,
	distributionProperty *int64) (string, error) {

	return c.WhcSendStoAsync(from, id, amount, redeem, distributionProperty).Receive()
}

func (c *Client) WhcSendFreeze(fromAddress string, id int64, amount string, frozenAddress string) (string, error) {
	return c.WhcSendFreezeAsync(fromAddress, id, amount, frozenAddress).Receive()
}

func (c *Client) WhcSendFreezeAsync(fromAddress string, id int64, amount string, frozenAddress string) FutureWhcSendFreezeResult {
	cmd := btcjson.NewWhcSendFreezeCmd(fromAddress, id, amount, frozenAddress)
	return c.sendCmd(cmd)
}

type FutureWhcSendFreezeResult chan *response

func (r FutureWhcSendFreezeResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var txid string
	err = json.Unmarshal(res, &txid)
	if err != nil {
		return "", err
	}

	return txid, nil
}

type FutureWhcCreatePayloadFreezeResult chan *response

func (r FutureWhcCreatePayloadFreezeResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadFreezeAsync(toAddress string, id int64, amount string) FutureWhcCreatePayloadFreezeResult {
	cmd := btcjson.NewWhcCreatePayloadFreezeCmd(toAddress, id, amount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadFreeze(toAddress string, id int64, amount string) (string, error) {
	return c.WhcCreatePayloadFreezeAsync(toAddress, id, amount).Receive()
}

type FutureWhcCreatePayloadUnFreezeResult chan *response

func (r FutureWhcCreatePayloadUnFreezeResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	var result string
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) WhcCreatePayloadUnFreezeAsync(toAddress string, id int64, amount string) FutureWhcCreatePayloadUnFreezeResult {
	cmd := btcjson.NewWhcCreatePayloadUnFreezeCmd(toAddress, id, amount)
	return c.sendCmd(cmd)
}

func (c *Client) WhcCreatePayloadUnFreeze(toAddress string, id int64, amount string) (string, error) {
	return c.WhcCreatePayloadUnFreezeAsync(toAddress, id, amount).Receive()
}

type FutureWhcGetFrozenBalanceResult chan *response

func (r FutureWhcGetFrozenBalanceResult) Receive() (*btcjson.FrozenBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result btcjson.FrozenBalanceResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetFrozenBalanceAsync(address string, id int64) FutureWhcGetFrozenBalanceResult {
	cmd := btcjson.NewWhcGetFrozenBalanceCmd(address, id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetFrozenBalance(address string, id int64) (*btcjson.FrozenBalanceResult, error) {
	return c.WhcGetFrozenBalanceAsync(address, id).Receive()
}

type FutureWhcGetFrozenBalanceForIdResult chan *response

func (r FutureWhcGetFrozenBalanceForIdResult) Receive() (*[]btcjson.FrozenBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []btcjson.FrozenBalanceResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetFrozenBalanceForIdAsync(id int64) FutureWhcGetFrozenBalanceForIdResult {
	cmd := btcjson.NewWhcGetFrozenBalanceForIdCmd(id)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetFrozenBalanceForId(id int64) (*[]btcjson.FrozenBalanceResult, error) {
	return c.WhcGetFrozenBalanceForIdAsync(id).Receive()
}

type FutureWhcGetFrozenBalanceForAddressResult chan *response

func (r FutureWhcGetFrozenBalanceForAddressResult) Receive() (*[]btcjson.FrozenBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var result []btcjson.FrozenBalanceResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) WhcGetFrozenBalanceForAddressAsync(address string) FutureWhcGetFrozenBalanceForAddressResult {
	cmd := btcjson.NewWhcGetFrozenBalanceForAddressCmd(address)
	return c.sendCmd(cmd)
}

func (c *Client) WhcGetFrozenBalanceForAddress(address string) (*[]btcjson.FrozenBalanceResult, error) {
	return c.WhcGetFrozenBalanceForAddressAsync(address).Receive()
}

func (c *Client) WhcSendUnFreeze(fromAddress string, id int64, amount string, frozenAddress string) (string, error) {
	return c.WhcSendUnFreezeAsync(fromAddress, id, amount, frozenAddress).Receive()
}

func (c *Client) WhcSendUnFreezeAsync(fromAddress string, id int64, amount string, frozenAddress string) FutureWhcSendFreezeResult {
	cmd := btcjson.NewWhcSendUnFreezeCmd(fromAddress, id, amount, frozenAddress)
	return c.sendCmd(cmd)
}
