package dataawsrdsclusters


type DataAwsRdsClustersFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_clusters#name DataAwsRdsClusters#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/rds_clusters#values DataAwsRdsClusters#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

