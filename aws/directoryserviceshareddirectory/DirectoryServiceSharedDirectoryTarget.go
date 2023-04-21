package directoryserviceshareddirectory


type DirectoryServiceSharedDirectoryTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/directory_service_shared_directory#id DirectoryServiceSharedDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/directory_service_shared_directory#type DirectoryServiceSharedDirectory#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

