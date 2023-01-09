package quickbooks

type ENV string

const (
	SANDBOX    ENV = "https://developer.api.intuit.com/.well-known/openid_sandbox_configuration"
	PRODUCTION     = "https://developer.api.intuit.com/.well-known/openid_configuration"
)
