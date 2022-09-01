package amplify


type AmplifyAppCustomRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app#source AmplifyApp#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app#target AmplifyApp#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app#condition AmplifyApp#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/amplify_app#status AmplifyApp#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

