package vpc

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS VPC.
type VpcEndpointConnectionAccepterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_connection_accepter#vpc_endpoint_id VpcEndpointConnectionAccepter#vpc_endpoint_id}.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_connection_accepter#vpc_endpoint_service_id VpcEndpointConnectionAccepter#vpc_endpoint_service_id}.
	VpcEndpointServiceId *string `field:"required" json:"vpcEndpointServiceId" yaml:"vpcEndpointServiceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/vpc_endpoint_connection_accepter#id VpcEndpointConnectionAccepter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
