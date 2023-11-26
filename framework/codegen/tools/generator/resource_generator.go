package main

import (
	specresource "github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
)

const (
	BoolAttributeType         = "BoolAttribute"
	StringAttributeType       = "StringAttribute"
	NumberAttributeType       = "NumberAttribute"
	Int64AttributeType        = "Int64Attribute"
	MapAttributeType          = "MapAttribute"
	ListAttributeType         = "ListAttribute"
	ObjectAttributeType       = "ObjectAttribute"
	SingleNestedAttributeType = "SingleNestedAttribute"
	ListNestedAttributeType   = "ListNestedAttribute"
)

const (
	BoolElementType   = "BoolType"
	StringElementType = "StringType"
	NumberElementType = "NumberType"
	Int64ElementType  = "Int64Type"
)

type ResourceGenerator struct {
	ResourceConfig ResourceConfig
	Schema         SchemaGenerator
}

type SchemaGenerator struct {
	Name        string
	Description string
	Attributes  AttributesGenerator
}

type AttributeGenerator struct {
	Name          string
	AttributeType string
	ElementType   string
	Required      bool
	Description   string
	Computed      bool
	Sensitive     bool

	NestedAttributes AttributesGenerator
}

func (g AttributeGenerator) String() string {
	return renderTemplate(attributeTemplate, g)
}

type AttributesGenerator []AttributeGenerator

func (g AttributesGenerator) String() string {
	return renderTemplate(attributesTemplate, g)
}

func NewResourceGenerator(cfg ResourceConfig, spec specresource.Resource) ResourceGenerator {
	return ResourceGenerator{
		ResourceConfig: cfg,
		Schema: SchemaGenerator{
			Name:        cfg.Name,
			Description: cfg.Description,
			Attributes:  generateAttributes(spec.Schema.Attributes),
		},
	}
}

func (g *ResourceGenerator) GenerateSchemaFunctionCode() string {
	return renderTemplate(schemaFunctionTemplate, g)
}

func (g *ResourceGenerator) GenerateCRUDStubCode() string {
	return renderTemplate(crudStubsTemplate, g)
}

func (g *ResourceGenerator) GenerateResourceCode() string {
	return renderTemplate(resourceTemplate, g)
}

func (g *ResourceGenerator) GenerateModelCode() string {
	return renderTemplate(modelTemplate, g)
}

func generateAttributes(attrs specresource.Attributes) []AttributeGenerator {
	generatedAttrs := []AttributeGenerator{}
	for _, attr := range attrs {
		generatedAttr := AttributeGenerator{
			Name: attr.Name,
		}
		switch {
		case attr.Bool != nil:
			if attr.Bool.Description != nil {
				generatedAttr.Description = *attr.Bool.Description
			}
			generatedAttr.AttributeType = BoolAttributeType
			generatedAttr.Required = isRequired(attr.Bool.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.Bool.ComputedOptionalRequired)
		case attr.String != nil:
			if attr.String.Description != nil {
				generatedAttr.Description = *attr.String.Description
			}
			generatedAttr.AttributeType = StringAttributeType
			generatedAttr.Required = isRequired(attr.String.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.String.ComputedOptionalRequired)
		case attr.Number != nil:
			if attr.Number.Description != nil {
				generatedAttr.Description = *attr.Number.Description
			}
			generatedAttr.AttributeType = NumberAttributeType
			generatedAttr.Required = isRequired(attr.Number.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.Number.ComputedOptionalRequired)
		case attr.Int64 != nil:
			if attr.Int64.Description != nil {
				generatedAttr.Description = *attr.Int64.Description
			}
			generatedAttr.AttributeType = Int64AttributeType
			generatedAttr.Required = isRequired(attr.Int64.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.Int64.ComputedOptionalRequired)
		case attr.Map != nil:
			if attr.Map.Description != nil {
				generatedAttr.Description = *attr.Map.Description
			}
			generatedAttr.AttributeType = MapAttributeType
			generatedAttr.Required = isRequired(attr.Map.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.Map.ComputedOptionalRequired)
			generatedAttr.ElementType = getElementType(attr.Map.ElementType)
		case attr.List != nil:
			if attr.List.Description != nil {
				generatedAttr.Description = *attr.List.Description
			}
			generatedAttr.AttributeType = ListAttributeType
			generatedAttr.Required = isRequired(attr.List.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.List.ComputedOptionalRequired)
			generatedAttr.ElementType = getElementType(attr.List.ElementType)
		case attr.SingleNested != nil:
			if attr.SingleNested.Description != nil {
				generatedAttr.Description = *attr.SingleNested.Description
			}
			generatedAttr.AttributeType = SingleNestedAttributeType
			generatedAttr.Required = isRequired(attr.SingleNested.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.SingleNested.ComputedOptionalRequired)
			generatedAttr.NestedAttributes = generateAttributes(attr.SingleNested.Attributes)
		case attr.ListNested != nil:
			if attr.ListNested.Description != nil {
				generatedAttr.Description = *attr.ListNested.Description
			}
			generatedAttr.AttributeType = ListNestedAttributeType
			generatedAttr.Required = isRequired(attr.ListNested.ComputedOptionalRequired)
			generatedAttr.Computed = isComputed(attr.ListNested.ComputedOptionalRequired)
			generatedAttr.NestedAttributes = generateAttributes(attr.ListNested.NestedObject.Attributes)
		}
		generatedAttrs = append(generatedAttrs, generatedAttr)
	}
	return generatedAttrs
}

func isComputed(c specschema.ComputedOptionalRequired) bool {
	return c == specschema.Computed || c == specschema.ComputedOptional
}

func isRequired(c specschema.ComputedOptionalRequired) bool {
	return c == specschema.Required
}

func getElementType(e specschema.ElementType) string {
	switch {
	case e.Bool != nil:
		return BoolElementType
	case e.String != nil:
		return StringElementType
	case e.Number != nil:
		return NumberElementType
	case e.Int64 != nil:
		return Int64ElementType
	}
	panic("unsupported element type")
}
