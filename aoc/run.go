package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/varbrad/advent-of-code-2021/utils"
)

var regex = regexp.MustCompile(`^day([0-9]+)$`)

func main() {
	result, err := filepath.Glob("./day*")
	if err != nil {
		panic(err)
	}

	sort.SliceStable(result, func(a, b int) bool {
		aR := regex.FindStringSubmatch(result[a])
		bR := regex.FindStringSubmatch(result[b])

		dayA := utils.ToInteger(aR[1])
		dayB := utils.ToInteger(bR[1])
		return dayA < dayB
	})

	for _, day := range result {
		cmd := exec.Command("go", "run", "./"+day)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("There was an error running day", day)
		}
	}
}
