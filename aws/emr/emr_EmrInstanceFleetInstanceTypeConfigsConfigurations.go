package emr


type EmrInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet#classification EmrInstanceFleet#classification}.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet#properties EmrInstanceFleet#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

