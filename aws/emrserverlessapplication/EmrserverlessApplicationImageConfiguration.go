package emrserverlessapplication


type EmrserverlessApplicationImageConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/emrserverless_application#image_uri EmrserverlessApplication#image_uri}.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

