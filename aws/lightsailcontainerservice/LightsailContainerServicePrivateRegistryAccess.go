package lightsailcontainerservice


type LightsailContainerServicePrivateRegistryAccess struct {
	// ecr_image_puller_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/lightsail_container_service#ecr_image_puller_role LightsailContainerService#ecr_image_puller_role}
	EcrImagePullerRole *LightsailContainerServicePrivateRegistryAccessEcrImagePullerRole `field:"optional" json:"ecrImagePullerRole" yaml:"ecrImagePullerRole"`
}

