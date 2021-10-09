package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	fmt.Println("Welcome on csv2birthcal !")
	js.Global().Set("parseCSV", parseCSVWrapper())
	<-make(chan bool) // avoids "Go program has already exited" error
}

func parseCSVWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid call, one string argument is expected"
		}

		csv := args[0].String()
		fmt.Println("csv" + csv)
		birthdays, err := fromCSV(strings.NewReader(csv))
		if err != nil {
			return fmt.Errorf("Could not parse CSV: %w", err).Error()
		}

		return birthdays.toIcal()
	})
}
