// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package manager

import (
	"context"

	"github.com/gocraft/dbr"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/util/ctxutil"
)

func BuildOwnerPathFilter(ctx context.Context, ownerPaths ...string) dbr.Builder {
	s := ctxutil.GetSender(ctx)
	if s == nil {
		return nil
	}
	accessPath := string(s.GetAccessPath())

	if len(ownerPaths) == 0 {
		return db.Prefix(constants.ColumnOwnerPath, accessPath)
	} else {
		var ops []dbr.Builder
		for _, ownerPath := range ownerPaths {
			ops = append(ops, db.Prefix(constants.ColumnOwnerPath, ownerPath))
		}
		return db.Or(ops...)
	}
}
