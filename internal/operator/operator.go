package operator

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
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
	Messages chan *interface{}

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	self     peer.ID
	doneChan chan bool
}

type GameMessage struct {
	MesssgeType uint32
	PlayerId uint32 // player id that commit action
	PlayerState Player  // this should be a deep copy of player
}

func Initialize(done chan bool) *Operator {
	o := new(Operator)
	o.doneChan = done
	return o
}
func (o *Operator) initializePubsub() {
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

	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(o.ctx, h)
	if err != nil {
		panic(err)
	}

	// setup local mDNS discovery
	if err := setupDiscovery(h); err != nil {
		panic(err)
	}
	subscribeAndDispatch(o.ctx, ps)
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
func subscribeAndDispatch(ctx context.Context, ps *pubsub.PubSub) {
	topic := "MESSAGE"
	_topic, _ := ps.Join(topic)
	sub, _ := _topic.Subscribe()
	go readLoop(ctx, sub)

}
func readLoop(ctx context.Context, sub *pubsub.Subscription) {
	for {
		msg, _ := sub.Next(ctx)
		// only forward messages delivered by others
		if msg.ReceivedFrom == cr.self {
			continue
		}
		switch msg.GetTopic() {
		case "MESSAGE":
			bytes := msg.GetData()
			payload := new(GameMessage)
			err := json.Unmarshal(bytes, payload)
			if err != nil {
				panic)(err)
			}
			break
		}
	}
}
func (o *Operator) SendMessage(topic string, msg interface{}) {
	raw, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	_topic, _ := o.ps.Join(topic)
	_topic.Publish(o.ctx, raw)

}
