// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type CreateBankAccountResponse struct {
	// OK
	BankAccountResponse *shared.BankAccountResponse
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}
