package remote

import (
	"fmt"
)

// Client is the interface that must be implemented for a remote state
// driver. It supports dumb put/get/delete, and the higher level structs
// handle persisting the state properly here.
type Client interface {
	Get() (*Payload, error)
	Put([]byte) error
	Delete() error
}

// Payload is the return value from the remote state storage.
type Payload struct {
	MD5  []byte
	Data []byte
}

// Factory is the factory function to create a remote client.
type Factory func(map[string]string) (Client, error)

// NewClient returns a new Client with the given type and configuration.
// The client is looked up in the BuiltinClients variable.
func NewClient(t string, conf map[string]string) (Client, error) {
	f, ok := BuiltinClients[t]
	if !ok {
		return nil, fmt.Errorf("unknown remote client type: %s", t)
	}

	return f(conf)
}

// BuiltinClients is the list of built-in clients that can be used with
// NewClient.
var BuiltinClients = map[string]Factory{
	"atlas":       atlasFactory,
	"azure":       azureFactory,
	"consul":      consulFactory,
	"etcd":        etcdFactory,
	"gcs":         gcsFactory,
	"http":        httpFactory,
	"s3":          s3Factory,
	"swift":       swiftFactory,
	"artifactory": artifactoryFactory,

	// This is used for development purposes only.
	"_local": fileFactory,
}
