package ec2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// AWS EC2.
type Ec2TrafficMirrorFilterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_traffic_mirror_filter#description Ec2TrafficMirrorFilter#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_traffic_mirror_filter#id Ec2TrafficMirrorFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_traffic_mirror_filter#network_services Ec2TrafficMirrorFilter#network_services}.
	NetworkServices *[]*string `field:"optional" json:"networkServices" yaml:"networkServices"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_traffic_mirror_filter#tags Ec2TrafficMirrorFilter#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ec2_traffic_mirror_filter#tags_all Ec2TrafficMirrorFilter#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

