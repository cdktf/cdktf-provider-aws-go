package redshiftserverlessworkgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftserverlessWorkgroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#namespace_name RedshiftserverlessWorkgroup#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#workgroup_name RedshiftserverlessWorkgroup#workgroup_name}.
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#base_capacity RedshiftserverlessWorkgroup#base_capacity}.
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// config_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#config_parameter RedshiftserverlessWorkgroup#config_parameter}
	ConfigParameter interface{} `field:"optional" json:"configParameter" yaml:"configParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#enhanced_vpc_routing RedshiftserverlessWorkgroup#enhanced_vpc_routing}.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#id RedshiftserverlessWorkgroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#publicly_accessible RedshiftserverlessWorkgroup#publicly_accessible}.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#security_group_ids RedshiftserverlessWorkgroup#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#subnet_ids RedshiftserverlessWorkgroup#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#tags RedshiftserverlessWorkgroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#tags_all RedshiftserverlessWorkgroup#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/redshiftserverless_workgroup#timeouts RedshiftserverlessWorkgroup#timeouts}
	Timeouts *RedshiftserverlessWorkgroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

