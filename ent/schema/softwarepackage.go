package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SoftwarePackage holds the schema definition for the SoftwarePackage entity.
type SoftwarePackage struct {
	ent.Schema
}

// Fields of the SoftwarePackage.
func (SoftwarePackage) Fields() []ent.Field {
	return []ent.Field {
		field.String("name").Unique(),
	}
}

// Edges of the SoftwarePackage.
func (SoftwarePackage) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("Versions", PackageVersion.Type),
	}
}
