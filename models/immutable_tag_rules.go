package models

type ImmutableTagRulesBody struct {
	ID             int                                 `json:"id,omitempty"`
	ProjectId      int                                 `json:"project_id,omitempty"`
	Priority       int                                 `json:"priority,omitempty"`
	ScopeSelectors ImmutableTagRulesBodyScopeSelectors `json:"scope_selectors,omitempty"`
	Disabled       bool                                `json:"disabled,omitempty"`
	Template       string                              `json:"template,omitempty"`
	Action         string                              `json:"action,omitempty"`
	TagSelectors   ImmutableTagRulesBodyTagSelectors   `json:"tag_selectors,omitempty"`
}

type ImmutableTagRulesBodyScopeSelectors struct {
	Repository []ImmutableTagRulesBodyScopeSelectorRepository `json:"repository,omitempty"`
}

type ImmutableTagRulesBodyScopeSelectorRepository struct {
	Decoration string `json:"decoration,omitempty"`
	Pattern    string `json:"pattern,omitempty"`
	Kind       string `json:"kind,omitempty"`
}

type ImmutableTagRulesBodyTagSelectors []struct {
	Decoration string `json:"decoration,omitempty"`
	Pattern    string `json:"pattern,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Extras     string `json:"extras,omitempty"`
}

type ImmutableTagRulesBodyTagSelector struct {
	Decoration string `json:"decoration,omitempty"`
	Pattern    string `json:"pattern,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Extras     string `json:"extras,omitempty"`
}
