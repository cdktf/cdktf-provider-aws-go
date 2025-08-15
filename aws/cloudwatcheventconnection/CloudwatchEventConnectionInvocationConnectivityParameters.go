// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventconnection


type CloudwatchEventConnectionInvocationConnectivityParameters struct {
	// resource_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/cloudwatch_event_connection#resource_parameters CloudwatchEventConnection#resource_parameters}
	ResourceParameters *CloudwatchEventConnectionInvocationConnectivityParametersResourceParameters `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

