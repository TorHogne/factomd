// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package meta_test

import (
	"encoding/hex"
	"testing"

	. "github.com/FactomProject/factomd/common/meta"
)

func TestNewBlockSigningKeyStruct(t *testing.T) {
	parts := []string{
		"00",
		"4E657720426C6F636B205369676E696E67204B6579",
		"888888d027c59579fc47a6fc6c4a5c0409c7c39bc38a86cb5fc0069978493762",
		"8473745873ec04073ecf005b0d2b6cfe2f05f88f025e0c0a83a40d1de696a9cb",
		"00000000495EAA80",
		"0125b0e7fd5e68b4dec40ca0cd2db66be84c02fe6404b696c396e3909079820f61",
		"0bb2cab2904a014bd915b276c350821620edb432ddfbceed3896e87e591a412712b7db6d8dad1a8313138ea919bbc9b7a1bd4ffe1d84d558b8a78ef7746f480d",
	}
	extIDs := [][]byte{}
	for _, v := range parts {
		b, _ := hex.DecodeString(v)
		extIDs = append(extIDs, b)
		//t.Logf("Len %v - %v", i, len(b))
	}
	msr := new(NewBlockSigningKeyStruct)
	err := msr.DecodeFromExtIDs(extIDs)
	if err != nil {
		t.Errorf("%v", err)
	}
	h := msr.GetChainID()
	if h.String() != "bcabf9375eeb361e1a6ca509c60763365fbde5a3c5d3d87911d254057b3aa84f" {
		t.Errorf("Wrong ChainID, expected bcabf9375eeb361e1a6ca509c60763365fbde5a3c5d3d87911d254057b3aa84f, got %v", h.String())
	}
}
