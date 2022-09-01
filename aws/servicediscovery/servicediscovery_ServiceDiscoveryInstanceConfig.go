package servicediscovery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS Service Discovery.
type ServiceDiscoveryInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_instance#attributes ServiceDiscoveryInstance#attributes}.
	Attributes *map[string]*string `field:"required" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_instance#instance_id ServiceDiscoveryInstance#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_instance#service_id ServiceDiscoveryInstance#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/service_discovery_instance#id ServiceDiscoveryInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

