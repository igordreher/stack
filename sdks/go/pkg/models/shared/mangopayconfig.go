// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MangoPayConfig struct {
	APIKey   string `json:"apiKey"`
	ClientID string `json:"clientID"`
	Endpoint string `json:"endpoint"`
	Name     string `json:"name"`
	// The frequency at which the connector will try to fetch new BalanceTransaction objects from MangoPay API.
	//
	PollingPeriod *string `json:"pollingPeriod,omitempty"`
}
