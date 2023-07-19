package appconfigextension


type AppconfigExtensionParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/appconfig_extension#name AppconfigExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/appconfig_extension#description AppconfigExtension#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/appconfig_extension#required AppconfigExtension#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

