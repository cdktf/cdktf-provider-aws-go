package apigatewayusageplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayUsagePlanConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#name ApiGatewayUsagePlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// api_stages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#api_stages ApiGatewayUsagePlan#api_stages}
	ApiStages interface{} `field:"optional" json:"apiStages" yaml:"apiStages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#description ApiGatewayUsagePlan#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#id ApiGatewayUsagePlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#product_code ApiGatewayUsagePlan#product_code}.
	ProductCode *string `field:"optional" json:"productCode" yaml:"productCode"`
	// quota_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#quota_settings ApiGatewayUsagePlan#quota_settings}
	QuotaSettings *ApiGatewayUsagePlanQuotaSettings `field:"optional" json:"quotaSettings" yaml:"quotaSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#tags ApiGatewayUsagePlan#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#tags_all ApiGatewayUsagePlan#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// throttle_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/api_gateway_usage_plan#throttle_settings ApiGatewayUsagePlan#throttle_settings}
	ThrottleSettings *ApiGatewayUsagePlanThrottleSettings `field:"optional" json:"throttleSettings" yaml:"throttleSettings"`
}

