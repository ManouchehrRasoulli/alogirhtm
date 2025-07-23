package kata

import "testing"

func TestLCS(t *testing.T) {
	t.Log(LCS("AGCAT", "GAC"))                           // GC
	t.Log(LCS("anothertest", "notatest"))                // nottest
	t.Log(LCS("nothardlythefinaltest", "zzzfinallyzzz")) // final
}
