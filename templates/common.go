package templates

type TemplateContext struct {
	Entity             string
	LowerEntity        string
	Initial            string
	Module             string
	ProjectName        string
	ProjectDescription string
}

type FileConfiguration struct {
	Entity          string
	Template        string
	TemplateContext TemplateContext
	Repeat          bool
	RepeatEntities  []string
}
