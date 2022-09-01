// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type FisExperimentTemplateStopCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#source FisExperimentTemplate#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#value FisExperimentTemplate#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

