// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type KeyspacesTablePointInTimeRecovery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/keyspaces_table#status KeyspacesTable#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

