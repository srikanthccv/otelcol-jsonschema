package components

import (
	"log"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

func KoanfLoad(b []byte, err error) *koanf.Koanf {
	if err != nil {
		log.Fatal(err)
	}
	var k = koanf.New(".")
	k.Load(rawbytes.Provider(b), json.Parser())
	return k
}
