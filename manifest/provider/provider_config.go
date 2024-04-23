// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// GetProviderConfigSchema contains the definitions of all configuration attributes
func GetProviderConfigSchema() *tfprotov6.Schema {
	b := tfprotov6.SchemaBlock{

		Attributes: []*tfprotov6.SchemaAttribute{
			{
				Name:            "host",
				Type:            tftypes.String,
				Description:     "The hostname (in form of URI) of Kubernetes master.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "username",
				Type:            tftypes.String,
				Description:     "The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "password",
				Type:            tftypes.String,
				Description:     "The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "insecure",
				Type:            tftypes.Bool,
				Description:     "Whether server should be accessed without verifying the TLS certificate.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "tls_server_name",
				Type:            tftypes.String,
				Description:     "Server name passed to the server for SNI and is used in the client to check server certificates against.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "client_certificate",
				Type:            tftypes.String,
				Description:     "PEM-encoded client certificate for TLS authentication.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "client_key",
				Type:            tftypes.String,
				Description:     "PEM-encoded client certificate key for TLS authentication.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "cluster_ca_certificate",
				Type:            tftypes.String,
				Description:     "PEM-encoded root certificates bundle for TLS authentication.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_paths",
				Type:            tftypes.List{ElementType: tftypes.String},
				Description:     "A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_path",
				Type:            tftypes.String,
				Description:     "Path to the kube config file. Can be set with KUBE_CONFIG_PATH.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_context",
				Type:            tftypes.String,
				Description:     "",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_context_auth_info",
				Type:            tftypes.String,
				Description:     "",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "config_context_cluster",
				Type:            tftypes.String,
				Description:     "",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "token",
				Type:            tftypes.String,
				Description:     "Token to authenticate an service account",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "proxy_url",
				Type:            tftypes.String,
				Description:     "URL to the proxy to be used for all API requests",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "ignore_annotations",
				Type:            tftypes.List{ElementType: tftypes.String},
				Description:     "List of Kubernetes metadata annotations to ignore across all resources handled by this provider for situations where external systems are managing certain resource annotations. Each item is a regular expression.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
			{
				Name:            "ignore_labels",
				Type:            tftypes.List{ElementType: tftypes.String},
				Description:     "List of Kubernetes metadata labels to ignore across all resources handled by this provider for situations where external systems are managing certain resource labels. Each item is a regular expression.",
				Required:        false,
				Optional:        true,
				Computed:        false,
				Sensitive:       false,
				DescriptionKind: 0,
				Deprecated:      false,
			},
		},
		BlockTypes: []*tfprotov6.SchemaNestedBlock{
			{
				TypeName: "exec",
				Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
				MinItems: 0,
				MaxItems: 1,
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name:            "api_version",
							Type:            tftypes.String,
							Required:        true,
							Optional:        false,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
						{
							Name:            "command",
							Type:            tftypes.String,
							Required:        true,
							Optional:        false,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
						{
							Name:            "env",
							Type:            tftypes.Map{ElementType: tftypes.String},
							Required:        false,
							Optional:        true,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
						{
							Name:            "args",
							Type:            tftypes.List{ElementType: tftypes.String},
							Required:        false,
							Optional:        true,
							Computed:        false,
							Sensitive:       false,
							DescriptionKind: 0,
							Deprecated:      false,
						},
					},
				},
			},
			{
				TypeName: "experiments",
				Nesting:  tfprotov6.SchemaNestedBlockNestingModeList,
				MinItems: 0,
				MaxItems: 1,
				Block: &tfprotov6.SchemaBlock{
					Description: "Enable and disable experimental features.",
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name:            "manifest_resource",
							Type:            tftypes.Bool,
							Required:        false,
							Optional:        true,
							Computed:        false,
							Sensitive:       false,
							Description:     "Enable the `kubernetes_manifest` resource.",
							DescriptionKind: 0,
							Deprecated:      false,
						},
					},
				},
			},
		},
	}

	return &tfprotov6.Schema{
		Version: 0,
		Block:   &b,
	}
}
