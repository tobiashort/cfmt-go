package main

import (
	"fmt"
	"os"

	"github.com/tobiashort/cfmt"
)

func main() {
	cfmt.Print("This ", "is ", "#r{red}\n")
	cfmt.Print("This ", "is ", "#g{green}\n")
	cfmt.Print("This ", "is ", "#y{yellow}\n")
	cfmt.Print("This ", "is ", "#b{blue}\n")
	cfmt.Print("This ", "is ", "#p{purple}\n")
	cfmt.Print("This ", "is ", "#c{cyan}\n")

	fmt.Printf("\n")

	cfmt.Printf("This is #r{%s}\n", "red")
	cfmt.Printf("This is #g{%s}\n", "green")
	cfmt.Printf("This is #y{%s}\n", "yellow")
	cfmt.Printf("This is #b{%s}\n", "blue")
	cfmt.Printf("This is #p{%s}\n", "purle")
	cfmt.Printf("This is #c{%s}\n", "cyan")

	fmt.Printf("\n")

	cfmt.Println("This", "is", "#r{red}")
	cfmt.Println("This", "is", "#g{green}")
	cfmt.Println("This", "is", "#y{yellow}")
	cfmt.Println("This", "is", "#b{blue}")
	cfmt.Println("This", "is", "#p{purple}")
	cfmt.Println("This", "is", "#c{cyan}")

	fmt.Printf("\n")

	cfmt.CPrint("c", "This ", "is ", "#r{red}\n")
	cfmt.CPrint("p", "This ", "is ", "#g{green}\n")
	cfmt.CPrint("b", "This ", "is ", "#y{yellow}\n")
	cfmt.CPrint("y", "This ", "is ", "#b{blue}\n")
	cfmt.CPrint("g", "This ", "is ", "#p{purple}\n")
	cfmt.CPrint("r", "This ", "is ", "#c{cyan}\n")

	fmt.Printf("\n")

	cfmt.CPrintf("c", "This is #r{red}\n")
	cfmt.CPrintf("p", "This is #g{green}\n")
	cfmt.CPrintf("b", "This is #y{yellow}\n")
	cfmt.CPrintf("y", "This is #b{blue}\n")
	cfmt.CPrintf("g", "This is #p{purple}\n")
	cfmt.CPrintf("r", "This is #c{cyan}\n")

	fmt.Printf("\n")

	cfmt.CPrintln("c", "This is #r{red}")
	cfmt.CPrintln("p", "This is #g{green}")
	cfmt.CPrintln("b", "This is #y{yellow}")
	cfmt.CPrintln("y", "This is #b{blue}")
	cfmt.CPrintln("g", "This is #p{purple}")
	cfmt.CPrintln("r", "This is #c{cyan}")

	fmt.Printf("\n")

	cfmt.Fprint(os.Stdout, "This ", "is ", "#r{red}\n")
	cfmt.Fprint(os.Stdout, "This ", "is ", "#g{green}\n")
	cfmt.Fprint(os.Stdout, "This ", "is ", "#y{yellow}\n")
	cfmt.Fprint(os.Stdout, "This ", "is ", "#b{blue}\n")
	cfmt.Fprint(os.Stdout, "This ", "is ", "#p{purple}\n")
	cfmt.Fprint(os.Stdout, "This ", "is ", "#c{cyan}\n")

	fmt.Printf("\n")

	cfmt.Fprintf(os.Stdout, "This is #r{%s}\n", "red")
	cfmt.Fprintf(os.Stdout, "This is #g{%s}\n", "green")
	cfmt.Fprintf(os.Stdout, "This is #y{%s}\n", "yellow")
	cfmt.Fprintf(os.Stdout, "This is #b{%s}\n", "blue")
	cfmt.Fprintf(os.Stdout, "This is #p{%s}\n", "purle")
	cfmt.Fprintf(os.Stdout, "This is #c{%s}\n", "cyan")

	fmt.Printf("\n")

	cfmt.Fprintln(os.Stdout, "This", "is", "#r{red}")
	cfmt.Fprintln(os.Stdout, "This", "is", "#g{green}")
	cfmt.Fprintln(os.Stdout, "This", "is", "#y{yellow}")
	cfmt.Fprintln(os.Stdout, "This", "is", "#b{blue}")
	cfmt.Fprintln(os.Stdout, "This", "is", "#p{purple}")
	cfmt.Fprintln(os.Stdout, "This", "is", "#c{cyan}")
}
