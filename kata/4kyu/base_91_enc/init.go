package kata

const encodeStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'\""

var (
	encode    = [91]byte{}
	decodeMap = [256]byte{}
)

func init() {
	copy(encode[:], encodeStr)

	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xff // no data
	}

	for i := 0; i < len(encode); i++ {
		decodeMap[encode[i]] = byte(i)
	}
}
