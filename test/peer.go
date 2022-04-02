package test

import (
	"math/rand"
	"testing"

	"github.com/riteshRcH/core/peer"

	mh "github.com/riteshRcH/go-edge-device-lib/multihash"
)

func RandPeerID() (peer.ID, error) {
	buf := make([]byte, 16)
	rand.Read(buf)
	h, _ := mh.Sum(buf, mh.SHA2_256, -1)
	return peer.ID(h), nil
}

func RandPeerIDFatal(t testing.TB) peer.ID {
	p, err := RandPeerID()
	if err != nil {
		t.Fatal(err)
	}
	return p
}
