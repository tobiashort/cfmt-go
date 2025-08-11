package main

import (
	"fmt"
	"os"

	"github.com/tobiashort/ansi-go"
	"github.com/tobiashort/cfmt-go"
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

	fmt.Printf("\n")

	fmt.Print(cfmt.Sprint("This ", "is ", "#r{red}\n"))
	fmt.Print(cfmt.Sprint("This ", "is ", "#g{green}\n"))
	fmt.Print(cfmt.Sprint("This ", "is ", "#y{yellow}\n"))
	fmt.Print(cfmt.Sprint("This ", "is ", "#b{blue}\n"))
	fmt.Print(cfmt.Sprint("This ", "is ", "#p{purple}\n"))
	fmt.Print(cfmt.Sprint("This ", "is ", "#c{cyan}\n"))

	fmt.Printf("\n")

	fmt.Print(cfmt.Sprintf("This is #r{%s}\n", "red"))
	fmt.Print(cfmt.Sprintf("This is #g{%s}\n", "green"))
	fmt.Print(cfmt.Sprintf("This is #y{%s}\n", "yellow"))
	fmt.Print(cfmt.Sprintf("This is #b{%s}\n", "blue"))
	fmt.Print(cfmt.Sprintf("This is #p{%s}\n", "purle"))
	fmt.Print(cfmt.Sprintf("This is #c{%s}\n", "cyan"))

	fmt.Printf("\n")

	fmt.Print(cfmt.Sprintln("This", "is", "#r{red}"))
	fmt.Print(cfmt.Sprintln("This", "is", "#g{green}"))
	fmt.Print(cfmt.Sprintln("This", "is", "#y{yellow}"))
	fmt.Print(cfmt.Sprintln("This", "is", "#b{blue}"))
	fmt.Print(cfmt.Sprintln("This", "is", "#p{purple}"))
	fmt.Print(cfmt.Sprintln("This", "is", "#c{cyan}"))

	fmt.Printf("\n")

	cfmt.Println("This", "is", "#rB{red and bold}")
	cfmt.Println("This", "is", "#gU{green and underlined}")
	cfmt.Println("This", "is", "#yR{yellow reversed}")

	fmt.Printf("\n")

	cfmt.Begin(ansi.Purple)
	fmt.Println("what follows now  is...")
	fmt.Println("...in purple.")
	cfmt.End()
	fmt.Println("and now back to normal")
}
