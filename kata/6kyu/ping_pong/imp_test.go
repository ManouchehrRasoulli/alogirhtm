package kata

import "testing"

func Test_PingPong(t *testing.T) {
	t.Log(PingPong("ping-pong-ping-pong-bonk-bing-doof"))                                         // ping
	t.Log(PingPong("pong-ping-dong-ping-pong-tink-bonk-pong-ping-doof"))                          // pong
	t.Log(PingPong("pong-ping-bink-ping-pong-donk"))                                              // ping
	t.Log(PingPong("pong-ping-bink-ping-pong-donk"))                                              // ping
	t.Log(PingPong("ping-pong-ping-pong-ping-pong-ping-pong-dong-tang-bing-tink-donk-donk-donk")) // ping
}
