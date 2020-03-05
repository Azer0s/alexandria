package interpreter

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func DoPreprocessing(str, dir string) (code string, ttl uint32) {
	ttlRegex, _ := regexp.Compile("^#default_ttl ([0-9]+)$")
	includeRegex, _ := regexp.Compile("^#include (.+)$")

	ttl = 43200 //default TTL is 43200

	lines := strings.Split(str, "\n")

	for i := 0; i < len(lines); i++ {
		if ttlRegex.MatchString(lines[i]) {
			match, _ := strconv.Atoi(ttlRegex.FindStringSubmatch(lines[i])[1])
			ttl = uint32(match)
			lines = remove(lines, i)
			i--
		} else if includeRegex.MatchString(lines[i]) {
			match := includeRegex.FindStringSubmatch(lines[i])[1]
			b, err := ioutil.ReadFile(path.Join(dir, match))
			if err != nil {
				log.Fatal(err)
			}

			includeLines := strings.Split(string(b), "\n")

			lines = remove(lines, i)

			before := make([]string, i)
			after := make([]string, len(lines)-i)
			copy(before, lines[0:i])
			copy(after, lines[i:])

			lines = append(before, includeLines...)
			lines = append(lines, after...)
			i--
		}
	}

	code = strings.Join(lines, "\n")

	return
}
