package aibridged

import "github.com/coder/coder/v2/x/aibridged/proto"

type DRPCServer interface {
	proto.DRPCRecorderServer
	proto.DRPCMCPConfiguratorServer
	proto.DRPCAuthorizerServer
}
