package emrinstancefleet


type EmrInstanceFleetLaunchSpecifications struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emr_instance_fleet#on_demand_specification EmrInstanceFleet#on_demand_specification}
	OnDemandSpecification interface{} `field:"optional" json:"onDemandSpecification" yaml:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/emr_instance_fleet#spot_specification EmrInstanceFleet#spot_specification}
	SpotSpecification interface{} `field:"optional" json:"spotSpecification" yaml:"spotSpecification"`
}

