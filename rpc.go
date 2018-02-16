package omnigo

import (
    "bytes"
    "github.com/gorilla/rpc/json"
    "net/http"
)

func makeJsonRpcRequest(url, responseTarget interface{}, args interface{}) err {
    msg, err := json.EncodeClient("omni", args)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    client := new(http.Client)
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    defer resp.Body.Close()
    err = json.DecodeClientResponse(resp.Body, responseTarget)
    if err != nil {
        return err
    }
}

func OmniSend(fromAddress, toAddress string, propertyId int, amount, redeemAddress, referenceAmount string) string {
    msg, err := json.EncodeClient("omni", [fromAddress, toAddress, propertyId, amount, redeemAddress, referenceAmount])
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")
    client := new(http.Client)
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    defer resp.Body.Close()
    var transactionHash string
    err = json.DecodeClientResponse(resp.Body, &transactionHash)
    if err != nil {
        return err
    }

    return nil
}

func OmniSendDexSell(fromAddress string, propertyIdForSale int, amountForSale, amountDesired string, paymentWindow int, minAcceptFee string, action int) string {
    msg, err := json.EncodeClientRequest("omni_senddexsell", [fromAddress, propertyIdForSale, amountForSale, amountDesired, paymentWindow, minAcceptFee, action])
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")
    client := new(http.Client)
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }

    defer resp.Body.Close()

    var transactionHash string
    err = json.DecodeClientResponse(resp.Body, &transactionHash)
    if err != nil {
        return "", err
    }

    return transactionHash, nil
}

func OmniSendDexAccept(fromAddress, toAddress string, propertyId int, amount string, override bool) string {
    var transactionHash string
    err := makeJsonRpcRequest(url, &transactionHash, fromAddress, toAddress, propertyId, amount, override)
    if err != nil {
        return "", err
    }

    return transactionHash, nil
}

func OmniSendIssuanceCrowdSale(fromAddress string, ecosystem, type, previousId number, category, subCategory, name, url, data string, propertyIdDesired int, tokensPerUnit string, deadline, earlyBonus, issuerPercentage int) string {
    var transactionHash string
    err := makeJsonRpcRequest(url, &transactionHash, fromAddress, ecosystem, type, previousId, category, subCategory, name, url, data, propertyIdDesired, tokensPerUnit, deadline, earlyBonus, issuerPercentage)
    if err != nil {
        return "", err
    }

    return transactionHash, nil
}

func OmniSendIssuanceFixed(fromAddress string, ecosystem, type, previousId int, category, subCategory, name, url, data, amount, tokensPerUnit string, deadline, earlyBonus, issuerPercentage int) string {
    var transactionHash string
    err := makeJsonRpcRequest(url, &transactionHash, fromAddress, ecosystem, type, previousId, category, subCategory, name, url, data, amount, tokensPerUnit, deadline, earlyBonus, issuerPercentage)
    if err != nil {
        return "", err
    }

    return transactionHash, nil
}


func OmniGetAllBalancesForAddress(address string) *[]Balance, error {
    msg, err := json.EncodeClientRequest("omni_getallbalancesforaddress", address)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(msg))
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")
    client := new(http.Client)
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    balances := []Balance{}
    err := json.DecodeClientResponse(resp.Body, &balances)
    if err != nil {
        return balances, err
    }

    return balances, nil
}