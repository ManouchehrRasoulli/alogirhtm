package _kyu

import "testing"

func TestSplitByWidth(t *testing.T) {
	t.Log(splitBySize("somekeytosomexts", 2))
}

func TestParseKey(t *testing.T) {
	t.Logf("%+v", parseKey("So"))
}

func TestEncode(t *testing.T) {
	t.Log(Encode("Ala has a cat", "gaderypoluki"))  // Gug hgs g cgt
	t.Log(Encode("ABCD", "gaderypoluki"))           // GBCE
	t.Log(Encode("gaderypoluki", "gaderypoluki"))   // agedyropulik
	t.Log(Encode("Hmdr nge brres", "regulaminowy")) // Hide our beers
}

func TestDecode(t *testing.T) {
	t.Log(Decode("Gug hgs g cgt", "gaderypoluki"))      // Ala has a cat
	t.Log(Decode("agedyropulik", "gaderypoluki"))       // gaderypoluki
	t.Log(Decode("Dkucr pu yhr ykbir", "politykarenu")) // Dance on the table
}
