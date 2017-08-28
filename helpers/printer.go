package helpers

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vlad-s/hcpxread/structs"
)

func PrintCommands() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Println()
	fmt.Fprintln(w, "99.\tExport")
	fmt.Fprintln(w, "0.\tExit")
	w.Flush()
}

func PrintInstances(h structs.HccapxInstances) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Println()
	for k, v := range h {
		fmt.Fprintf(w, "%d.\t[%s]\t%s\t%s\n", k+1, v.KeyVersion, v.ESSID, v.StationMAC)
	}
	w.Flush()
}

func PrintHccapx(h structs.HccapxInstance) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Key Version\tESSID\tESSID length\tBSSID\tClient MAC")
	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\n", h.KeyVersion, h.ESSID, h.ESSIDLength, h.StationMAC, h.ClientMAC)
	w.Flush()

	fmt.Println()
	fmt.Fprintln(w, "Handshake messages\tEAPOL Source\tAP message\tSTA message\tReplay counter match")
	mp := structs.MessagePairTable[h.MessagePair]
	fmt.Fprintf(w, "M%d + M%d\tM%d\tM%d\tM%d\t%v\n", mp.APMessage, mp.ClientMessage, mp.EAPOLSource,
		mp.APMessage, mp.ClientMessage, mp.ReplayCounterMatching)
	w.Flush()
}
