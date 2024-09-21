package operator

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"terminal_hack/internal/constants"
	"terminal_hack/internal/messages"
	"terminal_hack/internal/player"
	"time"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Minute

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "terminal-hack"

type Operator struct {
	Messages        chan *pubsub.Message
	SelfPlayerState chan *interface{} // TODO: can this bw added in a select {} with .Next()?

	ctx    context.Context
	ps     *pubsub.PubSub
	topics map[string]*pubsub.Topic
	sub    *pubsub.Subscription

	self     peer.ID
	doneChan chan bool
	logger   *slog.Logger
}

func New(logger *slog.Logger, done chan bool) *Operator {
	o := new(Operator)
	o.logger = logger
	o.doneChan = done
	o.logger.Info("Constructing operator")
	return o
}
func (o *Operator) InitializePubsub(_player *player.Player) {
	// parse some flags to set our nickname and the room to join
	// nickFlag := flag.String("nick", "", "nickname to use in chat. will be generated if empty")
	// roomFlag := flag.String("room", "awesome-chat-room", "name of chat room to join")
	flag.Parse()

	o.ctx = context.Background()

	// create a new libp2p Host that listens on a random TCP port
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}

	o.self = h.ID()
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(o.ctx, h)
	if err != nil {
		panic(err)
	}
	o.ps = ps
	// setup local mDNS discovery
	if err := setupDiscovery(h); err != nil {
		panic(err)
	}
	o.logger.Info("Setting up mDNS")
	o.topics = make(map[string]*pubsub.Topic)
	o.Messages = make(chan *pubsub.Message)
	o.subscribeAndDispatch(o.ctx)
	o.logger.Info("Dispatched a local listener")
	// TODO: add check to see if there any peers
	// send new player
	if len(o.ps.ListPeers(constants.TOPIC)) > 0 {
		// TODO: ask for state here, including reordering container render incase of other peers
		o.SendMessage(
			messages.GameMessageTopic,
			messages.GameMessage{
				MessageType: messages.GameBoardRequestType,
				Data:        messages.GameBoardRequest{},
			},
		)
		o.SendMessage(
			messages.GameMessageTopic,
			messages.GameMessage{
				MessageType: messages.AddPlayerType,
				Data: messages.AddPlayer{
					Player: *_player.Clone(),
				},
			})
	}
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	slog.Info(fmt.Sprintf("discovered new peer %s\n", pi.ID.String()))
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		slog.Error("error connecting to peer %s: %s\n", pi.ID.String(), err)
	}
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(h host.Host) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, DiscoveryServiceTag, &discoveryNotifee{h: h})
	return s.Start()
}
func (o *Operator) subscribeAndDispatch(ctx context.Context) {
	topics := []string{messages.GameMessageTopic}
	for _, topic := range topics {
		_topic := o.getTopic(topic)
		sub, err := _topic.Subscribe()
		if err != nil {
			o.logger.Error(err.Error())
		}
		go readLoop(ctx, o.self, sub, o.Messages)
	}

}
func readLoop(ctx context.Context, id peer.ID, sub *pubsub.Subscription, msgs chan *pubsub.Message) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			slog.Error(
				fmt.Sprintf("Error encountered during receive: %s", err.Error()),
			)
		}
		slog.Info(msg.ReceivedFrom.String())
		slog.Info(id.String())
		// only forward messages delivered by others
		if msg.ReceivedFrom.String() == id.String() {
			slog.Info("Ignoring message from self")
			continue
		}
		slog.Info("Message received from subscription")
		msgs <- msg
	}
}
func (o *Operator) SendMessage(topic string, msg messages.GameMessage) {
	raw, err := json.Marshal(msg)
	o.logger.Info(
		fmt.Sprintf("Sending payload: %b", raw),
	)
	if err != nil {
		panic(err)
	}
	_topic := o.getTopic(topic)
	_topic.Publish(o.ctx, raw)
	o.logger.Info(
		fmt.Sprintf("Message{%s:%s} published", topic, msg.MessageType),
	)

}
func (o *Operator) getTopic(t string) *pubsub.Topic {
	if o.topics[t] == nil {
		_topic, err := o.ps.Join(t)
		if err != nil {
			panic(err)
		}
		o.topics[t] = _topic
	}
	return o.topics[t]
}

func (o *Operator) GetPeerCount() int {
	return len(o.ps.ListPeers(messages.GameMessageTopic))
}
