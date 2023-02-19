package tsv

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var tIdRegexp = regexp.MustCompile(`tt[0]*`)
var nIdRegexp = regexp.MustCompile(`nm[0]*`)

const tab = "\t"

func LoopTsv(tsvPath string, yield func(uint, []string) error) error {
	f, err := os.Open(tsvPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	var lineno uint
	for fileScanner.Scan() {
		lineno = lineno + 1
		if lineno == 1 {
			continue
		}
		line := fileScanner.Text()
		rec := strings.Split(line, tab)

		if yield(lineno, rec) != nil {
			break
		}
	}

	return nil
}

func wcCountLines(path string) (int, error) {
	out, err := exec.Command("wc", "-l", path).Output()
	if err != nil {
		return 0, err
	}

	outStr := string(out)
	countStr := strings.Split(outStr, " ")[0]

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}

	return count - 1, nil
}

func tt2Int(tt string) uint {
	idStr := tIdRegexp.ReplaceAllString(tt, "")
	i, _ := strconv.ParseUint(idStr, 0, 32)
	return uint(i)
}

func nm2Int(nm string) uint {
	idStr := nIdRegexp.ReplaceAllString(nm, "")
	i, _ := strconv.ParseUint(idStr, 0, 32)
	return uint(i)
}
