package transportClient

import (
	"jrpc/src/rpc_common/entities"
	"jrpc/src/rpc_core/serializer"
)

const (
	DEFAULT_SERIALIZER = serializer.JSONSerializer
)

type RpcClient interface {
	SendRequest(req entities.RPCdata) (*entities.RPCdata, error)
}