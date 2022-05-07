package yamlx

import (
	"errors"
	"fmt"
	"strings"
)

type Map map[string]string

func (t *Map) set(data any, path string) {
	switch u := data.(type) {
	case map[string]any:
		for k, v := range u {
			if len(path) == 0 {
				t.set(v, fmt.Sprint(k))
			} else {
				t.set(v, path+"."+fmt.Sprint(k))
			}
		}
	case map[any]any:
		for k, v := range u {
			if len(path) == 0 {
				t.set(v, fmt.Sprint(k))
			} else {
				t.set(v, path+"."+fmt.Sprint(k))
			}
		}
	default:
		if len(path) != 0 {
			(*t)[path] = fmt.Sprint(data)
		}
	}
}

func (t *Map) Set(data any) {
	t.set(data, "")
}

func marshal(data map[string]string, result *map[any]any) error {
	done := map[string]bool{}
	for k, v := range data {
		split := strings.SplitN(k, ".", 2)
		root := split[0]
		if done[root] {
			continue
		}
		if len(split) == 1 {
			(*result)[root] = v
			continue
		}
		n := map[string]string{}
		count := 0
		final := false
		for k2, v2 := range data {
			split := strings.SplitN(k2, ".", 2)
			if root == split[0] {
				n[strings.TrimPrefix(k2, root+".")] = v2
				count++
			}
			if len(split) == 1 {
				final = true
			}
		}
		if count == 1 && final {
			return errors.New("got value and root node")
		}
		done[root] = true
		s := &map[any]any{}
		(*result)[root] = s
		err := marshal(n, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Map) Marshal() (map[any]any, error) {
	r := &map[any]any{}
	err := marshal(*t, r)
	if err != nil {
		return nil, err
	}
	return *r, nil
}
