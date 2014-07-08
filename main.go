package main

import (
  "github.com/bwilkins/toolbelt-go/commands/ssh"
  "os"
  //"github.com/bwilkins/aws/opsworks"
  //"fmt"
)


func main() {
  ssh.SshCmd.Run(ssh.SshCmd, os.Args[1:])

  //request := opsworks.DescribeInstancesRequest{StackId: "ce22e71b-9b8e-4356-ab87-efcf02847f19"}

  //foo, _ := opsworks.DescribeInstances(request)
  //fmt.Printf("%s", foo.Instances[0].Hostname)
}
