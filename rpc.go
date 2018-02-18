package omnigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type OmniRpcClient struct {
	Url      string
	Username string
	Password string
}

type clientRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	Id     uint64        `json:"id"`
}

type clientResponse struct {
	Result *json.RawMessage `json:"result"`
	Error  interface{}      `json:"error"`
	Id     uint64           `json:"id"`
}

func encodeClientRequest(method string, args ...interface{}) ([]byte, error) {
	var params []interface{}

	if len(args) == 0 {
		params = make([]interface{}, 0)
	} else {
		params = args
	}

	c := &clientRequest{
		Method: method,
		Params: params,
		Id:     uint64(rand.Int63()),
	}

	return json.Marshal(c)
}

func decodeClientResponse(r io.Reader, reply interface{}) error {
	var c clientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return err
	}

	if c.Error != nil {
		return fmt.Errorf("%v", c.Error)
	}

	if c.Result == nil {
		return errors.New("result is null")
	}

	return json.Unmarshal(*c.Result, reply)
}

func (omni *OmniRpcClient) invokeJsonRpcMethod(decodeTarget interface{}, method string, args ...interface{}) error {
	msg, err := encodeClientRequest("omni_getinfo")

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", omni.Url, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(omni.Username, omni.Password)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = decodeClientResponse(resp.Body, &decodeTarget)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (omni *OmniRpcClient) OmniSendDexSell(fromAddress string, propertyIdForSale int, amountForSale, amountDesired string, paymentWindow int, minAcceptFee string, action int) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_senddexsell", fromAddress, propertyIdForSale, amountForSale, amountDesired, paymentWindow, minAcceptFee, action)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendDexAccept(fromAddress, toAddress string, propertyId int, amount string, override bool) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_senddexaccept", fromAddress, toAddress, propertyId, amount, override)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendIssuanceCrowdSale(fromAddress string, ecosystem, tokenType, previousId int, category, subCategory, name, url, data string, propertyIdDesired int, tokensPerUnit string, deadline, earlyBonus, issuerPercentage int) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendissuancecrowdsale", fromAddress, ecosystem, tokenType, previousId, category, subCategory, name, url, data, propertyIdDesired, tokensPerUnit, deadline, earlyBonus, issuerPercentage)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendIssuanceFixed(fromAddress string, ecosystem, tokenType, previousId int, category, subCategory, name, url, data, amount, tokensPerUnit string, deadline, earlyBonus, issuerPercentage int) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendissuancefixed", fromAddress, ecosystem, tokenType, previousId, category, subCategory, name, url, data, amount, tokensPerUnit, deadline, earlyBonus, issuerPercentage)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendIssuanceManaged(fromAddress string, ecosystem, tokenType, previousId int, category, subCategory, name, url, data string) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendissuancemanaged", fromAddress, ecosystem, previousId, category, subCategory, name, url, data)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendsTo(fromAddress string, propertyId int, amount, redeemAddress string, distributionProperty int) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendsto", fromAddress, propertyId, amount, redeemAddress, distributionProperty)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendGrant(fromAddress, toAddress string, propertyId int, amount, memo string) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendgrant", fromAddress, toAddress, propertyId, amount, memo)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendRevoke(fromAddress string, propertyId int, amount, memo string) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendrevoke", fromAddress, propertyId, amount, memo)
	return txHash, err
}

func (omni *OmniRpcClient) OmniSendCloseCrowdSale(fromAddress string, propertyId int) (string, error) {
	var txHash string
	err := omni.invokeJsonRpcMethod(&txHash, "omni_sendclosecrowdsale", fromAddress, propertyId)
	return txHash, err
}

func (omni *OmniRpcClient) OmniGetAllBalancesForAddress(address string) (*[]Balance, error) {
	balances := []Balance{}
	err := omni.invokeJsonRpcMethod(&balances, "omni_getallbalancesforaddress", address)

	return &balances, err
}

func (omni *OmniRpcClient) GetInfo() (*OmniInfo, error) {
	info := OmniInfo{}
	err := omni.invokeJsonRpcMethod(&info, "omni_getinfo")

	return &info, err
}
