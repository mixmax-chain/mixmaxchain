// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://f2a0a5d26bd113fd02ee051d7dccbfe20b927e5293789b36f87b72b52657510fe9842087a91def572b66821ba8329e27069f302d5ddd5e05336b9159a50d289e@13.208.172.237:30311",
	"enode://edb0871c1a64c062098053e74486030ea1244b6ba361646edf7fcccd2666557d50e4fa22b3c206fce1c129dd75127ede2477bf06fe2215bb3c5f81b00549b8eb@35.241.119.30:30311",
	"enode://10344a545eeaf3b9828835c351ca5cc170d7d28ac5c8656b8247aba8dff814b230004bd22c096422769814eb6edc28815ad6b16713e6dde8ba9a6bc001a6da80@13.208.169.202:30311",
	"enode://3b76f1f1dc520f524d632f5f213ee3098c5a2215c7c95246453034dd0683b3b3db2e7090855a3bfb788993f96c8e48509ad2a6c34b56922e9891d39afcf119e1@34.124.209.36:30311",
	"enode://36ac7cfc29c6a4441e37f65967e8a7343cc4c4dbc29c5ea94da90510ec678147115ee683e0a59d4a08a0fe24ccb2430e5e4b600507efa9da0d9ba3df05daa8f4@34.163.55.112:30311",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
