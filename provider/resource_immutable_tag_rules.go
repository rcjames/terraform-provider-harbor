package provider

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-harbor/client"
	"github.com/BESTSELLER/terraform-provider-harbor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImmutableTagRules() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			// To do - Scope selectors
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"template": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "immutable_template",
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "immutable",
			},
			// To do - Tag selectors
		},
	}
}

func resourceImmutableTagRulesCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceImmutableTagRulesRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceImmutableTagRulesUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceImmutableTagRulesDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
