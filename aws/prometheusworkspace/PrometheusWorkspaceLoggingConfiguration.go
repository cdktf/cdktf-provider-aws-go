// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusworkspace


type PrometheusWorkspaceLoggingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/prometheus_workspace#log_group_arn PrometheusWorkspace#log_group_arn}.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

