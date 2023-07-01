package dataawsrdsclusters


type DataAwsRdsClustersFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/rds_clusters#name DataAwsRdsClusters#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/rds_clusters#values DataAwsRdsClusters#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

