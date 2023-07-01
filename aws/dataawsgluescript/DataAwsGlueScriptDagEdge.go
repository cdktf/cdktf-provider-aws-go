package dataawsgluescript


type DataAwsGlueScriptDagEdge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/glue_script#source DataAwsGlueScript#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/glue_script#target DataAwsGlueScript#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.2/docs/data-sources/glue_script#target_parameter DataAwsGlueScript#target_parameter}.
	TargetParameter *string `field:"optional" json:"targetParameter" yaml:"targetParameter"`
}

