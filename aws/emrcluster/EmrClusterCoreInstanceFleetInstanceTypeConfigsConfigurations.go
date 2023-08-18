package emrcluster


type EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/emr_cluster#classification EmrCluster#classification}.
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/emr_cluster#properties EmrCluster#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

