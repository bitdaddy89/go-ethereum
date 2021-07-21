// Copyright 2015 The go-etherdata Authors
// This file is part of the go-etherdata library.
//
// The go-etherdata library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etherdata library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etherdata library. If not, see <http://www.gnu.org/licenses/>.

package etd

import (
	"math/big"
	"sync"
	"time"

	"github.com/crypyto-panel/go-etherdata/etd/protocols/etd"
	"github.com/crypyto-panel/go-etherdata/etd/protocols/snap"
)

// etdPeerInfo represents a short summary of the `etd` sub-protocol metadata known
// about a connected peer.
type etdPeerInfo struct {
	Version    uint     `json:"version"`    // Etherdata protocol version negotiated
	Difficulty *big.Int `json:"difficulty"` // Total difficulty of the peer's blockchain
	Head       string   `json:"head"`       // Hex hash of the peer's best owned block
}

// etdPeer is a wrapper around etd.Peer to maintain a few extra metadata.
type etdPeer struct {
	*etd.Peer
	snapExt *snapPeer // Satellite `snap` connection

	syncDrop *time.Timer   // Connection dropper if `etd` sync progress isn't validated in time
	snapWait chan struct{} // Notification channel for snap connections
	lock     sync.RWMutex  // Mutex protecting the internal fields
}

// info gathers and returns some `etd` protocol metadata known about a peer.
func (p *etdPeer) info() *etdPeerInfo {
	hash, td := p.Head()

	return &etdPeerInfo{
		Version:    p.Version(),
		Difficulty: td,
		Head:       hash.Hex(),
	}
}

// snapPeerInfo represents a short summary of the `snap` sub-protocol metadata known
// about a connected peer.
type snapPeerInfo struct {
	Version uint `json:"version"` // Snapshot protocol version negotiated
}

// snapPeer is a wrapper around snap.Peer to maintain a few extra metadata.
type snapPeer struct {
	*snap.Peer
}

// info gathers and returns some `snap` protocol metadata known about a peer.
func (p *snapPeer) info() *snapPeerInfo {
	return &snapPeerInfo{
		Version: p.Version(),
	}
}
