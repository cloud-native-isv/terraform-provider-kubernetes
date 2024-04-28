// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"

	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"

	framework "github.com/hashicorp/terraform-provider-kubernetes/internal/framework/provider"
	"github.com/hashicorp/terraform-provider-kubernetes/kubernetes"
	manifest "github.com/hashicorp/terraform-provider-kubernetes/manifest/provider"
)

const (
	providerName = "registry.terraform.io/hashicorp/kubernetes"

	Version = "dev"
)

// Generate docs for website
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	if os.Getenv("TF_X_KUBERNETES_CODEGEN_PLUGIN6") == "1" {
		plugin6main()
		return
	}

	debugFlag := flag.Bool("debug", false, "Start provider in stand-alone debug mode.")
	flag.Parse()

	providers := []func() tfprotov5.ProviderServer{
		kubernetes.Provider().GRPCProvider,
		manifest.Provider(),
		providerserver.NewProtocol5(framework.New(Version)),
	}

	ctx := context.Background()
	muxer, err := tf5muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	opts := []tf5server.ServeOpt{}
	if *debugFlag {
		reattachConfigCh := make(chan *plugin.ReattachConfig)
		go func() {
			reattachConfig, err := waitForReattachConfig(reattachConfigCh)
			if err != nil {
				fmt.Printf("Error getting reattach config: %s\n", err)
				return
			}
			printReattachConfig(reattachConfig)
		}()
		opts = append(opts, tf5server.WithDebug(ctx, reattachConfigCh, nil))
	}

	tf5server.Serve(providerName, muxer.ProviderServer, opts...)
}

func plugin6main() {
	debugFlag := flag.Bool("debug", false, "Start provider in stand-alone debug mode.")
	flag.Parse()

	upgradedSdkServer, _ := tf5to6server.UpgradeServer(
		context.Background(),
		kubernetes.Provider().GRPCProvider,
	)

	upgradedManifestServer, _ := tf5to6server.UpgradeServer(
		context.Background(),
		manifest.Provider(),
	)

	providers := []func() tfprotov6.ProviderServer{
		func() tfprotov6.ProviderServer { return upgradedSdkServer },
		func() tfprotov6.ProviderServer { return upgradedManifestServer },
		providerserver.NewProtocol6(framework.New(Version)),
	}

	ctx := context.Background()
	muxer, err := tf6muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	opts := []tf6server.ServeOpt{}
	if *debugFlag {
		reattachConfigCh := make(chan *plugin.ReattachConfig)
		go func() {
			reattachConfig, err := waitForReattachConfig(reattachConfigCh)
			if err != nil {
				fmt.Printf("Error getting reattach config: %s\n", err)
				return
			}
			printReattachConfig(reattachConfig)
		}()
		opts = append(opts, tf6server.WithDebug(ctx, reattachConfigCh, nil))
	}

	tf6server.Serve(providerName, muxer.ProviderServer, opts...)
}

// convertReattachConfig converts plugin.ReattachConfig to tfexec.ReattachConfig
func convertReattachConfig(reattachConfig *plugin.ReattachConfig) tfexec.ReattachConfig {
	return tfexec.ReattachConfig{
		Protocol:        string(reattachConfig.Protocol),
		ProtocolVersion: reattachConfig.ProtocolVersion,
		Pid:             reattachConfig.Pid,
		Test:            true,
		Addr: tfexec.ReattachConfigAddr{
			Network: reattachConfig.Addr.Network(),
			String:  reattachConfig.Addr.String(),
		},
	}
}

// printReattachConfig prints the line the user needs to copy and paste
// to set the TF_REATTACH_PROVIDERS variable
func printReattachConfig(config *plugin.ReattachConfig) {
	reattachStr, err := json.Marshal(map[string]tfexec.ReattachConfig{
		"kubernetes": convertReattachConfig(config),
	})
	if err != nil {
		fmt.Printf("Error building reattach string: %s", err)
		return
	}
	fmt.Printf("# Provider server started\nexport TF_REATTACH_PROVIDERS='%s'\n", string(reattachStr))
}

// waitForReattachConfig blocks until a ReattachConfig is recieved on the
// supplied channel or times out after 2 seconds.
func waitForReattachConfig(ch chan *plugin.ReattachConfig) (*plugin.ReattachConfig, error) {
	select {
	case config := <-ch:
		return config, nil
	case <-time.After(2 * time.Second):
		return nil, fmt.Errorf("timeout while waiting for reattach configuration")
	}
}
