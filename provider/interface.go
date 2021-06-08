package provider

import (
	"io/fs"
	"net/http"
	"sync"
)

type FS struct {
	mp   *Map
	IoFS fs.FS
	Root string
}

type HTTP struct {
	mp     *Map
	Client *http.Client
}

type Map struct {
	lock    sync.Mutex
	mapping map[string]interface{}
}
