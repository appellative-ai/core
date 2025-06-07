package host

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"strings"
	"sync"
)

const (
	nameKey    = "name"
	pathKey    = "@path"
	configRole = "config"
)

type Resource struct {
	Role   string
	Name   string
	Config map[string]string
}

func DefineNetwork(path string, roles []string) (*MapT[string, Resource], []error) {
	if path == "" || len(roles) == 0 {
		return nil, []error{errors.New("error: network config path is empty or roles are empty")}
	}
	cfg, err := readConfig(path, configRole)
	if err != nil {
		return nil, []error{err}
	}
	result := make([]error, len(roles))
	var wg sync.WaitGroup
	var net = NewSyncMap[string, Resource]()
	var i int
	for _, role := range roles {
		value, ok := cfg[role]
		if !ok || value == "" {
			continue
		}
		if i != 0 {
			i++
		}
		wg.Add(1)
		go func(role1 string, err *error) {
			defer wg.Done()
			rsc := Resource{Role: role1}

			m := parseConfig(value)
			rsc.Name = m[nameKey]
			if m[pathKey] != "" {
				rsc.Config, *err = readConfig(newPath(path, m[pathKey]), role1)
				if *err != nil {
					return
				}
			}
			net.Store(rsc.Role, rsc)
		}(role, &result[i])
	}
	wg.Wait()
	return net, packErrors(result)
}

func packErrors(errs []error) []error {
	var result []error
	for _, err := range errs {
		if err != nil {
			result = append(result, err)
		}
	}
	return result
}

func newPath(path, fileName string) string {
	i := strings.LastIndex(path, "/")
	if i == -1 {
		return "error: invalid path"
	}
	//i2 := strings.Index(fileName, "=")
	//if i2 == -1 {
	//	return "error: invalid file name"
	//}
	return path[:i+1] + fileName
}

func readConfig(path string, role string) (map[string]string, error) {
	buf, err := iox.ReadFile(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v: for %v role", err, role))
	}
	var cfg map[string]string
	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%v: for %v role", err, role))
	}
	return cfg, nil
}

func parseConfig(s string) map[string]string {
	var m = make(map[string]string)

	tokens := strings.Split(s, ",")
	for _, t := range tokens {
		pairs := strings.Split(t, "=")
		if len(pairs) < 2 || pairs[1] == "" {
			continue
		}
		m[pairs[0]] = pairs[1]
	}
	return m
}
