# hcpxread

`hcpxread` is an interactive tool made to view, parse, and export `.hccapx` files.
You can learn more about the HCCAPX format [from the official docs](https://hashcat.net/wiki/doku.php?id=hccapx).

Long story short,
>hccapx is a custom format, specifically developed for hashcat.

`hcpxread` was designed based on the official HCCAPX specifications:
![HCCAPX specifications](https://hashcat.net/wiki/lib/exe/fetch.php?cache=&media=hccapx_specs.jpg)

## Features
* Interactive menu
* Reads and outputs AP data
* Shows summary of the loaded access points

## Usage 
```
$ go get github.com/vlad-s/hcpxread
$ hcpxread
   _                                       _
  | |__   ___ _ ____  ___ __ ___  __ _  __| |
  | '_ \ / __| '_ \ \/ / '__/ _ \/ _` |/ _` |
  | | | | (__| |_) >  <| | |  __/ (_| | (_| |
  |_| |_|\___| .__/_/\_\_|  \___|\__,_|\__,_|
             |_|
  
  Usage of hcpxread:
    -capture file
      	The HCCAPX file to read
    -debug
      	Show additional, debugging info
```

## Example
```
$ hcpxread -capture wpa.hccapx
INFO[0000] Opened file for reading                       name=wpa.hccapx size="6.5 KB"
INFO[0000] Searching for HCPX headers...
INFO[0000] Finished searching for headers                indexes=17
INFO[0000] Summary: 17 networks, 0 WPA/17 WPA2, 16 unique APs

1.  [WPA2] XXX                B0:48:7A:BF:07:A4
2.  [WPA2] XXXXX              08:10:77:5B:AC:ED
...
17. [WPA2] XXXXXX             64:70:02:9E:4D:1A
0.  Exit

network > 1

Message Pair                  |ESSID Length
0 - replay counter matching   |3

Key Version   |ESSID   |BSSID               |Client MAC
WPA2          |XXX     |B0:48:7A:BF:07:A4   |88:9F:FA:89:10:2E

```

## To do
* Export individual or a range of APs to an external file
* Add more data in the output
* ~~Add the message pair table~~
* Debugging flag should make the output more verbose