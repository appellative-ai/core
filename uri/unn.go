package uri

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	UnnScheme = "unn"
	UnnPrefix = UnnScheme + Colon
	Colon     = ":"
	Slash     = "/"
	Fragment  = "#"
)

type Unn struct {
	Domain    string
	Namespace string
	Class     string
	Path      string
	Resource  string
	Fragment  string
	Err       error
}

func ParseUnn(uri string) *Unn {
	u := new(Unn)
	newUri := uprootUnn(uri)
	tokens := strings.Split(newUri, Colon)
	if len(tokens) < 3 || tokens[0] == "" {
		u.Err = errors.New(fmt.Sprintf("invalid argument: invalid number of components [%v] [%v]", uri, len(tokens)))
		return u
	}
	for i := 0; i < len(tokens); i++ {
		switch i {
		case 0:
			u.Domain = tokens[i]
		case 1:
			u.Namespace = tokens[i]
		case 2:
			parseClass(tokens[i], u)
		case 3:
			parseResource(tokens[i], u)
		}
	}
	return u
}

func parseClass(s string, u *Unn) error {
	i := strings.Index(s, Slash)
	if i == -1 {
		return errors.New(fmt.Sprintf("invalid argument: no path for agent [%v]", s))
	}
	u.Class = s[:i]
	u.Path = s[i+1:]
	return nil
}

func parseResource(s string, u *Unn) {
	i := strings.Index(s, Fragment)
	if i == -1 {
		u.Resource = s

	} else {
		u.Resource = s[:i]
		u.Fragment = s[i+1:]
	}
}

func uprootUnn(uri string) string {
	if strings.HasPrefix(uri, UnnPrefix) {
		return uri[len(UnnPrefix):]
	}
	u, err := url.Parse(uri)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	path := ""
	if strings.HasPrefix(u.Path, Slash) {
		path = u.Path[1:]
	} else {
		path = u.Path
	}
	if u.Fragment != "" {
		return path + Fragment + u.Fragment
	}
	return path
}
