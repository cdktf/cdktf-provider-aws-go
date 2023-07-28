package dynamodbtable


type DynamodbTableAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/dynamodb_table#name DynamodbTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/dynamodb_table#type DynamodbTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

