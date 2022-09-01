package transfer


type TransferAccessPosixProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#gid TransferAccess#gid}.
	Gid *float64 `field:"required" json:"gid" yaml:"gid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#uid TransferAccess#uid}.
	Uid *float64 `field:"required" json:"uid" yaml:"uid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transfer_access#secondary_gids TransferAccess#secondary_gids}.
	SecondaryGids *[]*float64 `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
}

