package dataawsamiids

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsAmiIdsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#owners DataAwsAmiIds#owners}.
	Owners *[]*string `field:"required" json:"owners" yaml:"owners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#executable_users DataAwsAmiIds#executable_users}.
	ExecutableUsers *[]*string `field:"optional" json:"executableUsers" yaml:"executableUsers"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#filter DataAwsAmiIds#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#id DataAwsAmiIds#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#include_deprecated DataAwsAmiIds#include_deprecated}.
	IncludeDeprecated interface{} `field:"optional" json:"includeDeprecated" yaml:"includeDeprecated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#name_regex DataAwsAmiIds#name_regex}.
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#sort_ascending DataAwsAmiIds#sort_ascending}.
	SortAscending interface{} `field:"optional" json:"sortAscending" yaml:"sortAscending"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/data-sources/ami_ids#timeouts DataAwsAmiIds#timeouts}
	Timeouts *DataAwsAmiIdsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

