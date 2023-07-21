package dataawsredshiftorderablecluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsRedshiftOrderableClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/redshift_orderable_cluster#cluster_type DataAwsRedshiftOrderableCluster#cluster_type}.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/redshift_orderable_cluster#cluster_version DataAwsRedshiftOrderableCluster#cluster_version}.
	ClusterVersion *string `field:"optional" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/redshift_orderable_cluster#id DataAwsRedshiftOrderableCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/redshift_orderable_cluster#node_type DataAwsRedshiftOrderableCluster#node_type}.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/data-sources/redshift_orderable_cluster#preferred_node_types DataAwsRedshiftOrderableCluster#preferred_node_types}.
	PreferredNodeTypes *[]*string `field:"optional" json:"preferredNodeTypes" yaml:"preferredNodeTypes"`
}

