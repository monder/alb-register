package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elbv2"
)

type targetConfig struct {
	ARN  *string
	Port *int64
}

func main() {

	targets := make([]targetConfig, 0)
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--") {
			argValue := strings.SplitN(arg[2:], "=", 2)
			if argValue[0] == "port" {
				port, err := strconv.ParseInt(argValue[1], 10, 64)
				if err != nil {
					log.Fatal(fmt.Errorf("Invalid port number '%s'\n", argValue[1]))
				}
				if len(targets) == 0 {
					log.Fatal(fmt.Errorf("port number should follow ARN\n"))
				}
				targets[len(targets)-1].Port = aws.Int64(port)

			} else {
				log.Fatal(fmt.Errorf("Unknown option '%s'\n", arg))
			}
		} else {
			targets = append(targets, targetConfig{ARN: aws.String(arg)})
		}
	}

	sess, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	meta := ec2metadata.New(sess)

	instanceId, err := meta.GetMetadata("instance-id")
	if err != nil {
		log.Fatal(err)
	}

	region, err := meta.Region()
	if err != nil {
		log.Fatal(err)
	}

	for _, target := range targets {
		_, err = elbv2.New(sess, &aws.Config{Region: aws.String(region)}).RegisterTargets(&elbv2.RegisterTargetsInput{
			TargetGroupArn: target.ARN,
			Targets: []*elbv2.TargetDescription{
				{Id: aws.String(instanceId), Port: target.Port},
			},
		})

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Instance '%s' registred in target group '%s'", instanceId, *target.ARN)
	}

	os.Exit(0)
}
