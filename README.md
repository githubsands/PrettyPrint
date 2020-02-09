Using pretty printer:

import (
    prettyprint "github.com/githubsands/PrettyPrint"
)

var printer *prettyprint.Printer

func init() {
	printer := prettyprint.NewPrinter(prettyprint.PrinterOptions{CountFunction: true})
	printer.Start()
}
