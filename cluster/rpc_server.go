package cluster

import (
	"context"
	"github.com/topfreegames/pitaya/protos"
)

// PitayaServer is the server API for Pitaya service.
type PitayaServer interface {
	Call(context.Context, *protos.Request) (*protos.Response, error)
	PushToUser(context.Context, *protos.Push) (*protos.Response, error)
	SessionBindRemote(context.Context, *protos.BindMsg) (*protos.Response, error)
	KickUser(context.Context, *protos.KickMsg) (*protos.KickAnswer, error)
}
