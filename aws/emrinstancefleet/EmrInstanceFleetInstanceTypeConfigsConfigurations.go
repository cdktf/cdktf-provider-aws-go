package emrinstancefleet


type EmrInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emr_instance_fleet#classification EmrInstanceFleet#classification}.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emr_instance_fleet#properties EmrInstanceFleet#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

