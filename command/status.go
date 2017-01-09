package command

import (
	"strings"

	"github.com/DaveBlooman/fasten/output"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
)

type StatusCommand struct {
	Meta
}

func (c *StatusCommand) Run(args []string) int {

	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		output.Error("failed to create AWS session")
		return 1
	}

	svc := lightsail.New(sess)

	serverParams := &lightsail.GetInstancesInput{}
	serverResp, err := svc.GetInstances(serverParams)
	if err != nil {
		output.Error(err.Error())
		return 1
	}

	instance := serverResp.Instances[0]

	flags := map[string]string{"Name": *instance.Name,
		"IP Address": *instance.PublicIpAddress,
		"State":      *instance.State.Name,
		"Keypair":    *instance.SshKeyName,
		"Username":   *instance.Username,
		"Location":   *instance.Location.AvailabilityZone,
	}

	output.Status("Service Status", flags)

	return 0

}

func (c *StatusCommand) Synopsis() string {
	return ""
}

func (c *StatusCommand) Help() string {
	helpText := `Shows Information about deployment

`
	return strings.TrimSpace(helpText)
}
