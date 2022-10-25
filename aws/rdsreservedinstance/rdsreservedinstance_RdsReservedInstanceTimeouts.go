package rdsreservedinstance


type RdsReservedInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rds_reserved_instance#create RdsReservedInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rds_reserved_instance#delete RdsReservedInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/rds_reserved_instance#update RdsReservedInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

