package mediaconvertqueue


type MediaConvertQueueReservationPlanSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/media_convert_queue#commitment MediaConvertQueue#commitment}.
	Commitment *string `field:"required" json:"commitment" yaml:"commitment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/media_convert_queue#renewal_type MediaConvertQueue#renewal_type}.
	RenewalType *string `field:"required" json:"renewalType" yaml:"renewalType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/media_convert_queue#reserved_slots MediaConvertQueue#reserved_slots}.
	ReservedSlots *float64 `field:"required" json:"reservedSlots" yaml:"reservedSlots"`
}

