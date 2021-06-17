package token

import "context"

type TokenAuth struct {
	AppKey    string
	AppSecret string
}

func (c *TokenAuth) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  c.AppKey,
		"appkey": c.AppSecret,
	}, nil
}

//是否基于TLS认证进行安全传输
func (c *TokenAuth) RequireTransportSecurity() bool {
	return false
}
