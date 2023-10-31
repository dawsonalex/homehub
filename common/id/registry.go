package id

import (
	"errors"
	"sync"
)

var registry = makeRegistry()

var ErrPrefixExists = errors.New("id prefix already exists")

type Registry struct {
	typeMap map[string]string // a map of ID prefixes to keys
	sync.Mutex
}

func (r *Registry) get(prefix string) (string, bool) {
	r.Lock()
	defer r.Unlock()

	prefix = padPrefix(prefix)
	if key, keyExists := registry.typeMap[prefix]; keyExists {
		return key, true
	}
	return "", false
}

func makeRegistry() *Registry {
	return &Registry{
		typeMap: make(map[string]string),
	}
}

func RegisterType(typePrefix, typeKey string) error {
	registry.Lock()
	defer registry.Unlock()

	typePrefix = padPrefix(typePrefix)
	if _, idExists := registry.typeMap[typePrefix]; idExists {
		return ErrPrefixExists
	}

	registry.typeMap[typePrefix] = typeKey
	return nil
}
