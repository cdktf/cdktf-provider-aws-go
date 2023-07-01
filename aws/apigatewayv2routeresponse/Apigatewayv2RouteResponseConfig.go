package apigatewayv2routeresponse

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Apigatewayv2RouteResponseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/apigatewayv2_route_response#api_id Apigatewayv2RouteResponse#api_id}.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/apigatewayv2_route_response#route_id Apigatewayv2RouteResponse#route_id}.
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/apigatewayv2_route_response#route_response_key Apigatewayv2RouteResponse#route_response_key}.
	RouteResponseKey *string `field:"required" json:"routeResponseKey" yaml:"routeResponseKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/apigatewayv2_route_response#id Apigatewayv2RouteResponse#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/apigatewayv2_route_response#model_selection_expression Apigatewayv2RouteResponse#model_selection_expression}.
	ModelSelectionExpression *string `field:"optional" json:"modelSelectionExpression" yaml:"modelSelectionExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/apigatewayv2_route_response#response_models Apigatewayv2RouteResponse#response_models}.
	ResponseModels *map[string]*string `field:"optional" json:"responseModels" yaml:"responseModels"`
}

