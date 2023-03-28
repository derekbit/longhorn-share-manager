package client

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	rpc "github.com/longhorn/longhorn-share-manager/pkg/rpc"
	"github.com/longhorn/longhorn-share-manager/pkg/types"
)

type ShareManagerClient struct {
	Address string
}

func NewShareManagerClient(address string) *ShareManagerClient {
	return &ShareManagerClient{
		Address: address,
	}
}

func (cli *ShareManagerClient) FilesystemTrim(isEncryptedDevice bool) error {
	conn, err := grpc.Dial(cli.Address, grpc.WithInsecure())
	if err != nil {
		return errors.Wrapf(err, "cannot connect share manager service to %v", cli.Address)
	}
	defer conn.Close()

	client := rpc.NewShareManagerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	_, err = client.FilesystemTrim(ctx, &rpc.FilesystemTrimRequest{IsEncryptedDevice: isEncryptedDevice})
	return err
}
