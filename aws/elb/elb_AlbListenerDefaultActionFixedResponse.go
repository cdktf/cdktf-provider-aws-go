package elb


type AlbListenerDefaultActionFixedResponse struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#content_type AlbListener#content_type}.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#message_body AlbListener#message_body}.
	MessageBody *string `field:"optional" json:"messageBody" yaml:"messageBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/alb_listener#status_code AlbListener#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

