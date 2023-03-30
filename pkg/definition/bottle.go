package definition

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Bottle() map[string]string {
	var (
		path string
		res  = map[string]string{}
	)

	dir, err := os.Getwd()
	if err != nil {
		log.Panic("could not get dir", err)
	}

	find := false

	for i := 0; i < 6; i++ {
		_ = filepath.Walk(dir, func(p string, info fs.FileInfo, err error) error {
			if strings.HasSuffix(p, ".bottle") {
				path = p
				find = true
			}
			return nil
		})

		if find {
			break
		}

		dir += "/.."
	}

	if path == "" {
		log.Panic("could not find .bottle file")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Panic("could not open .bottle file", err)
	}

	for _, value := range strings.Split(string(file), "\n") {
		split := strings.Split(value, ":")
		if len(split) != 2 {
			log.Panic("invalid key:value .bottle file declaration")
		}

		res[split[0]] = strings.TrimSpace(split[1])
	}

	return res
}
