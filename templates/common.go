package templates

type TemplateContext struct {
	Entity             string
	LowerEntity        string
	Initial            string
	Module             string
	ProjectName        string
	ProjectDescription string
	ProjectVersion     string
	Frontend           bool
	NGINX              bool
	MySQL              bool
	Postgres           bool
	MongoDB            bool
	Redis              bool
}

type FileConfiguration struct {
	Entity          string
	Template        string
	TemplateContext TemplateContext
	Repeat          bool
	RepeatEntities  []string
}
