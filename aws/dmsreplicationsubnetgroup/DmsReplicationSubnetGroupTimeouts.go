package dmsreplicationsubnetgroup


type DmsReplicationSubnetGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/dms_replication_subnet_group#create DmsReplicationSubnetGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/dms_replication_subnet_group#delete DmsReplicationSubnetGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/dms_replication_subnet_group#update DmsReplicationSubnetGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

