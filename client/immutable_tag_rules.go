package client

import (
	"log"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-harbor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ImmutableTagRulesBody(d *schema.ResourceData) models.ImmutableTagRulesBody {
	body := models.ImmutableTagRulesBody{
		ID:       d.Get("id").(int)
		Priority: d.Get("priority").(int)
		Disabled: d.Get("disabled").(bool)
		Template: d.Get("template").(string)
		Action:   d.Get("action").(string)
	}
}
