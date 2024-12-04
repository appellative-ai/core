package uri

import (
	"errors"
	"fmt"
	"net/url"
)

const (
	DomainPath     = "domain"
	DomainRootPath = "/" + DomainPath
)

// ValidateURL - validate a URL
func ValidateURL(url *url.URL, domain string) (p *Parsed, err error) {
	if url == nil {
		return &Parsed{}, errors.New("error: URL is nil")
	}
	if len(domain) == 0 {
		return &Parsed{}, errors.New("error: domain is empty")
	}
	if url.Path == DomainRootPath {
		return &Parsed{Path: DomainPath}, nil
	}
	if url.RawQuery != "" {
		p = Uproot(url.Path + "?" + url.RawQuery)
	} else {
		p = Uproot(url.Path)
	}
	if !p.Valid {
		return p, p.Err
	}
	if p.Domain != domain {
		return p, errors.New(fmt.Sprintf("error: invalid URI, domain does not match: \"%v\" \"%v\"", url.Path, domain))
	}
	if len(p.Path) == 0 {
		return p, errors.New(fmt.Sprintf("error: invalid URI, path only contains a domain: \"%v\"", url.Path))
	}
	return p, nil
}
