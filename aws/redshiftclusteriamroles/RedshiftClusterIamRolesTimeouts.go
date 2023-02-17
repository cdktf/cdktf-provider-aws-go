package redshiftclusteriamroles


type RedshiftClusterIamRolesTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshift_cluster_iam_roles#create RedshiftClusterIamRoles#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshift_cluster_iam_roles#delete RedshiftClusterIamRoles#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/redshift_cluster_iam_roles#update RedshiftClusterIamRoles#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

