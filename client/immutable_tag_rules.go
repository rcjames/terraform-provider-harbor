package client

import (
	"github.com/BESTSELLER/terraform-provider-harbor/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ImmutableTagRulesBody(d *schema.ResourceData) models.ImmutableTagRulesBody {
	body := models.ImmutableTagRulesBody{
		ID:       d.Get("id").(int),
		Priority: d.Get("priority").(int),
		Disabled: d.Get("disabled").(bool),
		Template: d.Get("template").(string),
		Action:   d.Get("action").(string),
	}

	scopeSelectors := d.Get("scope_selectors").(interface{})
	body.ScopeSelectors = expandScopeSelectors(scopeSelectors)

	tagSelectors := d.Get("tag_selectors").([]interface{})
	body.TagSelectors = expandTagSelectors(tagSelectors)

	return body
}

func expandScopeSelectors(scopeSelectors interface{}) models.ImmutableTagRulesBodyScopeSelectors {
	selectors := models.ImmutableTagRulesBodyScopeSelectors{}

	for k, v := range scopeSelectors {
		if k == "repository" {
			repository := models.ImmutableTagRulesBodyScopeSelectorRepository{
				Decoration: v["decoration"].(string),
				Pattern:    v["pattern"].(string),
				Kind:       v["kind"].(string),
			}
			selectors.Repository = repository
		}
	}

	return selectors
}

func expandTagSelectors(tagSelectors []interface) []models.ImmutableTagRulesBodyTagSelector {
	selectors := []models.ImmutableTagRulesBodyTagSelector{}

	for _, data := range tagSelectors {
		append(selectors, model.ImmutableTagRulesBodyTagSelector{
			Decoration: data["decoration"].(string),
			Pattern:    data["pattern"].(string),
			Kind:       data["kind"].(string),
		})
	}

	return selectors
}
