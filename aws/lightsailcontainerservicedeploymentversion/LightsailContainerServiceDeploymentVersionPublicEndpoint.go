// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainerservicedeploymentversion


type LightsailContainerServiceDeploymentVersionPublicEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/lightsail_container_service_deployment_version#container_name LightsailContainerServiceDeploymentVersion#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/lightsail_container_service_deployment_version#container_port LightsailContainerServiceDeploymentVersion#container_port}.
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/lightsail_container_service_deployment_version#health_check LightsailContainerServiceDeploymentVersion#health_check}
	HealthCheck *LightsailContainerServiceDeploymentVersionPublicEndpointHealthCheck `field:"required" json:"healthCheck" yaml:"healthCheck"`
}

