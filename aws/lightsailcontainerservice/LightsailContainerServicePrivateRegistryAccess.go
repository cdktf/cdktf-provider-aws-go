// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainerservice


type LightsailContainerServicePrivateRegistryAccess struct {
	// ecr_image_puller_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/lightsail_container_service#ecr_image_puller_role LightsailContainerService#ecr_image_puller_role}
	EcrImagePullerRole *LightsailContainerServicePrivateRegistryAccessEcrImagePullerRole `field:"optional" json:"ecrImagePullerRole" yaml:"ecrImagePullerRole"`
}

