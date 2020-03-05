package cfg

import (
	"github.com/Azer0s/alexandria/dns/cfg/interpreter"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
)

type Zones []interpreter.Zone

func Parse(file string) Zones {
	b, _ := ioutil.ReadFile(file)
	str := string(b)

	dir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		log.Fatal(err)
	}

	str, ttl := interpreter.DoPreprocessing(str, dir)
	zones := interpreter.ParseConfig(interpreter.LexConfig(str), ttl)
	return zones
}
