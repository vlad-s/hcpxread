package helpers

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/sirupsen/logrus"
	"github.com/vlad-s/hcpxread/structs"
)

var Logger = logrus.New()
var debug *bool

func SetDebugging(d bool) {
	if debug == nil {
		debug = new(bool)
	}
	*debug = d
}

func Debug() bool {
	if debug == nil {
		return false
	}
	return *debug
}

func ClearScreen(nl ...bool) {
	c := "\033[H\033[2J"
	if len(nl) == 1 && nl[0] {
		c += "\n"
	}
	fmt.Print(c)
}

func SearchHeaders(content []byte) (pos []int) {
	for i, j := 0, 0; i < len(content)-4; i++ {
		j = i + 4
		if bytes.Equal(content[i:j], structs.HcpxHeader) {
			if *debug {
				Logger.WithField("index", i).Debug("Got HCPX header")
			}
			pos = append(pos, i)
		}
	}
	return
}

func ParseHccapx(b []byte) (h structs.HccapxInstance) {
	essid := bytes.Replace(b[10:42], []byte{0}, []byte{}, -1)

	h = structs.HccapxInstance{
		Signature:   b[0:4],
		Version:     b[4:8],
		MessagePair: structs.MessagePair(b[8:9][0]),

		ESSIDLength: uint8(b[9:10][0]),
		ESSID:       string(essid),

		KeyVersion: structs.WPAKey(b[42:43][0]),
		HashValue:  b[43:59],

		StationMAC:   b[59:65],
		StationNonce: b[65:97],
		ClientMAC:    b[97:103],
		ClientNonce:  b[103:135],

		EAPOL: b[137:],
	}

	EAPOLLength := new(big.Int).SetBytes(b[135:137]).Uint64()
	h.EAPOLLength = uint16(EAPOLLength)

	return
}
