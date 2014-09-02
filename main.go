package main

import (
  "os"
  "os/user"
  "path"
  "fmt"

  "github.com/bwilkins/aws"
  "github.com/bwilkins/toolbelt/config"
  "github.com/bwilkins/toolbelt/commands"
)


func main() {
  usr, _ := user.Current()
  config.SetConfig(path.Join(usr.HomeDir, ".toolbelt.yml"))
  aws.SetAccessCredentials( aws.Credentials{config.Config.OpsWorks.AccessId, config.Config.OpsWorks.SecretKey} )

  if len(os.Args) > 1 {
    commandName := os.Args[1]
	commandArgs := os.Args[2:]

	commands.RunCommand(commandName, commandArgs)
  } else {
	fmt.Println("No command given, aborting.")
  }
}
