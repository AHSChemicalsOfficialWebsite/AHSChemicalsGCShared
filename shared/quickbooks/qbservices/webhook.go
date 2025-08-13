package qbservices

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/quickbooks"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/quickbooks/qbmodels"
	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/utils"
	"github.com/go-jose/go-jose/v4/json"
)

func VerifyQuickBooksWebhookSignature(body []byte, signature string) bool {

	mac := hmac.New(sha256.New, []byte(quickbooks.QUICKBOOKS_WEBHOOK_VERIFY_TOKEN))
	mac.Write(body)

	expectedMAC := mac.Sum(nil)
	//Quickbooks sends the signature in base64 format, so need to encode the generated hmac as well
	expectedSignature := base64.StdEncoding.EncodeToString(expectedMAC) 

	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

func GetQBCustomerFromEntityID(id string) (*qbmodels.QBCustomer, error){
	reqURL := quickbooks.QUICKBOOKS_GET_CUSTOMER_URL
	params :=  url.Values{}
	params.Set("id", id)

	reqURL = reqURL + "?" + params.Encode()
	client := http.Client{Timeout: 10*time.Second}

	resp, err := client.Get(reqURL)
	if err != nil{
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		return nil, errors.New("Failed to get the customer from quickbooks" + string(respBody))
	}

	var respPaylod utils.SuccessPaylod[string]
	err = json.Unmarshal(respBody, &respPaylod)
	if err != nil{
		return nil, err
	}

	var quickbooksCustomer qbmodels.QBCustomersResponse
	err = json.Unmarshal([]byte(respPaylod.Data), &quickbooksCustomer)
	if err != nil{
		return nil, err
	}
	
	return &quickbooksCustomer.QueryResponse.Customer[0], nil
}

func GetQBProductFromEntityID(id string) (*qbmodels.QBItem, error){
	reqURL := quickbooks.QUICKBOOKS_GET_PRODUCT_URL
	params :=  url.Values{}
	params.Set("id", id)

	reqURL = reqURL + "?" + params.Encode()
	client := http.Client{Timeout: 10*time.Second}

	resp, err := client.Get(reqURL)
	if err != nil{
		return nil, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		return nil, errors.New("Failed to get item from quickbooks" + string(respBody))
	}
	
	var respPaylod utils.SuccessPaylod[string]
	err = json.Unmarshal(respBody, &respPaylod)
	if err != nil{
		return nil, errors.New("Could not get item from quickbooks" + string(respBody))
	}

	var quickbooksItem qbmodels.QBItemsResponse
	err = json.Unmarshal([]byte(respPaylod.Data), &quickbooksItem)
	if err != nil{
		return nil, err
	}

	return &quickbooksItem.QueryResponse.Item[0], nil
}