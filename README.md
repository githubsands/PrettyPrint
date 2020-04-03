Using pretty printer:

```golang
import (
    prettyprint "github.com/githubsands/PrettyPrint"
)

var printer *prettyprint.Printer

type Object struct {
    Name string
    Age int
}

func main() {
	printer := prettyprint.NewPrinter(prettyprint.PrinterOptions{CountFunction: true})
	printer.Start()
	printer.PrintCheck()
	
	testObject := Object{Name: Github, Age: 12}
	
	printer.PrintVar(testObject)
	
	printer.PrintCheck()
}
```
