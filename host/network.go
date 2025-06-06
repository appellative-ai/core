package host

import (
	"encoding/json"
	"errors"
	"github.com/behavioral-ai/core/iox"
	"strings"
	"sync"
)

const (
	AtPath = "@path"
)

type Resource struct {
	Role   string
	Name   string
	Err    error
	Config map[string]string
}

func DefineNetwork(path string, roles []string) (*MapT[string, Resource], error) {
	if path == "" || len(roles) == 0 {
		return nil, errors.New("error: network config path is empty or roles are empty")
	}
	cfg, err := readConfig(path)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var net = NewSyncMap[string, Resource]()
	for _, role := range roles {
		if c, ok := cfg[role]; ok {
			wg.Add(1)
			go func(role1 string) {
				defer wg.Done()
				var rsc Resource
				var m map[string]string
				rsc.Role = role1
				m, rsc.Err = parseConfig(c)
				rsc.Name = m["name"]
				if m[AtPath] != "" {
					rsc.Config, rsc.Err = readConfig(newPath(path, m[AtPath]))
				}
				//fmt.Printf("Resource -> %v\n", rsc)
				net.Store(rsc.Role, rsc)
			}(role)
		}
	}

	/*
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


	*/
	wg.Wait()
	return net, nil
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
