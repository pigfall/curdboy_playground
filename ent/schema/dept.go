package schema

import (
	"entgo.io/ent"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Dept holds the schema definition for the Dept entity.
type Dept struct {
	ent.Schema
}

func (Dept) Annotations() []schema.Annotation {
		return []schema.Annotation{
			entproto.Message(entproto.PackageName(PACKAGE_NAME)) ,// tell the entproto ,the schema needed to be generated to proto
		entproto.BuildInCURDService(),
		} 
}

const DEPT_ID="dept_id"

// Fields of the Dept.
func (Dept) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").// the primary key
			StorageKey(DEPT_ID). // override the primary key from id to dept_id, the primary column in db will be dept_id
			NotEmpty(). 
			MaxLen(100).
			Annotations(
				entproto.Field(
					1, // will map to proto file 1
					entproto.FieldName(DEPT_ID), // will map to proto field with name dept_id
				),
			), // add out proto annotation
	}
}

const DPET_ID="dept_id"
const DEPT="dept"

// Edges of the Dept.
func (Dept) Edges() []ent.Edge {
	return []ent.Edge{
		// dept m2m user
		edge.To(
			USER,User.Type,
		).StorageKey(
			edge.Columns(DEPT_ID,USER_ID),// the column for join table will be dept_id,user_id
		).Annotations(
			entproto.FieldIgnore(),// ignore the edge field for the proto of dept
		),
	}
}
