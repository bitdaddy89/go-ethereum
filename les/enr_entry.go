// Copyright 2019 The go-etherdata Authors
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

package les

import (
	"github.com/crypyto-panel/go-etherdata/core/forkid"
	"github.com/crypyto-panel/go-etherdata/p2p/dnsdisc"
	"github.com/crypyto-panel/go-etherdata/p2p/enode"
	"github.com/crypyto-panel/go-etherdata/rlp"
)

// lesEntry is the "les" ENR entry. This is set for LES servers only.
type lesEntry struct {
	// Ignore additional fields (for forward compatibility).
	VfxVersion uint
	Rest       []rlp.RawValue `rlp:"tail"`
}

func (lesEntry) ENRKey() string { return "les" }

// etdEntry is the "etd" ENR entry. This is redeclared here to avoid depending on package etd.
type etdEntry struct {
	ForkID forkid.ID
	Tail   []rlp.RawValue `rlp:"tail"`
}

func (etdEntry) ENRKey() string { return "etd" }

// setupDiscovery creates the node discovery source for the etd protocol.
func (etd *LightEtherdata) setupDiscovery() (enode.Iterator, error) {
	it := enode.NewFairMix(0)

	// Enable DNS discovery.
	if len(etd.config.EthDiscoveryURLs) != 0 {
		client := dnsdisc.NewClient(dnsdisc.Config{})
		dns, err := client.NewIterator(etd.config.EthDiscoveryURLs...)
		if err != nil {
			return nil, err
		}
		it.AddSource(dns)
	}

	// Enable DHT.
	if etd.udpEnabled {
		it.AddSource(etd.p2pServer.DiscV5.RandomNodes())
	}

	forkFilter := forkid.NewFilter(etd.blockchain)
	iterator := enode.Filter(it, func(n *enode.Node) bool { return nodeIsServer(forkFilter, n) })
	return iterator, nil
}

// nodeIsServer checks whether n is an LES server node.
func nodeIsServer(forkFilter forkid.Filter, n *enode.Node) bool {
	var les lesEntry
	var etd etdEntry
	return n.Load(&les) == nil && n.Load(&etd) == nil && forkFilter(etd.ForkID) == nil
}
