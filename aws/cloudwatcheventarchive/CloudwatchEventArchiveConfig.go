package cloudwatcheventarchive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchEventArchiveConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_archive#event_source_arn CloudwatchEventArchive#event_source_arn}.
	EventSourceArn *string `field:"required" json:"eventSourceArn" yaml:"eventSourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_archive#name CloudwatchEventArchive#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_archive#description CloudwatchEventArchive#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_archive#event_pattern CloudwatchEventArchive#event_pattern}.
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_archive#id CloudwatchEventArchive#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/cloudwatch_event_archive#retention_days CloudwatchEventArchive#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
}

