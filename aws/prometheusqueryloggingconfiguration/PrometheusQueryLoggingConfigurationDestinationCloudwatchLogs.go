// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusqueryloggingconfiguration


type PrometheusQueryLoggingConfigurationDestinationCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/prometheus_query_logging_configuration#log_group_arn PrometheusQueryLoggingConfiguration#log_group_arn}.
	LogGroupArn *string `field:"required" json:"logGroupArn" yaml:"logGroupArn"`
}

