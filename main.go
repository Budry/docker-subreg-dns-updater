package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/Budry/docker-subreg-dns-updater/ip"
	"github.com/Budry/docker-subreg-dns-updater/subreg"
	"github.com/Budry/docker-subreg-dns-updater/utils"
)

func main() {

	service := subreg.NewSubregCz("", false, &subreg.BasicAuth{})
	publicIp, err := ip.GetPublicIp()
	if err != nil {
		log.Fatalln(err)
	}
	dnsManager := subreg.DNSManager{
		Ip:     publicIp,
		Client: service,
	}

	err = dnsManager.Update(os.Getenv("HOSTNAME"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	ctx := context.Background()

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	result, _ := cli.Events(ctx, types.EventsOptions{})

	for {
		message := <-result
		if message.Action == "create" {
			inspect, err := cli.ContainerInspect(ctx, message.ID)
			if err != nil {
				continue
			}
			envMap := utils.SplitKeyValueSlice(inspect.Config.Env)
			err = dnsManager.Update(envMap["VIRTUAL_HOST"])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
		}
	}
}
