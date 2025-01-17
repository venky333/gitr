package cli

import log "github.com/sirupsen/logrus"

type Flag string

const (
	Dry    Flag = "dry"
	CreDir Flag = "create-dir"
)

func HandleFlagErr(err error, flag Flag) {
	if err != nil {
		log.Fatalf("error parsing %s flag. err %v", flag, err)
	}
}
