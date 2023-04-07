package emrserverlessapplication


type EmrserverlessApplicationImageConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#image_uri EmrserverlessApplication#image_uri}.
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
}

