package ssh

import (
  "github.com/bwilkins/toolbelt/command"
  "github.com/bwilkins/aws/opsworks"
  "log"
  "strings"
  "os"
  "os/exec"
)

var SshCmd = &command.Command{
  UsageLine: "ssh",
  ShortDesc: "connects to a remote server as specified on command-line",
  LongDesc:  ``,
}

var stackName string
var instanceName string

func runSsh(cmd *command.Command, args []string) {
  cmd.Flags.Parse(args)

  if stackName == "" {
    log.Fatal("Cannot SSH without a stack name")
    return
  }

  stacksResponse, err := opsworks.DescribeStacks(opsworks.DescribeStacksRequest{})
  if err != nil {
    log.Fatal(err.Error())
  }

  stack := selectStackFromResponse(stackName, stacksResponse)

  instancesResponse, err := opsworks.DescribeInstances(opsworks.DescribeInstancesRequest{StackId: stack.StackId })

  instance := selectInstanceFromResponse(instanceName, instancesResponse)

  shellCommand := exec.Command("ssh", "-l", "bugg", instance.PublicIp)

  shellCommand.Stdin = os.Stdin
  shellCommand.Stdout = os.Stdout
  shellCommand.Stderr = os.Stderr

  shellCommand.Run()
}

func init() {
  SshCmd.Run = runSsh

  SshCmd.Flags.Usage = func() { SshCmd.Usage() }

  SshCmd.Flags.StringVar(&stackName, "s", "", "which stack find an instance to ssh into")
  SshCmd.Flags.StringVar(&stackName, "stack", "", "which stack find an instance to ssh into")
  SshCmd.Flags.StringVar(&instanceName, "i", "", "which instance to ssh into (fuzzy match)")
  SshCmd.Flags.StringVar(&instanceName, "instance", "", "which instance to ssh into (fuzzy match)")
}

func selectStackFromResponse(stackName string, response *opsworks.DescribeStacksResponse) opsworks.Stack {
  var selectedStack opsworks.Stack
  for _, mStack := range response.Stacks {
    if mStack.Name == stackName {
      selectedStack = mStack
      break
    }
  }

  return selectedStack
}

func selectInstanceFromResponse(instanceName string, response *opsworks.DescribeInstancesResponse) opsworks.Instance {
  var selectedInstance opsworks.Instance

  if len(response.Instances) <= 0 {
    log.Fatal("No instances available for %s", stackName)
  }

  if instanceName == "" {
    return response.Instances[0]
  }

  for _, mInstance := range response.Instances {
    if strings.Contains(mInstance.InstanceId, instanceName) {
      selectedInstance = mInstance
      break
    }
  }

  return selectedInstance
}
