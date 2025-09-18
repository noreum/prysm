package util

import (
	"testing"

	"github.com/OffchainLabs/prysm/v6/config/params"
	"github.com/OffchainLabs/prysm/v6/consensus-types/blocks"
	"github.com/OffchainLabs/prysm/v6/testing/require"
	"github.com/OffchainLabs/prysm/v6/time/slots"
)

func TestInclusionProofs(t *testing.T) {
	ds, err := slots.EpochStart(params.BeaconConfig().DenebForkEpoch)
	require.NoError(t, err)
	_, blobs := GenerateTestDenebBlockWithSidecar(t, [32]byte{}, ds, params.BeaconConfig().MaxBlobsPerBlock(ds))
	for i := range blobs {
		require.NoError(t, blocks.VerifyKZGInclusionProof(blobs[i]))
	}
}
