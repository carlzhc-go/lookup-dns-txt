///usr/bin/true; exec go run "$0" "$@"
// -*- mode: go -*-
package main

import (
	"net"
	"os"
	"strings"
	. "fmt"
)

const (
	hesiod =  "/etc/hesiod.conf"
)

var (
	hesiodSettings map[string]string
	catalog = "sshkey"
	user string
)

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}

	return v
}

func init() {
	hesiodSettings = make(map[string]string)
}

func main() {
	if len(os.Args) < 2 {
		Fprintf(os.Stderr, "Usage: %s [catalog] <name>\n", os.Args[0])
		os.Exit(1)
	} else if len(os.Args) > 2 {
		catalog = os.Args[1]
		user = os.Args[2]
	} else {
		user = os.Args[1]
	}


	hesiodContent := string(Must(os.ReadFile(hesiod)))
	for _, l := range strings.Split(hesiodContent, "\n") {
		if len(l) == 0 { continue }
		l = strings.TrimSpace(l)
		firstChar := []byte(l)[0]
		if firstChar == '#' {
			continue
		}

		kv := strings.Split(l, "=")
		if len(kv) == 2 {
			k := strings.TrimSpace(kv[0])
			v := strings.TrimSpace(kv[1])
			hesiodSettings[k] = v
		} else {
			Fprintf(os.Stderr, "unknow setting: %v\n", kv)
		}
	}

	record := user + "." + catalog + hesiodSettings["lhs"] + hesiodSettings["rhs"]
	txts := Must(net.LookupTXT(record))
	for _, txt := range txts {
		Printf("%s\n", txt)
	}
}
