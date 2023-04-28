package opensearchdomain


type OpensearchDomainEbsOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opensearch_domain#ebs_enabled OpensearchDomain#ebs_enabled}.
	EbsEnabled interface{} `field:"required" json:"ebsEnabled" yaml:"ebsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opensearch_domain#iops OpensearchDomain#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opensearch_domain#throughput OpensearchDomain#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opensearch_domain#volume_size OpensearchDomain#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opensearch_domain#volume_type OpensearchDomain#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

