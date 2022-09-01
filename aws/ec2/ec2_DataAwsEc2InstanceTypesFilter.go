package ec2


type DataAwsEc2InstanceTypesFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_instance_types#name DataAwsEc2InstanceTypes#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/ec2_instance_types#values DataAwsEc2InstanceTypes#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

