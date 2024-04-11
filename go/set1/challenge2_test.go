package set1

import (
	"encoding/hex"
	"testing"
)

func TestChallenge2(t *testing.T) {
	a, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		t.Error("could not decode hex string")
	}
	b, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		t.Error("could not decode hex string")
	}
	ans, err := FixedXOR(a, b)
	if err != nil {
		t.Error("could not get XOR result")
	}
	actual := hex.EncodeToString(ans)
	expected := "746865206b696420646f6e277420706c6179"
	if actual != expected {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}
