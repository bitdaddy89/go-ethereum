// Copyright 2020 The go-etherdata Authors
// This file is part of go-etherdata.
//
// go-etherdata is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-etherdata is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-etherdata. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"net"

	"github.com/crypyto-panel/go-etherdata/cmd/devp2p/internal/etdtest"
	"github.com/crypyto-panel/go-etherdata/crypto"
	"github.com/crypyto-panel/go-etherdata/internal/utesting"
	"github.com/crypyto-panel/go-etherdata/p2p"
	"github.com/crypyto-panel/go-etherdata/p2p/rlpx"
	"github.com/crypyto-panel/go-etherdata/rlp"
	"gopkg.in/urfave/cli.v1"
)

var (
	rlpxCommand = cli.Command{
		Name:  "rlpx",
		Usage: "RLPx Commands",
		Subcommands: []cli.Command{
			rlpxPingCommand,
			rlpxEthTestCommand,
		},
	}
	rlpxPingCommand = cli.Command{
		Name:   "ping",
		Usage:  "ping <node>",
		Action: rlpxPing,
	}
	rlpxEthTestCommand = cli.Command{
		Name:      "etd-test",
		Usage:     "Runs tests against a node",
		ArgsUsage: "<node> <chain.rlp> <genesis.json>",
		Action:    rlpxEthTest,
		Flags: []cli.Flag{
			testPatternFlag,
			testTAPFlag,
		},
	}
)

func rlpxPing(ctx *cli.Context) error {
	n := getNodeArg(ctx)
	fd, err := net.Dial("tcp", fmt.Sprintf("%v:%d", n.IP(), n.TCP()))
	if err != nil {
		return err
	}
	conn := rlpx.NewConn(fd, n.Pubkey())
	ourKey, _ := crypto.GenerateKey()
	_, err = conn.Handshake(ourKey)
	if err != nil {
		return err
	}
	code, data, _, err := conn.Read()
	if err != nil {
		return err
	}
	switch code {
	case 0:
		var h etdtest.Hello
		if err := rlp.DecodeBytes(data, &h); err != nil {
			return fmt.Errorf("invalid handshake: %v", err)
		}
		fmt.Printf("%+v\n", h)
	case 1:
		var msg []p2p.DiscReason
		if rlp.DecodeBytes(data, &msg); len(msg) == 0 {
			return fmt.Errorf("invalid disconnect message")
		}
		return fmt.Errorf("received disconnect message: %v", msg[0])
	default:
		return fmt.Errorf("invalid message code %d, expected handshake (code zero)", code)
	}
	return nil
}

// rlpxEthTest runs the etd protocol test suite.
func rlpxEthTest(ctx *cli.Context) error {
	if ctx.NArg() < 3 {
		exit("missing path to chain.rlp as command-line argument")
	}
	suite, err := etdtest.NewSuite(getNodeArg(ctx), ctx.Args()[1], ctx.Args()[2])
	if err != nil {
		exit(err)
	}
	// check if given node supports etd66, and if so, run etd66 protocol tests as well
	is66Failed, _ := utesting.Run(utesting.Test{Name: "Is_66", Fn: suite.Is_66})
	if is66Failed {
		return runTests(ctx, suite.EthTests())
	}
	return runTests(ctx, suite.AllEthTests())
}
