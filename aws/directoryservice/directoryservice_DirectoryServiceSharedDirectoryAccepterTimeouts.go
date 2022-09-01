package directoryservice


type DirectoryServiceSharedDirectoryAccepterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/directory_service_shared_directory_accepter#create DirectoryServiceSharedDirectoryAccepter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/directory_service_shared_directory_accepter#delete DirectoryServiceSharedDirectoryAccepter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

