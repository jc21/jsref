package provider

import (
	"net/url"

	"github.com/lestrrat-go/pdebug/v3"
	"github.com/pkg/errors"
)

// NewMap ...
func NewMap() *Map {
	return &Map{
		mapping: make(map[string]any),
	}
}

// Set ...
func (mp *Map) Set(key string, v any) error {
	mp.lock.Lock()
	defer mp.lock.Unlock()

	mp.mapping[key] = v
	return nil
}

// Get ...
func (mp *Map) Get(key *url.URL) (res any, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("Map.Get(%s)", key).BindError(&err)
		defer g.End()
	}

	mp.lock.Lock()
	defer mp.lock.Unlock()

	v, ok := mp.mapping[key.String()]
	if !ok {
		return nil, errors.New("not found")
	}

	return v, nil
}

// Reset ...
func (mp *Map) Reset() error {
	mp.lock.Lock()
	defer mp.lock.Unlock()

	mp.mapping = make(map[string]any)
	return nil
}
