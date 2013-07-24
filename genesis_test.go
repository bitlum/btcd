// Copyright (c) 2013 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcwire_test

import (
	"bytes"
	"github.com/conformal/btcwire"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	pver := uint32(60002)

	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := btcwire.GenesisBlock.BtcEncode(&buf, pver)
	if err != nil {
		t.Errorf("TestGenesisBlock: %v", err)
		return
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Errorf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
		return
	}

	// Check hash of the block against expected hash.
	hash, err := btcwire.GenesisBlock.Header.BlockSha(pver)
	if err != nil {
		t.Errorf("BlockSha: %v", err)
	}
	if !btcwire.GenesisHash.IsEqual(&hash) {
		t.Errorf("TestGenesisBlock: Genesis block hash does not appear valid - "+
			"got %v, want %v", spew.Sdump(hash),
			spew.Sdump(btcwire.GenesisHash))
		return
	}
}

// TestTestNetGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
func TestTestNetGenesisBlock(t *testing.T) {
	pver := uint32(60002)

	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := btcwire.TestNetGenesisBlock.BtcEncode(&buf, pver)
	if err != nil {
		t.Errorf("TestTestNetGenesisBlock: %v", err)
		return
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testNetGenesisBlockBytes) {
		t.Errorf("TestTestNetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testNetGenesisBlockBytes))
		return
	}

	// Check hash of the block against expected hash.
	hash, err := btcwire.TestNetGenesisBlock.Header.BlockSha(pver)
	if err != nil {
		t.Errorf("BlockSha: %v", err)
	}
	if !btcwire.TestNetGenesisHash.IsEqual(&hash) {
		t.Errorf("TestTestNetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(btcwire.TestNetGenesisHash))
		return
	}
}

// TestTestNet3GenesisBlock tests the genesis block of the test network (version
// 3) for validity by checking the encoded bytes and hashes.
func TestTestNet3GenesisBlock(t *testing.T) {
	pver := uint32(60002)

	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := btcwire.TestNet3GenesisBlock.BtcEncode(&buf, pver)
	if err != nil {
		t.Errorf("TestTestNet3GenesisBlock: %v", err)
		return
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testNet3GenesisBlockBytes) {
		t.Errorf("TestTestNet3GenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testNet3GenesisBlockBytes))
		return
	}

	// Check hash of the block against expected hash.
	hash, err := btcwire.TestNet3GenesisBlock.Header.BlockSha(pver)
	if err != nil {
		t.Errorf("BlockSha: %v", err)
	}
	if !btcwire.TestNet3GenesisHash.IsEqual(&hash) {
		t.Errorf("TestTestNet3GenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(btcwire.TestNet3GenesisHash))
		return
	}
}

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 60002.
var genesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0x29, 0xab, 0x5f, 0x49, /* |K.^J)._I| */
	0xff, 0xff, 0x00, 0x1d, 0x1d, 0xac, 0x2b, 0x7c, /* |......+|| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |..M.....| */
	0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, /* |..EThe T| */
	0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2f, /* |imes 03/| */
	0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, 0x39, /* |Jan/2009| */
	0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6c, /* | Chancel| */
	0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x62, /* |lor on b| */
	0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, /* |rink of | */
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x62, /* |second b| */
	0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, 0x66, /* |ailout f| */
	0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x73, /* |or banks| */
	0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, 0x05, /* |........| */
	0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, 0x04, /* |*....CA.| */
	0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, 0x27, /* |g....UH'| */
	0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, 0x10, /* |.g..q0..| */
	0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, 0xa6, /* |\..(.9..| */
	0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, 0xb6, /* |yb...a..| */
	0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, 0xc4, /* |I..?L.8.| */
	0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, 0xde, /* |.U......| */
	0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, 0x57, /* |\8M....W| */
	0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, 0x5f, /* |.Lp+k.._|*/
	0xac, 0x00, 0x00, 0x00, 0x00, /* |.....|    */
}

// testNetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the regression test network as of protocol version 60002.
var testNetGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0xda, 0xe5, 0x49, 0x4d, /* |K.^J)._I| */
	0xff, 0xff, 0x7f, 0x20, 0x02, 0x00, 0x00, 0x00, /* |......+|| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |..M.....| */
	0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, /* |..EThe T| */
	0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2f, /* |imes 03/| */
	0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, 0x39, /* |Jan/2009| */
	0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6c, /* | Chancel| */
	0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x62, /* |lor on b| */
	0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, /* |rink of | */
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x62, /* |second b| */
	0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, 0x66, /* |ailout f| */
	0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x73, /* |or banks| */
	0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, 0x05, /* |........| */
	0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, 0x04, /* |*....CA.| */
	0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, 0x27, /* |g....UH'| */
	0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, 0x10, /* |.g..q0..| */
	0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, 0xa6, /* |\..(.9..| */
	0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, 0xb6, /* |yb...a..| */
	0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, 0xc4, /* |I..?L.8.| */
	0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, 0xde, /* |.U......| */
	0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, 0x57, /* |\8M....W| */
	0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, 0x5f, /* |.Lp+k.._|*/
	0xac, 0x00, 0x00, 0x00, 0x00, /* |.....|    */
}

// testNet3GenesisBlockBytes are the wire encoded bytes for the genesis block of
// the test network (version 3) as of protocol version 60002.
var testNet3GenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0xda, 0xe5, 0x49, 0x4d, /* |K.^J)._I| */
	0xff, 0xff, 0x00, 0x1d, 0x1a, 0xa4, 0xae, 0x18, /* |......+|| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |..M.....| */
	0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, /* |..EThe T| */
	0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2f, /* |imes 03/| */
	0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, 0x39, /* |Jan/2009| */
	0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6c, /* | Chancel| */
	0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x62, /* |lor on b| */
	0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, /* |rink of | */
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x62, /* |second b| */
	0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, 0x66, /* |ailout f| */
	0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x73, /* |or banks| */
	0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, 0x05, /* |........| */
	0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, 0x04, /* |*....CA.| */
	0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, 0x27, /* |g....UH'| */
	0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, 0x10, /* |.g..q0..| */
	0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, 0xa6, /* |\..(.9..| */
	0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, 0xb6, /* |yb...a..| */
	0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, 0xc4, /* |I..?L.8.| */
	0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, 0xde, /* |.U......| */
	0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, 0x57, /* |\8M....W| */
	0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, 0x5f, /* |.Lp+k.._|*/
	0xac, 0x00, 0x00, 0x00, 0x00, /* |.....|    */
}
