package templates

const RestHandlerTemplate = `package {{.LowerEntity}}

import (
    "github.com/gin-gonic/gin"
	"{{.Module}}/internal/core/ports"
	"net/http"
)

type {{.Entity}}Handler struct {
	{{.LowerEntity}}Service ports.{{.Entity}}Servicer
}

func (h {{.Entity}}Handler) Create{{.Entity}}(c *gin.Context) {
	var r ports.Create{{.Entity}}Request

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n{{.Initial}}, err := h.{{.LowerEntity}}Service.Create(nil, r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"{{.LowerEntity}}": n{{.Initial}}})
}

func (h {{.Entity}}Handler) Update{{.Entity}}(c *gin.Context) {
	var r ports.Update{{.Entity}}Request

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u{{.Initial}}, err := h.{{.LowerEntity}}Service.Update(nil, r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"{{.LowerEntity}}": u{{.Initial}}})
}

func (h {{.Entity}}Handler) Get{{.Entity}}(c *gin.Context) {
	{{.Initial}}{{.Initial}}, err := h.{{.LowerEntity}}Service.Get(nil, r, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"{{.LowerEntity}}": {{.Initial}}{{.Initial}}})
}

func (h {{.Entity}}Handler) Delete{{.Entity}}(c *gin.Context) {
	err := h.{{.LowerEntity}}Service.Delete(nil, r, payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h {{.Entity}}Handler) List(c *gin.Context) {
	{{.Initial}}s, err := h.{{.LowerEntity}}Service.List(nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"{{.TableSuffix}}": {{.Initial}}s})
}

func New{{.Entity}}Handler({{.Initial}}s ports.{{.Entity}}Servicer) {{.Entity}}Handler {
	return {{.Entity}}Handler{
		{{.LowerEntity}}Service: {{.Initial}}s,
	}
}
`

func GetRestHandlerFileConfiguration(tc TemplateContext) []FileConfiguration {
	return []FileConfiguration{
		{
			Template:        RestHandlerTemplate,
			TemplateContext: tc,
		},
	}
}
