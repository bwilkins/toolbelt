package stacks

import (
  "log"
  "strconv"

  "github.com/bwilkins/toolbelt/command"
  "github.com/bwilkins/toolbelt/encoding/term_table"
  "github.com/bwilkins/aws/opsworks"
)

var StacksCmd = &command.Command{
  UsageLine: "stacks",
  ShortDesc: "connects to a remote server as specified on command-line",
  LongDesc:  ``,
}

func init() {
  StacksCmd.Run = runStacks

  StacksCmd.Flags.Usage = func() { StacksCmd.Usage() }
}

func runStacks(cmd *command.Command, args []string) {
  stacksResponse, err := opsworks.DescribeStacks(opsworks.DescribeStacksRequest{})
  if err != nil {
    log.Fatal(err.Error())
  }

  table := term_table.NewTermTable("Name", "Region", "# Layers", "# Instances", "ID")
  for _, stack := range stacksResponse.Stacks {
    var layersCount, instanceCount int
    layersResponse, err := opsworks.DescribeLayers(opsworks.DescribeLayersRequest{ StackId: stack.StackId })
    if err != nil {
      layersCount = 0
    } else {
      layersCount = len(layersResponse.Layers)
    }
    instancesResponse, err := opsworks.DescribeInstances(opsworks.DescribeInstancesRequest{ StackId: stack.StackId })
    if err != nil {
      instanceCount = 0
    } else {
      instanceCount = len(instancesResponse.Instances)
    }

    layers := strconv.FormatInt(int64(layersCount), 10)
    instances := strconv.FormatInt(int64(instanceCount), 10)

    table.WriteRow(stack.Name, stack.Region, layers, instances, stack.StackId)
  }

  table.PrintTable()
}
