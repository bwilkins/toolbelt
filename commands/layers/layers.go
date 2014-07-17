package layers

import (
  "log"

  "github.com/bwilkins/toolbelt/command"
  "github.com/bwilkins/toolbelt/encoding/term_table"
  "github.com/bwilkins/aws/opsworks"
)

var LayersCmd = &command.Command{
  UsageLine: "stacks",
  ShortDesc: "connects to a remote server as specified on command-line",
  LongDesc:  ``,
}

var stackName string

func init() {
  LayersCmd.Run = runLayers

  LayersCmd.Flags.Usage = func() { LayersCmd.Usage() }

  LayersCmd.Flags.StringVar(&stackName, "s", "", "which stack to list instances for")
  LayersCmd.Flags.StringVar(&stackName, "stack", "", "which stack to list instances for")
}

func runLayers(cmd *command.Command, args []string) {
  cmd.Flags.Parse(args)

  if stackName == "" {
    log.Fatal("Cannot SSH without a stack name")
    return
  }

  stacksResponse, err := opsworks.DescribeStacks(opsworks.DescribeStacksRequest{})
  if err != nil {
    log.Fatal(err.Error())
  }
  var stackId string
  for _, stack := range stacksResponse.Stacks {
    if stack.Name == stackName {
      stackId = stack.StackId
      break
    }
  }

  if stackId == "" {
    log.Fatal("Could not find stack")
  }

  layersResponse, err := opsworks.DescribeLayers(opsworks.DescribeLayersRequest{ StackId: stackId })
  if err != nil {
    log.Fatal(err.Error())
  }

  table := term_table.NewTermTable("Stack Name", "Name", "Short Name", "ID")
  for _, layer := range layersResponse.Layers {
    table.WriteRow(stackName, layer.Name, layer.Shortname, layer.LayerId)
  }

  table.PrintTable()
}
