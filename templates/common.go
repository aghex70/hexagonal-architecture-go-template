package templates

type TemplateContext struct {
	Entity             string
	LowerEntity        string
	Initial            string
	Module             string
	ProjectName        string
	ProjectDescription string
	ProjectVersion     string
}

type FileConfiguration struct {
	Entity          string
	Template        string
	TemplateContext TemplateContext
	Repeat          bool
	RepeatEntities  []string
}

type ComposeConfiguration struct {
	Frontend bool
	NGINX    bool
	MySQL    bool
	Postgres bool
	MongoDB  bool
	Redis    bool
}
