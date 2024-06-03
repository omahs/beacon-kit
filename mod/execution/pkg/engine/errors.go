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

package engine

import "github.com/berachain/beacon-kit/mod/errors"

var (
	// ErrExecutionClientDisconnected represents an error when
	/// the execution client is disconnected.
	ErrExecutionClientDisconnected = errors.New(
		"execution client disconnected")

	// ErrAcceptedSyncingPayloadStatus represents an error when
	// the payload status is SYNCING or ACCEPTED.
	ErrAcceptedSyncingPayloadStatus = errors.New(
		"payload status is SYNCING or ACCEPTED")

	// ErrInvalidPayloadStatus represents an error when the
	// payload status is INVALID.
	ErrInvalidPayloadStatus = errors.New(
		"payload status is INVALID")

	// ErrBadBlockProduced represents an error when the beacon
	// chain has produced a bad block.
	ErrBadBlockProduced = errors.New(
		"proposer has produced a bad block, RIP walrus",
	)
)
