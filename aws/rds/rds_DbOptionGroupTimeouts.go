package rds


type DbOptionGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/db_option_group#delete DbOptionGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

