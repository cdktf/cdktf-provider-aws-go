package controltowercontrol


type ControltowerControlTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/controltower_control#create ControltowerControl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/controltower_control#delete ControltowerControl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

