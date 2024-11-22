// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncresolver


type AppsyncResolverPipelineConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.77.0/docs/resources/appsync_resolver#functions AppsyncResolver#functions}.
	Functions *[]*string `field:"optional" json:"functions" yaml:"functions"`
}

