package emr


type EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#classification EmrCluster#classification}.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#properties EmrCluster#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

