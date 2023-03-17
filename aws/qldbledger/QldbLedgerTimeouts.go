package qldbledger


type QldbLedgerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/qldb_ledger#create QldbLedger#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/qldb_ledger#delete QldbLedger#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

