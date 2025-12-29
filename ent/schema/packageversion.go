package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PackageVersion holds the schema definition for the PackageVersion entity.
type PackageVersion struct {
	ent.Schema
}

// Fields of the PackageVersion.
func (PackageVersion) Fields() []ent.Field {
	return []ent.Field {
		field.String("Version").Unique(),
		field.Time("IndexedAt").Default(time.Now),
	}
}

// Edges of the PackageVersion.
func (PackageVersion) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("Package", SoftwarePackage.Type).Ref("Versions").Unique(),
	}
}
