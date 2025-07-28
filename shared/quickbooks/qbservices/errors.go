package qbservices

import (
	"encoding/json"
	"errors"

	"github.com/HarshMohanSason/AHSChemicalsGCShared/shared/quickbooks/qbmodels"
)

func ReturnQBOAuthError(err error) error {
	var QBOAuthErrorResponse qbmodels.QBOAuthErrorResponse
	unMarshalErr := json.Unmarshal([]byte(err.Error()), &QBOAuthErrorResponse)
	if unMarshalErr != nil {
		return errors.New("An error occurred while authenticating with Quickbooks. Please try again later.")
	}
	return errors.New(QBOAuthErrorResponse.Error)
}

func ReturnQBFaultError(err error) error {
	var QBFaultErrorResponse qbmodels.QBFaultErrorResponse
	unMarshalErr := json.Unmarshal([]byte(err.Error()), &QBFaultErrorResponse)
	//Return original error message if an error occurs
	if unMarshalErr != nil {
		return errors.New("An error occurred while syncing with Quickbooks. Please try again later.")
	}
	return errors.New(QBFaultErrorResponse.Fault.ErrorList[0].Message)
}