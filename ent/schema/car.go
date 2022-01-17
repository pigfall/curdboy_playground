package schema

import (
	"entgo.io/ent"
	"google.golang.org/protobuf/types/descriptorpb"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

func (Car) Annotations() []schema.Annotation {
		return []schema.Annotation{
			entproto.Message(entproto.PackageName(PACKAGE_NAME)) ,// tell the entproto ,the schema needed to be generated to proto
		entproto.BuildInCURDService(),
		} 
}

const CAR_ID= "car_id"

const CAR="car"

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").// the primary key
			StorageKey(CAR_ID). // override the primary key from id to car_id, the primary column in db will be car_id
			NotEmpty(). 
			MaxLen(100).
			Annotations(
				entproto.Field(
					1, // will map to proto file 1
					entproto.FieldName(CAR_ID), // will map to proto field with name car_id
				),
			), // add out proto annotation
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From(
			USER,User.Type,
		).Ref(CAR).
		Unique().
		Annotations(
			entproto.Field(
				2,
				entproto.FieldName(USER_ID),// the proto of car will add the field which name is user_id
				entproto.Type(descriptorpb.FieldDescriptorProto_TYPE_STRING),// the type of field will be string
			),
		),
		
	}
}
