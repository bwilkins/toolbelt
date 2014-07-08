package command

import (
  "flag"
  "fmt"
  "os"
  "strings"
)

//Stolen mostly from go command source code
type Command struct {
  //Run is the method to run the command
  Run func(cmd *Command, args []string)

  //UsageLine is the one-liner for usage of this command
  UsageLine string
  //ShortDesc is the short descripton of the command
  ShortDesc string
  //LongDesc is the potentially long description of the command
  LongDesc  string

  //Flags holds the flag state for this command
  Flags     flag.FlagSet
}

func (c *Command) Usage() {
  fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
  fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.LongDesc))
  os.Exit(2)
}
