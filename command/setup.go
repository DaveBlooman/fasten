package command

import (
	"io/ioutil"
	"strings"

	"github.com/DaveBlooman/fasten/output"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
)

type SetupCommand struct {
	Meta
}

// func (c *SetupCommand) Run(args []string) int {
// 	config := &aws.Config{
// 		Region: aws.String("us-east-1"),
// 	}
//
// 	sess, err := session.NewSession(config)
// 	if err != nil {
// 		output.Error("failed to create AWS session")
// 		return 1
// 	}
//
// 	svc := lightsail.New(sess)
//
// 	params := &lightsail.GetInstancesInput{}
// 	resp, err := svc.GetInstances(params)
//
// 	if err != nil {
// 		// Print the error, cast err to awserr.Error to get the Code and
// 		// Message from an error.
// 		fmt.Println(err.Error())
// 		return 1
// 	}
//
// 	fmt.Println(resp)
// 	return 0
// }

func (c *SetupCommand) Run(args []string) int {

	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		output.Error("failed to create AWS session")
		return 1
	}

	svc := lightsail.New(sess)

	var params *lightsail.DownloadDefaultKeyPairInput
	resp, err := svc.DownloadDefaultKeyPair(params)
	if err != nil {
		output.Error(err.Error())
		return 1
	}

	fasteKeyPair := []byte(*resp.PrivateKeyBase64)
	err = ioutil.WriteFile("fasten.pem", fasteKeyPair, 0644)
	if err != nil {
		output.Error(err.Error())
		return 1
	}

	output.Success("Downloaded Key Pair")

	instanceParams := &lightsail.CreateInstancesInput{
		AvailabilityZone: aws.String("us-east-1a"),
		BlueprintId:      aws.String("amazon_linux_2016_09_0"),
		BundleId:         aws.String("micro_1_0"),
		InstanceNames: []*string{
			aws.String("fasten-instance"),
		},
	}

	_, err = svc.CreateInstances(instanceParams)
	if err != nil {
		output.Error(err.Error())
		return 1
	}

	startParams := &lightsail.StartInstanceInput{
		InstanceName: aws.String("fasten-instance"),
	}

	_, startErr := svc.StartInstance(startParams)
	if startErr != nil {
		output.Error(startErr.Error())
		return 1
	}

	serverParams := &lightsail.GetInstancesInput{}
	serverResp, err := svc.GetInstances(serverParams)
	if err != nil {
		output.Error(err.Error())
		return 1
	}

	output.Success("Instance launched. IP address: " + *serverResp.Instances[0].PublicIpAddress)

	return 0

}

func (c *SetupCommand) Synopsis() string {
	return ""
}

func (c *SetupCommand) Help() string {
	helpText := `Shows Information about deployment

`
	return strings.TrimSpace(helpText)
}
