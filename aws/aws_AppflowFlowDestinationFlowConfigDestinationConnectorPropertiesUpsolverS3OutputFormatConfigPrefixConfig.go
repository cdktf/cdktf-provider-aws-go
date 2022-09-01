// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowDestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#prefix_type AppflowFlow#prefix_type}.
	PrefixType *string `field:"required" json:"prefixType" yaml:"prefixType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#prefix_format AppflowFlow#prefix_format}.
	PrefixFormat *string `field:"optional" json:"prefixFormat" yaml:"prefixFormat"`
}

