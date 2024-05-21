package bn254_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Skate-Org/AVS/lib/crypto/bls/bn254"
	"github.com/stretchr/testify/assert"
)

func TestBlsKeyCreation(t *testing.T) {
	tests := map[string]struct {
		keyPath  string
		password string
		wantErr  bool
	}{
		"valid bls key save": {
			keyPath:  "./operator_keys_test_directory/test.bls.key.json",
			password: "test",
			wantErr:  false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Cleanup(func() {
				dir := filepath.Dir(tt.keyPath)
				_ = os.RemoveAll(dir)
			})
			randomKey, err := bn254.GenRandomBlsKeys()
			assert.NoError(t, err)

			err = randomKey.SaveToFile(tt.keyPath, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			readKeyPair, err := bn254.ReadPrivateKeyFromFile(tt.keyPath, tt.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, randomKey, readKeyPair)
			}
		})
	}
}

