package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

// Key = Instance ID
type InstanceTags map[string][]InstanceTag

type InstanceTag struct {
	Key   string
	Value string
}

func parseTagEnvVars() (InstanceTags, error) {
	result := make(InstanceTags)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0][0:5] == "TAG__" {
			split_key := strings.Split(pair[0], "__")
			if len(split_key) != 3 {
				return InstanceTags{}, fmt.Errorf("Invalid format for '%s' environment variable key. Must be in the format of 'TAG__instanceid__tagname'")
			}
			result[split_key[1]] = append(result[split_key[1]], InstanceTag{
				Key:   split_key[2],
				Value: pair[1],
			})
		}
	}
	return result, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "ec2-mock"
	app.Usage = "Mock (parts of) the Amazon EC2 API"
	app.Flags = []cli.Flag{
		cli.Uint64Flag{
			Name:   "port, p",
			EnvVar: "PORT",
			Value:  uint64(33333),
			Usage:  "Listen Port",
		},
	}
	app.Action = func(c *cli.Context) error {
		server := Server{}
		server.Port = c.Uint64("port")
		tags, err := parseTagEnvVars()
		if err != nil {
			return err
		}
		server.Tags = tags
		server.Start()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
