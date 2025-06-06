package host

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/iox"
	"strings"
)

const (
	AtPath = "@path"
)

type Resource struct {
	Name   string
	Config map[string]string
}

func DefineNetwork(path string) ([]Resource, error) {
	var ops []Resource
	if path == "" {
		return nil, errors.New("error: network config path is empty")
	}
	net, err := readConfig(path)
	if err != nil {
		return nil, err
	}
	for k, v := range net {
		if strings.HasPrefix(v, AtPath) {
			cfg, err2 := readConfig(newPath(path, v))
			if err2 != nil {
				return ops, err2
			}
			ops = append(ops, Resource{Name: k, Config: cfg})
		} else {
			m, err1 := parseConfig(v)
			if err1 != nil {
				return ops, err1
			}
			ops = append(ops, Resource{Name: k, Config: m})
		}
	}
	return ops, nil
}

func newPath(path, fileName string) string {
	i := strings.LastIndex(path, "/")
	if i == -1 {
		return "error: invalid path"
	}
	i2 := strings.Index(fileName, "=")
	if i2 == -1 {
		return "error: invalid file name"
	}
	return path[:i+1] + fileName[i2+1:]
}

func readConfig(path string) (map[string]string, error) {
	buf, err := iox.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var net map[string]string
	err = json.Unmarshal(buf, &net)
	return net, err
}

func parseConfig(s string) (map[string]string, error) {
	var m = make(map[string]string)

	tokens := strings.Split(s, ",")
	for _, t := range tokens {
		pairs := strings.Split(t, "=")
		if len(pairs) < 2 || pairs[1] == "" {
			continue
		}
		m[pairs[0]] = pairs[1]
	}
	return m, nil
}
