// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package job

import (
	"context"

	nfpb "openpitrix.io/notification/pkg/pb"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/util/jsonutil"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

type Client struct {
	nfpb.NotificationClient
}

func NewClient() (*Client, error) {
	conn, err := manager.NewClient(constants.NotificationHost, constants.NotificationPort)
	if err != nil {
		return nil, err
	}
	return &Client{
		NotificationClient: nfpb.NewNotificationClient(conn),
	}, nil
}

func (c *Client) SendEmailNotification(ctx context.Context, title, content, owner string, addresses []string) error {
	_, err := c.CreateNotification(ctx, &nfpb.CreateNotificationRequest{
		ContentType: pbutil.ToProtoString(constants.NfContentTypeInvite),
		Title:       pbutil.ToProtoString(title),
		Content:     pbutil.ToProtoString(content),
		Owner:       pbutil.ToProtoString(owner),
		AddressInfo: pbutil.ToProtoString(
			jsonutil.ToString(map[string][]string{
				constants.NfTypeEmail: addresses,
			}),
		),
	})

	if err != nil {
		logger.Error(ctx, "Failed to send email, %+v", err)
	}

	return err
}
