package keyring

import (
	"github.com/hashicorp/vault/helper/keysutil"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func Factory(conf *logical.BackendConfig) (logical.Backend, error) {
	b := Backend(conf)
	be, err := b.Backend.Setup(conf)
	if err != nil {
		return nil, err
	}

	return be, nil
}

func Backend(conf *logical.BackendConfig) *backend {
	var b backend
	b.Backend = &framework.Backend{
		Paths: []*framework.Path{
			b.pathKeys(),
			b.pathListKeys(),
			b.pathRotate(),
			b.pathValueVersion(),
			b.pathValue(),
		},
	}

	return &b
}

type backend struct {
	*framework.Backend
	lm *keysutil.LockManager
}
