package ec2


type SpotInstanceRequestLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_instance_request#id SpotInstanceRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_instance_request#name SpotInstanceRequest#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/spot_instance_request#version SpotInstanceRequest#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

