package elb


type ElbListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elb#instance_port Elb#instance_port}.
	InstancePort *float64 `field:"required" json:"instancePort" yaml:"instancePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elb#instance_protocol Elb#instance_protocol}.
	InstanceProtocol *string `field:"required" json:"instanceProtocol" yaml:"instanceProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elb#lb_port Elb#lb_port}.
	LbPort *float64 `field:"required" json:"lbPort" yaml:"lbPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elb#lb_protocol Elb#lb_protocol}.
	LbProtocol *string `field:"required" json:"lbProtocol" yaml:"lbProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/elb#ssl_certificate_id Elb#ssl_certificate_id}.
	SslCertificateId *string `field:"optional" json:"sslCertificateId" yaml:"sslCertificateId"`
}

