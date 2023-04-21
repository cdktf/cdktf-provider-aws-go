package ivschatroom


type IvschatRoomMessageReviewHandler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ivschat_room#fallback_result IvschatRoom#fallback_result}.
	FallbackResult *string `field:"optional" json:"fallbackResult" yaml:"fallbackResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/ivschat_room#uri IvschatRoom#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

