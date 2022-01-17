package schema

import (
	"entgo.io/ent"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

const PACKAGE_NAME= "contacts"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
		return []schema.Annotation{
			entproto.Message(entproto.PackageName(PACKAGE_NAME)) ,// tell the entproto ,the schema needed to be generated to proto
		entproto.BuildInCURDService(),
		} 
}

const USER_ID="user_id"

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").// the primary key
		StorageKey(USER_ID). // override the primary key from id to  user_id, the primary column in db will be user_id
			NotEmpty(). 
			MaxLen(100).
			Annotations(
				entproto.Field(
					1, // will map to proto file 1
					entproto.FieldName(USER_ID), // will map to proto field with name user_id
				),
			), // add out proto annotation

			field.String("phone").
			NotEmpty().
			MaxLen(100).
			Optional(). // optional field
			Comment("手机号"). // comment
			Annotations(
				entproto.Field(
					2, // if without the entproto.FieldName option, the mapped proto field name will be the same as schema field. 
				),
			),

			// 人种
			field.Enum("race").
				Values( // white, black,yellow
						RACE_YELLOW,
						RACE_WHITE,
						RACE_BLACK,
				).
				Annotations(
					entproto.Field(
						3,
					),
					entproto.Enum( // map to proto enum
						map[string]int32{
							RACE_YELLOW:1,
							RACE_WHITE:2,
							RACE_BLACK:3,
						},
					),
				),
	}
}

const RACE_YELLOW ="YELLOW"
const RACE_WHITE ="WHITE"
const RACE_BLACK ="BLACK"

const USER="user"

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// user o2m car
		// user has many cars 
		edge.To(
			CAR,Car.Type,
		).StorageKey(
			edge.Column(USER_ID), // in schema of car, will add the field user_id
		).Annotations(
			entproto.FieldIgnore(), // in proto of user, will ignore this edge
		),

		// user m2m dept back ref
		edge.From(
				DEPT,Dept.Type,
		).Ref(USER).
		Annotations(
				entproto.FieldIgnore(),
		),
	}
}

