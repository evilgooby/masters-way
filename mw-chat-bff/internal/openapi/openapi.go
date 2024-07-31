package openapi

import (
	openapiChat "mw-chat-bff/apiAutogenerated/chat"
	openapiGeneral "mw-chat-bff/apiAutogenerated/general"
	"mw-chat-bff/internal/auth"
	"mw-chat-bff/internal/config"
	"net/http"
)

type authTransport struct {
	rt http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	if token, ok := ctx.Value(auth.ContextKeyAuthorization).(string); ok {
		req.Header.Set(auth.HeaderKeyAuthorization, token)
	}
	return t.rt.RoundTrip(req)
}

func MakeGeneralAPIConfig(cfg *config.Config) *openapiGeneral.Configuration {
	return &openapiGeneral.Configuration{
		Host:   cfg.GeneralAPIHost,
		Scheme: "http",
		Servers: openapiGeneral.ServerConfigurations{
			{
				URL:         cfg.GeneralBaseURL,
				Description: "mw-general",
			},
		},
		HTTPClient: &http.Client{
			Transport: &authTransport{rt: http.DefaultTransport},
		},
	}
}

func MakeGeneralAPIClient(cfg *config.Config) *openapiGeneral.APIClient {
	generalAPIConfig := MakeGeneralAPIConfig(cfg)
	return openapiGeneral.NewAPIClient(generalAPIConfig)
}

func MakeChatAPIConfig(cfg *config.Config) *openapiChat.Configuration {
	return &openapiChat.Configuration{
		Host:   cfg.ChatAPIHost,
		Scheme: "http",
		Servers: openapiChat.ServerConfigurations{
			{
				URL:         cfg.ChatBaseURL,
				Description: "mw-chat",
			},
		},
		HTTPClient: &http.Client{
			Transport: &authTransport{rt: http.DefaultTransport},
		},
	}
}

func MakeChatAPIClient(cfg *config.Config) *openapiChat.APIClient {
	chatAPIConfig := MakeChatAPIConfig(cfg)
	return openapiChat.NewAPIClient(chatAPIConfig)
}
