
package data

import(
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"github.com/pigfall/curdboy_playground/ent"
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	"github.com/pigfall/curdboy_playground/ent/predicate"
	"github.com/pigfall/curdboy_playground/ent/car"
	"github.com/pigfall/curdboy_playground/ent/dept"
	"github.com/pigfall/curdboy_playground/ent/user"

)




// fields in schema

 // field in protobuf// CarId// UserId
//
// filter
func buildCarQueryFilter(req *pb.Car)predicate.Car{
	
	filters := make([]predicate.Car,0,1)
		
		
			if len(req.CarId) >0 {
				filters = append(filters,car.IDEQ(req.CarId))
			}
		
		
			// userid car_id
						
							if req.UserId != nil{
								filters = append(
									filters,
									car.HasUserWith(
										user.IDEQ(req.UserId.GetValue()),
									),
								)
							}
  //TODO

	if len(filters) == 0{
		return nil
	}

	return car.And(filters...)
}

func mutationSetCar(mutation *ent.CarMutation,req *pb.Car){
		
			if len((req.CarId)) > 0{
				mutation.SetID(req.CarId)
			}
		
			
			
						
						if len(req.UserId.GetValue()) > 0{
							mutation.SetUserID(req.UserId.GetValue())
						}
}


// fields in schema

 // field in protobuf// DeptId
//
// filter
func buildDeptQueryFilter(req *pb.Dept)predicate.Dept{
	
	filters := make([]predicate.Dept,0,1)
		
		
			if len(req.DeptId) >0 {
				filters = append(filters,dept.IDEQ(req.DeptId))
			}
  //TODO

	if len(filters) == 0{
		return nil
	}

	return dept.And(filters...)
}

func mutationSetDept(mutation *ent.DeptMutation,req *pb.Dept){
		
			if len((req.DeptId)) > 0{
				mutation.SetID(req.DeptId)
			}
}


// fields in schema// Phone// Race

 // field in protobuf// UserId// Phone// Race
//
// filter
func buildUserQueryFilter(req *pb.User)predicate.User{
	
	filters := make([]predicate.User,0,1)
		
		
			if len(req.UserId) >0 {
				filters = append(filters,user.IDEQ(req.UserId))
			}
		
		
			// phone user_id
				
					if req.Phone != nil{
								filters = append(
									filters,
									user.PhoneEQ(req.Phone.GetValue()),
								)
						}
				
		
		
			// race user_id
				
				
					
								if  req.Race > 0{
									filters = append(
										filters,
										user.RaceEQ(
											user.Race(
												pb.User_Race_name[int32(req.Race)],
											),
										),
									)
								}
  //TODO

	if len(filters) == 0{
		return nil
	}

	return user.And(filters...)
}

func mutationSetUser(mutation *ent.UserMutation,req *pb.User){
		
			if len((req.UserId)) > 0{
				mutation.SetID(req.UserId)
			}
		
			
			
						if req.Phone != nil{mutation.SetPhone(req.Phone.GetValue())
						}
		
			
			
							mutation.SetRace(
								user.Race(
									pb.User_Race_name[int32(req.Race)],
								),
							)
}


func jsonMarshal(v interface{})([]byte,error){
	marshalOption := protojson.MarshalOptions{}
	msg := v.(proto.Message)
	return marshalOption.Marshal(msg)
}

func jsonUnmarshal(dataBytes []byte,v interface{})error{
		unmarshalOptions := protojson.UnmarshalOptions{
			DiscardUnknown: true,
		}
		pbMsg := v.(proto.Message)
		return unmarshalOptions.Unmarshal(dataBytes,pbMsg)
}
