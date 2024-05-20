package transportClient

import (
	"minerrpc/rpc_common/entities"
	"minerrpc/rpc_core/serializer"
)

const (
	DEFAULT_SERIALIZER = serializer.JSONSerializer
)

type RpcClient interface {
	SendRequest(req entities.RPCdata) (*entities.RPCdata, error)
}
