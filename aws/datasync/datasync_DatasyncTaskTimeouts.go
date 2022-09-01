package datasync


type DatasyncTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/datasync_task#create DatasyncTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

