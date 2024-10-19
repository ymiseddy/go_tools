package idhash_test

import (
	"bytes"
	"testing"

	"github.com/ymiseddy/go_tools/idhash"
)

var samplePassword = "chavez"
var sampleIncorrectPassword = "dhavez"
var sampleKeyLength uint32 = 64

func TestArgon2IdHashGenerator_NoPepperSimpleHashVerifies(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	hash, err := hasher.GenerateHashBase64(samplePassword)
	if err != nil {
		t.Errorf("Failed to generate hash: %v", err)
	}

	verified, err := hasher.VerifyBase64(samplePassword, hash)
	if err != nil {
		t.Errorf("Failed to verify: %v", err)
	}
	if !verified {
		t.Fail()
	}

}

func TestArgon2IdHashGenerator_NoPepperSimpleHashDoesNotVerifyWhenWrongPassword(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	hash, err := hasher.GenerateHashBase64(samplePassword)
	if err != nil {
		t.Errorf("Failed to generate hash: %v", err)
	}

	verified, err := hasher.VerifyBase64(sampleIncorrectPassword, hash)
	if err != nil {
		t.Errorf("Failed to verify: %v", err)
	}

	if verified {
		t.Fail()
	}
}

func TestArgon2IdHashGenerator_WithPepperHashVerifies(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	hasher.Pepper = []byte{23, 21, 24, 17, 19, 21, 22, 27}

	hash, err := hasher.GenerateHashBase64(samplePassword)
	if err != nil {
		t.Errorf("Failed to generate hash: %v", err)
	}

	verified, err := hasher.VerifyBase64(samplePassword, hash)
	if err != nil {
		t.Errorf("Failed to verify: %v", err)
	}

	if !verified {
		t.Fail()
	}
}

func TestArgon2IdHashGenerator_WithPepperHashFailsToVerifyWithIncorrectPepper(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	hasher.Pepper = []byte{23, 21, 24, 17, 19, 21, 22, 27}

	hash, err := hasher.GenerateHashBase64(samplePassword)
	if err != nil {
		t.Errorf("Failed to generate hash: %v", err)
	}

	hasher.Pepper = []byte{23, 21, 24, 17, 19, 21, 22, 77}
	verified, err := hasher.VerifyBase64(samplePassword, hash)
	if err != nil {
		t.Errorf("Failed to verify: %v", err)
	}

	if verified {
		t.Fail()
	}
}

func TestArgon2IdHashGenerator_RawHashWithoutSalt_ShouldGenerateTheSpecifiedKeySize(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	key := hasher.GenerateKeyFromBytes([]byte("Hello"), sampleKeyLength, nil)

	if uint32(len(key)) != sampleKeyLength {
		t.Fail()
	}
}

func TestArgon2IdHashGenerator_RawHashWithoutSalt_ShouldGenerateTheSameKeyForTheSameBytes(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	key := hasher.GenerateKeyFromBytes([]byte("Hello"), sampleKeyLength, nil)
	keyCheck := hasher.GenerateKeyFromBytes([]byte("Hello"), sampleKeyLength, nil)

	if !bytes.Equal(key, keyCheck) {
		t.Fail()
	}
}

func TestArgon2IdHashGenerator_RawHashWithoutSalt_ShouldNotGenerateTheSameKeyForDifferingBytes(t *testing.T) {
	hasher := idhash.DefaultArgon2Id
	key := hasher.GenerateKeyFromBytes([]byte("Hello"), sampleKeyLength, nil)
	key2 := hasher.GenerateKeyFromBytes([]byte("Hello2"), sampleKeyLength, nil)

	if bytes.Equal(key, key2) {
		t.Fail()
	}
}

func TestArgon2IdHashGenerator_ImplementsHashGenerator(t *testing.T) {
	// This test will fail to build if we ever fail to implement the interface.
	var _ idhash.HashGenerator = idhash.DefaultArgon2Id
}
