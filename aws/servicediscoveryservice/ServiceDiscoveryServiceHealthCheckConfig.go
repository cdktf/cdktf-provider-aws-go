// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryservice


type ServiceDiscoveryServiceHealthCheckConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/service_discovery_service#failure_threshold ServiceDiscoveryService#failure_threshold}.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/service_discovery_service#resource_path ServiceDiscoveryService#resource_path}.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/service_discovery_service#type ServiceDiscoveryService#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

