package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"sync"
	"sync/atomic"

	"github.com/go-rat/utils/convert"
)

// Copyright © 2014, Roger Peppe
// github.com/rogpeppe/fastuuid
// All rights reserved.

var (
	uuidSeed    [24]byte
	uuidCounter uint64
	uuidSetup   sync.Once
)

// UUID generates an universally unique identifier (UUID)
// WARNING: This implementation is not cryptographically secure and MUST NOT be used for key generation
func UUID() string {
	// Setup seed & counter once
	uuidSetup.Do(func() {
		if _, err := rand.Read(uuidSeed[:]); err != nil {
			return
		}
		uuidCounter = binary.LittleEndian.Uint64(uuidSeed[:8])
	})
	if atomic.LoadUint64(&uuidCounter) <= 0 {
		return "00000000-0000-0000-0000-000000000000"
	}
	// first 8 bytes differ, taking a slice of the first 16 bytes
	x := atomic.AddUint64(&uuidCounter, 1)
	_uuid := uuidSeed
	binary.LittleEndian.PutUint64(_uuid[:8], x)
	_uuid[6], _uuid[9] = _uuid[9], _uuid[6]

	// RFC4122 v4
	_uuid[6] = (_uuid[6] & 0x0f) | 0x40
	_uuid[8] = _uuid[8]&0x3f | 0x80

	// create UUID representation of the first 128 bits
	b := make([]byte, 36)
	hex.Encode(b[0:8], _uuid[0:4])
	b[8] = '-'
	hex.Encode(b[9:13], _uuid[4:6])
	b[13] = '-'
	hex.Encode(b[14:18], _uuid[6:8])
	b[18] = '-'
	hex.Encode(b[19:23], _uuid[8:10])
	b[23] = '-'
	hex.Encode(b[24:], _uuid[10:16])

	return convert.UnsafeString(b)
}
