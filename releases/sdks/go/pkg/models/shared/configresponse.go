// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type ConfigResponse struct {
	Data WebhooksConfig `json:"data"`
}

func (o *ConfigResponse) GetData() WebhooksConfig {
	if o == nil {
		return WebhooksConfig{}
	}
	return o.Data
}
