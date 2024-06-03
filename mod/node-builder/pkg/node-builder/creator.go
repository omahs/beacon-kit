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

package nodebuilder

import (
	"io"
	"reflect"

	"cosmossdk.io/log"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/app"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/comet"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

// NodeBuilder is a struct that holds the.
func (nb *NodeBuilder[T]) AppCreator(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) T {
	// Check for goleveldb cause bad project.
	if appOpts.Get("app-db-backend") == "goleveldb" {
		panic("goleveldb is not supported")
	}

	app := *app.NewBeaconKitApp(
		logger, db, traceStore, true,
		appOpts,
		nb.appInfo.DepInjectConfig,
		nb.chainSpec,
		append(
			server.DefaultBaseappOptions(appOpts),
			func(bApp *baseapp.BaseApp) {
				bApp.SetParamStore(comet.NewConsensusParamsStore(nb.chainSpec))
			})...,
	)
	return reflect.ValueOf(app).Convert(
		reflect.TypeOf((*T)(nil)).Elem()).Interface().(T)
}
