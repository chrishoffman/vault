package keyring

import (
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func (b *backend) pathValue() *framework.Path {
	return &framework.Path{
		Pattern: "value/" + framework.GenericNameRegex("name"),

		Callbacks: map[logical.Operation]framework.OperationFunc{},
	}
}

func (b *backend) pathValueVersion() *framework.Path {
	return &framework.Path{
		Pattern: "value/" + framework.GenericNameRegex("name") + "/" +
			framework.GenericNameRegex("version"),

		Fields: map[string]*framework.FieldSchema{
			"name": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "The backend key name",
			},

			"version": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "the version of key to retrieve",
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{},
	}
}
