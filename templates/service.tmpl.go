package templates

const ServiceTmplStr = `package services

import (
	"context"
	"github.com/jinzhu/copier"
	"{{ .ProjectName }}/dao"
	"{{ .ProjectName }}/dao/models"
	"{{ .ProjectName }}/services/internal"
)

type {{ .StructName }}Service struct {
	Q *dao.Query
}

func New{{ .StructName }}Service() *{{ .StructName }}Service {
	return &{{ .StructName }}Service{Q: dao.Q,}
}

// GetList 查询列表
type GetList{{ .StructName }}Params struct {
{{- range .SearchableFields }}
{{- if eq .Name "Id" }}
	{{ .Name }} {{ .Type }} ` + "`form:\"{{ .TagName }}\" binding:\"required\"`" + `
{{- else }}
	{{ .Name }} {{ .Type }} ` + "`form:\"{{ .TagName }}\"`" + `
{{- end }}
{{- end }}
	internal.PageInfo
}

func (s *{{ .StructName }}Service) GetList(ctx context.Context, params *GetList{{ .StructName }}Params) ([]*models.{{ .StructName }}, error) {
	querySet := s.Q.{{ .StructName }}
	stmt := querySet.WithContext(ctx)
{{- range .SearchableFields }}
{{- if .IsPointer }}
	if params.{{ .Name }} != nil {
		stmt = stmt.Where(querySet.{{ .Name }}.Eq(*params.{{ .Name }}))
	}
{{- else }}
	stmt = stmt.Where(querySet.{{ .Name }}.Eq(params.{{ .Name }}))
{{- end }}
{{- end }}

	return stmt.Find()
}

// Create 创建
type Create{{ .StructName }}Params struct {
{{- range .AllFields }}
	{{ .Name }} {{ .Type }} ` + "`json:\"{{ .TagName }}\"`" + `
{{- end }}
}

func (s *{{ .StructName }}Service) Create(ctx context.Context, params *Create{{ .StructName }}Params) error {
	var obj models.{{ .StructName }}
	copier.Copy(&obj, params)
	return s.Q.{{ .StructName }}.WithContext(ctx).Create(&obj)
}

// Update 更新
type Update{{ .StructName }}Params struct {
{{- range .AllFields }}
{{- if eq .Name "Id" }}
	{{ .Name }} {{ .Type }} ` + "`form:\"{{ .TagName }}\" binding:\"required\"`" + `
{{- else }}
	{{ .Name }} {{ .Type }} ` + "`form:\"{{ .TagName }}\"`" + `
{{- end }}
{{- end }}
}

func (s *{{ .StructName }}Service) Update(ctx context.Context, params *Update{{ .StructName }}Params) error {
	var obj = s.Q.{{ .StructName }}

	_, err := s.Q.{{ .StructName }}.WithContext(ctx).Where(obj.ID.Eq(params.ID)).Updates(params)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
type Delete{{ .StructName }}Params struct {
{{- range .AllFields }}
{{- if eq .Name "ID" }}
	{{ .Name }} {{ .Type }} ` + "`form:\"{{ .TagName }}\" binding:\"required\"`" + `
{{- end }}
{{- end }}
}

func (s *{{ .StructName }}Service) Delete(ctx context.Context, params *Delete{{ .StructName }}Params) error {
	_, err := s.Q.{{ .StructName }}.WithContext(ctx).
		Where(s.Q.{{ .StructName }}.ID.Eq(params.ID)).
		Delete()
	return err
}
`
