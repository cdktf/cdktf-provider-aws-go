package dboptiongroup


type DbOptionGroupOptionOptionSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/db_option_group#name DbOptionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/resources/db_option_group#value DbOptionGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

