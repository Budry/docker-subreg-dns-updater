package main

import (
	"context"
	"fmt"
	"os"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/Budry/subreg-dns-updater-cli/ip"
	"github.com/Budry/subreg-dns-updater-cli/subreg"
	"github.com/Budry/subreg-dns-updater-cli/utils"
)

func main() {

	service := subreg.NewSubregCz("", false, &subreg.BasicAuth{})
	dnsManager := subreg.DNSManager{
		Ip:     ip.GetPublicIp(),
		Client: service,
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
