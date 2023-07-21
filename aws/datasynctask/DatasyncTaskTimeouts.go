package datasynctask


type DatasyncTaskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/datasync_task#create DatasyncTask#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

