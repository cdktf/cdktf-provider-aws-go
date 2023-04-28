package redshiftserverlessusagelimit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftserverlessUsageLimitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshiftserverless_usage_limit#amount RedshiftserverlessUsageLimit#amount}.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshiftserverless_usage_limit#resource_arn RedshiftserverlessUsageLimit#resource_arn}.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshiftserverless_usage_limit#usage_type RedshiftserverlessUsageLimit#usage_type}.
	UsageType *string `field:"required" json:"usageType" yaml:"usageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshiftserverless_usage_limit#breach_action RedshiftserverlessUsageLimit#breach_action}.
	BreachAction *string `field:"optional" json:"breachAction" yaml:"breachAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshiftserverless_usage_limit#id RedshiftserverlessUsageLimit#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/redshiftserverless_usage_limit#period RedshiftserverlessUsageLimit#period}.
	Period *string `field:"optional" json:"period" yaml:"period"`
}

