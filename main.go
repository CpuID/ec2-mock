package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

func parseTagEnvVars() (InstanceTags, error) {
	result := make(InstanceTags)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if len(pair[0]) > 5 && pair[0][0:5] == "TAG__" {
			split_key := strings.Split(pair[0], "__")
			if len(split_key) != 3 {
				return InstanceTags{}, fmt.Errorf("Invalid format for '%s' environment variable key. Must be in the format of 'TAG__instanceid__tagname'", split_key)
			}
			use_instance_id := strings.Replace(split_key[1], "_", "-", 1)
			result[use_instance_id] = append(result[use_instance_id], InstanceTag{
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
