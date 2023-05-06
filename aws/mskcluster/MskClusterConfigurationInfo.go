package mskcluster


type MskClusterConfigurationInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/msk_cluster#arn MskCluster#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/msk_cluster#revision MskCluster#revision}.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
}

