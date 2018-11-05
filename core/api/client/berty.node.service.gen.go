// this file was generated by protoc-gen-gotemplate
package client

import (
	"context"
	"io"

	"berty.tech/core/api/node"
	"berty.tech/core/api/p2p"
	"berty.tech/core/entity"
)

// Service returns the native gRPC client
func (c *Client) Node() node.ServiceClient {
	return node.NewServiceClient(c.conn)
}
func (c *Client) EventStream(ctx context.Context, input *node.EventStreamInput) ([]*p2p.Event, error) {
	stream, err := c.Node().EventStream(ctx, input)
	if err != nil {
		return nil, err
	}
	var entries []*p2p.Event
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
func (c *Client) EventList(ctx context.Context, input *node.EventListInput) ([]*p2p.Event, error) {
	stream, err := c.Node().EventList(ctx, input)
	if err != nil {
		return nil, err
	}
	var entries []*p2p.Event
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
func (c *Client) ContactList(ctx context.Context, input *node.ContactListInput) ([]*entity.Contact, error) {
	stream, err := c.Node().ContactList(ctx, input)
	if err != nil {
		return nil, err
	}
	var entries []*entity.Contact
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
func (c *Client) ConversationList(ctx context.Context, input *node.ConversationListInput) ([]*entity.Conversation, error) {
	stream, err := c.Node().ConversationList(ctx, input)
	if err != nil {
		return nil, err
	}
	var entries []*entity.Conversation
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
func (c *Client) MonitorBandwidth(ctx context.Context, input *p2p.BandwidthStats) ([]*p2p.BandwidthStats, error) {
	stream, err := c.Node().MonitorBandwidth(ctx, input)
	if err != nil {
		return nil, err
	}
	var entries []*p2p.BandwidthStats
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
func (c *Client) MonitorPeers(ctx context.Context, input *node.Void) ([]*p2p.Peer, error) {
	stream, err := c.Node().MonitorPeers(ctx, input)
	if err != nil {
		return nil, err
	}
	var entries []*p2p.Peer
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
