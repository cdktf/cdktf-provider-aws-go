// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}.
	Endpoints *map[string]*string `field:"required" json:"endpoints" yaml:"endpoints"`
}

