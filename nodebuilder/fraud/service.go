package fraud

import (
	"context"

	"github.com/celestiaorg/celestia-node/fraud"
)

// Module encompasses the behavior necessary to subscribe and broadcast fraud proofs within the network.
// Any method signature changed here needs to also be changed in the API struct.
type Module interface {
	fraud.Service
}

// API is a wrapper around Module for the RPC.
// TODO(@distractedm1nd): These structs need to be autogenerated.
//
//go:generate go run github.com/golang/mock/mockgen -destination=mocks/api.go -package=mocks . Module
type API struct {
	Subscribe func(fraud.ProofType) (fraud.Subscription, error)
	Get       func(context.Context, fraud.ProofType) ([]fraud.Proof, error)
}