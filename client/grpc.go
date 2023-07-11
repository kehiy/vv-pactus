package client

import (
	"context"
	"errors"
	"log"

	"github.com/k0kubun/pp"
	"github.com/pactus-project/pactus/crypto"
	"github.com/pactus-project/pactus/crypto/bls"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	NetworkClient     pactus.NetworkClient
	Conn              *grpc.ClientConn
}

func NewClient(endpoint string) (*Client, error) {
	conn, err := grpc.Dial(endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	pp.Println("connection established...")

	return &Client{
		NetworkClient:     pactus.NewNetworkClient(conn),
		Conn:              conn,
	}, nil

}

func (c *Client) GetNetworkInfo() (*pactus.GetNetworkInfoResponse, error) {
	networkInfo, err := c.NetworkClient.GetNetworkInfo(context.Background(), &pactus.GetNetworkInfoRequest{})
	if err != nil {
		log.Printf("error obtaining network information: %v", err)

		return nil, err
	}

	return networkInfo, nil
}

func (c *Client) GetPeerInfo(address string) (*pactus.PeerInfo, *bls.PublicKey, error) {
	networkInfo, _ := c.GetNetworkInfo()
	crypto.PublicKeyHRP = "tpublic"
	if networkInfo != nil {
		for _, p := range networkInfo.Peers {
			for _, key := range p.ConsensusKeys {
				pub, _ := bls.PublicKeyFromString(key)
				if pub != nil {
					if pub.Address().String() == address {
						return p, pub, nil
					}
				}
			}
		}
	}
	return nil, nil, errors.New("peer does not exist")
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
