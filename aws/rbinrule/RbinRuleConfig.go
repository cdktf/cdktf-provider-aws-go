package rbinrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RbinRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#resource_type RbinRule#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// retention_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#retention_period RbinRule#retention_period}
	RetentionPeriod *RbinRuleRetentionPeriod `field:"required" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#description RbinRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// lock_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#lock_configuration RbinRule#lock_configuration}
	LockConfiguration *RbinRuleLockConfiguration `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#resource_tags RbinRule#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#tags RbinRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#tags_all RbinRule#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/rbin_rule#timeouts RbinRule#timeouts}
	Timeouts *RbinRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

