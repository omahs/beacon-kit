// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package deposit

import (
	"context"
	"math/big"

	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

// BeaconBlock is an interface for beacon blocks.
type BeaconBlock[
	BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT],
	DepositT interface{ GetIndex() uint64 },
	ExecutionPayloadT interface{ GetNumber() math.U64 },
] interface {
	GetSlot() math.U64
	GetBody() BeaconBlockBodyT
}

type BeaconBlockBody[
	DepositT interface{ GetIndex() uint64 },
	ExecutionPayloadT interface{ GetNumber() math.U64 },
] interface {
	GetDeposits() []DepositT
	GetExecutionPayload() ExecutionPayloadT
}

// BlockEvent is an interface for block events.
type BlockEvent[
	BeaconBlockT BeaconBlock[BeaconBlockBodyT, DepositT, ExecutionPayloadT],
	BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT],
	DepositT interface{ GetIndex() uint64 },
	ExecutionPayloadT interface{ GetNumber() math.U64 },
] interface {
	Context() context.Context
	Block() BeaconBlockT
}

// BlockFeed is an interface for subscribing to block events.
type BlockFeed[
	BeaconBlockT BeaconBlock[BeaconBlockBodyT, DepositT, ExecutionPayloadT],
	BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT],
	BlockEventT BlockEvent[
		BeaconBlockT, BeaconBlockBodyT, DepositT, ExecutionPayloadT],
	DepositT interface{ GetIndex() uint64 },
	ExecutionPayloadT interface{ GetNumber() math.U64 },
	SubscriptionT interface {
		Unsubscribe()
	}] interface {
	Subscribe(chan<- (BlockEventT)) SubscriptionT
}

// Contract is the ABI for the deposit contract.
type Contract[DepositT any] interface {
	// ReadDeposits reads deposits from the deposit contract.
	ReadDeposits(
		ctx context.Context,
		blockNumber math.U64,
	) ([]DepositT, error)
}

// Deposit is an interface for deposits.
type Deposit[DepositT, WithdrawalCredentialsT any] interface {
	// New creates a new deposit.
	New(
		crypto.BLSPubkey,
		WithdrawalCredentialsT,
		math.U64,
		crypto.BLSSignature,
		uint64,
	) DepositT
}

// EthClient is an interface for interacting with the Ethereum 1.0 client.
type EthClient interface {
	BlockByNumber(
		ctx context.Context,
		number *big.Int,
	) (*engineprimitives.Block, error)
}

// Store defines the interface for managing deposit operations.
type Store[DepositT any] interface {
	// PruneFromInclusive prunes the deposit store from the given
	// index for N indexes.
	PruneFromInclusive(index uint64, numPrune uint64) error
	// EnqueueDeposits adds a list of deposits to the deposit store.
	EnqueueDeposits(deposits []DepositT) error
}

type StorageBackend[
	AvailabilityStoreT any,
	BeaconStateT any,
	BlobSidecarsT any,
	DepositStoreT Store[DepositT],
	DepositT any,
] interface {
	// DepositStore returns the deposit store for the given context.
	DepositStore(context.Context) DepositStoreT
}
