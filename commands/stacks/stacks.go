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
  tableRowCount := len(stacksResponse.Stacks)

  stacksDetailsChan := make(chan []string, tableRowCount)

  for _, stack := range stacksResponse.Stacks {
    go fetchStackDetails(stack, stacksDetailsChan)
  }

  var row []string
  for i := 0; i < tableRowCount; i+=1 {
    row = <-stacksDetailsChan
    table.WriteRow(row...)
  }

  table.PrintTable()
}


func fetchStackDetails(stack opsworks.Stack, c chan []string) {
  var layerCountString, instanceCountString string
  layerCountChan := make(chan int64, 1)
  instanceCountChan := make(chan int64, 1)

  fetchLayerCount(stack.StackId, layerCountChan)
  fetchInstanceCount(stack.StackId, instanceCountChan)

  layerCountString = strconv.FormatInt(<-layerCountChan, 10)
  instanceCountString = strconv.FormatInt(<-instanceCountChan, 10)

  c <- []string{stack.Name, stack.Region, layerCountString, instanceCountString, stack.StackId}
}

func fetchLayerCount(stackId string, ret chan int64) {
  layersResponse, err := opsworks.DescribeLayers(opsworks.DescribeLayersRequest{ StackId: stackId })
  if err == nil {
    ret<- int64(len(layersResponse.Layers))
  } else {
    ret<- int64(0)
  }
}

func fetchInstanceCount(stackId string, ret chan int64) {
  instanceResponse, err := opsworks.DescribeInstances(opsworks.DescribeInstancesRequest{ StackId: stackId })
  if err == nil {
    ret<- int64(len(instanceResponse.Instances))
  } else {
    ret<- int64(0)
  }
}
