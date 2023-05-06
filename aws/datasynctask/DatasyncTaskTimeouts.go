package datasynctask


type DatasyncTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/datasync_task#create DatasyncTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

