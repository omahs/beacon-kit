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

package engineprimitives

import "github.com/berachain/beacon-kit/mod/errors"

var (
	// ErrInvalidTimestamp indicates that the provided timestamp is not valid.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// ErrInvalidRandao indicates that the provided RANDAO value is not valid.
	ErrInvalidRandao = errors.New("invalid randao")

	// ErrNilWithdrawals indicates that the withdrawals are in a
	// Capella versioned payload.
	ErrNilWithdrawals = errors.New("nil withdrawals post capella")

	// ErrEmptyPrevRandao indicates that the previous RANDAO value is empty.
	ErrEmptyPrevRandao = errors.New("empty randao")

	// ErrFailedToUnmarshalTx indicates that the transaction could not be
	// unmarshaled.
	ErrFailedToUnmarshalTx = errors.New("failed to unmarshal transaction")

	// ErrInvalidVersionedHash indicates that the versioned hash is invalid.
	ErrInvalidVersionedHash = errors.New("invalid versioned hash")

	// ErrMismatchedNumVersionedHashes indicates that the number of blobs in the
	// payload does not match the expected number.
	ErrMismatchedNumVersionedHashes = errors.New(
		"mismatch in number of versioned hashes",
	)

	// ErrPayloadBlockHashMismatch represents an error when the
	// block hash in the payload does not match from the assembled
	// block.
	ErrPayloadBlockHashMismatch = errors.New(
		"block hash in payload does not match assembled block",
	)
)
