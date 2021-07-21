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

package etd

import (
	"github.com/crypyto-panel/go-etherdata/core"
	"github.com/crypyto-panel/go-etherdata/core/forkid"
	"github.com/crypyto-panel/go-etherdata/p2p/enode"
	"github.com/crypyto-panel/go-etherdata/rlp"
)

// etdEntry is the "etd" ENR entry which advertises etd protocol
// on the discovery network.
type etdEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e etdEntry) ENRKey() string {
	return "etd"
}

// startEthEntryUpdate starts the ENR updater loop.
func (etd *Etherdata) startEthEntryUpdate(ln *enode.LocalNode) {
	var newHead = make(chan core.ChainHeadEvent, 10)
	sub := etd.blockchain.SubscribeChainHeadEvent(newHead)

	go func() {
		defer sub.Unsubscribe()
		for {
			select {
			case <-newHead:
				ln.Set(etd.currentEthEntry())
			case <-sub.Err():
				// Would be nice to sync with etd.Stop, but there is no
				// good way to do that.
				return
			}
		}
	}()
}

func (etd *Etherdata) currentEthEntry() *etdEntry {
	return &etdEntry{ForkID: forkid.NewID(etd.blockchain.Config(), etd.blockchain.Genesis().Hash(),
		etd.blockchain.CurrentHeader().Number.Uint64())}
}
