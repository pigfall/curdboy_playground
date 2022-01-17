
package biz

import(
	
	"github.com/pigfall/curdboy_playground/ent"
	
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	
	"strings"
	
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	
)

// define scheme to pb convert func


func toPBCar(model *ent.Car)(*pb.Car){
	output := &pb.Car{}
	
		
		
		output.CarId = model.ID
			
		
		
		
			if model.Edges.User != nil{
					output.UserId = wrapperspb.String(model.Edges.User.ID)
				}

	return output
}
func toPBDept(model *ent.Dept)(*pb.Dept){
	output := &pb.Dept{}
	
		
		
		output.DeptId = model.ID
			

	return output
}
func toPBUser(model *ent.User)(*pb.User){
	output := &pb.User{}
	
		
		
		output.UserId = model.ID
			
			
		
		
		output.Phone = wrapperspb.String(model.Phone)
			
			
		
		
		
								output.Race = pb.User_Race(pb.User_Race_value[strings.ToUpper(model.Race.String())])
			
			

	return output
}


