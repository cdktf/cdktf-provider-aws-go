// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type LocationTrackerAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/location_tracker_association#create LocationTrackerAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/location_tracker_association#delete LocationTrackerAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

