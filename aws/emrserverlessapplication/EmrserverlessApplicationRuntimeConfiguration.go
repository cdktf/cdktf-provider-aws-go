// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationRuntimeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/emrserverless_application#classification EmrserverlessApplication#classification}.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/emrserverless_application#properties EmrserverlessApplication#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

