// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codeartifactrepository


type CodeartifactRepositoryExternalConnections struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/resources/codeartifact_repository#external_connection_name CodeartifactRepository#external_connection_name}.
	ExternalConnectionName *string `field:"required" json:"externalConnectionName" yaml:"externalConnectionName"`
}

