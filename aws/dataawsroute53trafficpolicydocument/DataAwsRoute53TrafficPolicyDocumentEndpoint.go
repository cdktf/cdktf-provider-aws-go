package dataawsroute53trafficpolicydocument


type DataAwsRoute53TrafficPolicyDocumentEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/route53_traffic_policy_document#id DataAwsRoute53TrafficPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/route53_traffic_policy_document#region DataAwsRoute53TrafficPolicyDocument#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/route53_traffic_policy_document#type DataAwsRoute53TrafficPolicyDocument#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/data-sources/route53_traffic_policy_document#value DataAwsRoute53TrafficPolicyDocument#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

