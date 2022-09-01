package ec2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS EC2.
type EipAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#allocation_id EipAssociation#allocation_id}.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#allow_reassociation EipAssociation#allow_reassociation}.
	AllowReassociation interface{} `field:"optional" json:"allowReassociation" yaml:"allowReassociation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#id EipAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#instance_id EipAssociation#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#network_interface_id EipAssociation#network_interface_id}.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#private_ip_address EipAssociation#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eip_association#public_ip EipAssociation#public_ip}.
	PublicIp *string `field:"optional" json:"publicIp" yaml:"publicIp"`
}

