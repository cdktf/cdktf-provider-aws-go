// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type FisExperimentTemplateTargetFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#path FisExperimentTemplate#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#values FisExperimentTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

