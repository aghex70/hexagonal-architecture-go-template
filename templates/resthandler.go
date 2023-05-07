package templates

const RestHandlerTemplate = `package {{.LowerEntity}}

import (
    "github.com/gin-gonic/gin"
	_ "{{.Module}}/docs"
	"{{.Module}}/internal/core/domain"
	"{{.Module}}/internal/core/ports"
	"{{.Module}}/internal/handlers"
	"net/http"
)

type {{.Entity}}Handler struct {
	{{.LowerEntity}}Service ports.{{.Entity}}Servicer
}

// Create{{.Entity}} godoc
// @Summary Create a {{.LowerEntity}}
// @Description Create a {{.LowerEntity}}
// @ID create-{{.LowerEntity}}
// @Accept json
// @Produce json
// @Param {{.LowerEntity}} body domain.{{.Entity}} true "{{.LowerEntity}} data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.{{.Entity}}, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /{{.LowerEntity}} [post]
func (h {{.Entity}}Handler) Create{{.Entity}}(c *gin.Context) {
	var r ports.Create{{.Entity}}Request

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n{{.Initial}}, err := h.{{.LowerEntity}}Service.Create(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"{{.LowerEntity}}": n{{.Initial}}})
}

// Update{{.Entity}} godoc
// @Summary Update a {{.LowerEntity}}
// @Description Update a {{.LowerEntity}}
// @ID update-{{.LowerEntity}}
// @Accept json
// @Produce json
// @Param  id path string true "{{.LowerEntity}} ID" {{.LowerEntity}} body domain.{{.Entity}} true "{{.LowerEntity}} data"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.{{.Entity}}, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /{{.LowerEntity}}/{id} [put]
func (h {{.Entity}}Handler) Update{{.Entity}}(c *gin.Context) {
	var r ports.Update{{.Entity}}Request

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u{{.Initial}}, err := h.{{.LowerEntity}}Service.Update(context.TODO(), r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"{{.LowerEntity}}": u{{.Initial}}})
}

// Get{{.Entity}} godoc
// @Summary Get a {{.LowerEntity}}
// @Description Get a {{.LowerEntity}} by ID
// @ID get-{{.LowerEntity}}
// @Accept json
// @Produce json
// @Param id path string true "{{.LowerEntity}} ID"
// @Success 200 {object} handlers.JSONOKResponse{data=domain.{{.Entity}}, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /{{.LowerEntity}}/{id} [get]
func (h {{.Entity}}Handler) Get{{.Entity}}(c *gin.Context) {
	uuid := c.Param("uuid")
	{{.Initial}}{{.Initial}}, err := h.{{.LowerEntity}}Service.Get(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"{{.LowerEntity}}": {{.Initial}}{{.Initial}}})
}

// Delete{{.Entity}} godoc
// @Summary Delete a {{.LowerEntity}}
// @Description Delete a {{.LowerEntity}} by ID
// @ID delete-{{.LowerEntity}}
// @Accept json
// @Produce json
// @Param id path string true "{{.LowerEntity}} ID"
// @Success 204 {object} handlers.JSONOKResponse{data=domain.{{.Entity}}, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /{{.LowerEntity}}/{id} [delete]
func (h {{.Entity}}Handler) Delete{{.Entity}}(c *gin.Context) {
	uuid := c.Param("uuid")
	err := h.{{.LowerEntity}}Service.Delete(context.TODO(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}

// List godoc
// @Summary Retrieve all {{.TableSuffix}}
// @Description Retrieve all {{.TableSuffix}}
// @ID list-{{.LowerEntity}}
// @Accept json
// @Produce json
// @Success 200 {object} handlers.JSONOKResponse{data=domain.{{.Entity}}, statusCode=int}
// @Failure 400 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 404 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Failure 500 {object} handlers.JSONErrorResponse{message=string, statusCode=int, message=string}
// @Router /{{.TableSuffix}} [get]
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
