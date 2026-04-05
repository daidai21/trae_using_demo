package identity

import (
	"context"
	"net/http"
)

const (
	HeaderBusinessIdentity = "X-Business-Identity"
	QueryBusinessIdentity  = "identity"
	CookieBusinessIdentity = "business_identity"
)

type IdentityResolver interface {
	Resolve(ctx context.Context, req *http.Request) (*BusinessIdentity, error)
	ResolveFromUser(userID uint) (*BusinessIdentity, error)
}

type DefaultIdentityResolver struct {
	defaultIdentity *BusinessIdentity
}

func NewDefaultIdentityResolver() *DefaultIdentityResolver {
	return &DefaultIdentityResolver{
		defaultIdentity: NewBusinessIdentity(CountryCN, ModeNormal),
	}
}

func (r *DefaultIdentityResolver) Resolve(ctx context.Context, req *http.Request) (*BusinessIdentity, error) {
	if identityStr := req.URL.Query().Get(QueryBusinessIdentity); identityStr != "" {
		if id, err := Parse(identityStr); err == nil {
			return id, nil
		}
	}

	if identityStr := req.Header.Get(HeaderBusinessIdentity); identityStr != "" {
		if id, err := Parse(identityStr); err == nil {
			return id, nil
		}
	}

	if cookie, err := req.Cookie(CookieBusinessIdentity); err == nil && cookie.Value != "" {
		if id, err := Parse(cookie.Value); err == nil {
			return id, nil
		}
	}

	return r.defaultIdentity, nil
}

func (r *DefaultIdentityResolver) ResolveFromUser(userID uint) (*BusinessIdentity, error) {
	return r.defaultIdentity, nil
}
