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

	//scopeSelectors := d.Get("scope_selectors").(interface{})
	body.ScopeSelectors = d.Get("scope_selectors").(models.ImmutableTagRulesBodyScopeSelectors)
	//body.ScopeSelectors = expandScopeSelectors(scopeSelectors)

	body.TagSelectors = d.Get("tag_selectors").(models.ImmutableTagRulesBodyTagSelectors)
	//body.TagSelectors = expandTagSelectors(tagSelectors)

	return body
}

/*
func expandScopeSelectors(ss interface{}) models.ImmutableTagRulesBodyScopeSelectors {
	selectors := models.ImmutableTagRulesBodyScopeSelectors{}
	scopeSelectors := ss.(map[string][]models.ImmutableTagRulesBodyScopeSelectorRepository)

	for k, v := range scopeSelectors {
		if k == "repository" {
			repositories := []models.ImmutableTagRulesBodyScopeSelectorRepository{}

			for _, i := range v {
				item, _ := i.(models.ImmutableTagRulesBodyScopeSelectorRepository)
				repositories = append(repositories, models.ImmutableTagRulesBodyScopeSelectorRepository{
					Decoration: item.Decoration.(string),
					Pattern:    item.Pattern.(string),
					Kind:       item.Kind.(string),
				})
			}

			selectors.Repository = repositories
		}
	}

	return selectors
}

func expandTagSelectors(tagSelectors []interface{}) models.ImmutableTagRulesBodyTagSelectors {
	selectors := models.ImmutableTagRulesBodyTagSelectors{}

	for _, d := range tagSelectors {
		data := d.(models.ImmutableTagRulesBodyTagSelector)
		selectors = append(selectors, models.ImmutableTagRulesBodyTagSelector{
			Decoration: data.Decoration.(string),
			Pattern:    data.Pattern.(string),
			Kind:       data.Kind.(string),
			Extras:     data.Extras.(string),
		})
	}

	return selectors
}
*/
