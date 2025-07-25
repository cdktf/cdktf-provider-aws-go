// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusqueryloggingconfiguration


type PrometheusQueryLoggingConfigurationDestinationFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/prometheus_query_logging_configuration#qsp_threshold PrometheusQueryLoggingConfiguration#qsp_threshold}.
	QspThreshold *float64 `field:"required" json:"qspThreshold" yaml:"qspThreshold"`
}

