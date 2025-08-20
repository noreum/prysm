package stateutil

import (
	"fmt"

	"github.com/OffchainLabs/prysm/v6/encoding/ssz"
)

// ExecutionPayloadAvailabilityRoot computes the merkle root of an execution payload availability bitvector.
func ExecutionPayloadAvailabilityRoot(bitvector []byte) ([32]byte, error) {
	if len(bitvector)%32 != 0 {
		return [32]byte{}, fmt.Errorf("bitvector length %d not multiple of 32", len(bitvector))
	}

	chunks := make([][32]byte, len(bitvector)/32)
	for i := 0; i < len(chunks); i++ {
		start := i * 32
		chunks[i] = *(*[32]byte)(bitvector[start : start+32])
	}

	root, err := ssz.BitwiseMerkleize(chunks, uint64(len(chunks)), uint64(len(chunks)))
	if err != nil {
		return [32]byte{}, fmt.Errorf("could not merkleize execution payload availability: %w", err)
	}

	return root, nil
}
