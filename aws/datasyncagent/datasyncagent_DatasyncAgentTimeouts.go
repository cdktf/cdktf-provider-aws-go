package datasyncagent


type DatasyncAgentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_agent#create DatasyncAgent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

