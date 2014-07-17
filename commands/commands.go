package commands


import (
  "errors"
  "log"
  "github.com/bwilkins/toolbelt/command"
  "github.com/bwilkins/toolbelt/commands/ssh"
  "github.com/bwilkins/toolbelt/commands/stacks"
  "github.com/bwilkins/toolbelt/commands/layers"
)

func SelectCommand(commandName string) (*command.Command, error) {
  selectedCommand, exists := commandMap[commandName]
  if exists {
    return selectedCommand, nil
  }

  return nil, errors.New("Unrecognised command " + commandName)
}

func RunCommand(commandName  string, commandArgs []string) {
  selectedCommand, err := SelectCommand(commandName)
  if err != nil {
    log.Fatal(err.Error())
  }

  selectedCommand.Run(selectedCommand, commandArgs)
}

var commandMap map[string]*command.Command

func init() {
  commandMap = map[string]*command.Command{
    "ssh": ssh.SshCmd,
    "ssh-command": ssh.SshCmd,
    "stacks": stacks.StacksCmd,
    "layers": layers.LayersCmd,
  }
}
