// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type FisExperimentTemplateAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#action_id FisExperimentTemplate#action_id}.
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#name FisExperimentTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#description FisExperimentTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#parameter FisExperimentTemplate#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#start_after FisExperimentTemplate#start_after}.
	StartAfter *[]*string `field:"optional" json:"startAfter" yaml:"startAfter"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fis_experiment_template#target FisExperimentTemplate#target}
	Target *FisExperimentTemplateActionTarget `field:"optional" json:"target" yaml:"target"`
}

