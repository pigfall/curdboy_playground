
package data






import(
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	"context"
	"github.com/pigfall/curdboy_playground/internal/biz"
	"github.com/pigfall/curdboy_playground/ent"
	predicate "github.com/pigfall/curdboy_playground/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"errors"
			
			
			preCar "github.com/pigfall/curdboy_playground/ent/car"
			
			
			preDept "github.com/pigfall/curdboy_playground/ent/dept"
)

// 实现接口


type UserService struct{
	cli *ent.ClientWrapper
}

// 定义构造方法
func NewUserService(cli *ent.ClientWrapper) biz.DataUserServiceIfce{
	return &UserService{
		cli:cli,
	}
}



	
	func(this *UserService)Create (ctx context.Context,req *pb.User)(id string,err error){
		createBuilder := this.cli.User.Create()
		mutation := createBuilder.Mutation()
		mutationSetUser(mutation,req)
		model,err := createBuilder.Save(ctx)
		if err != nil{
			return "",err
		}
		return model.ID,nil
	}
	

	
	func(this *UserService)Update (ctx context.Context,req *pb.User)(err error){
		updateBuilder := this.cli.User.Update()
		mutation := updateBuilder.Mutation()
		mutationSetUser(mutation,req)
		_,err = updateBuilder.Save(ctx)
		if err != nil{
			return err
		}
		return nil
	}
	

	
	func(this *UserService)Delete (ctx context.Context,id string)(error){
		return this.cli.User.DeleteOneID(id).Exec(ctx)
	}
	

	
		func(this *UserService)FindById (ctx context.Context,id string)(*ent.User,error){
			query := this.cli.User.Query().Where(predicate.IDEQ(id))
				
				
			model,err := query.Only(ctx)
			if err != nil{
				return nil,err
			}
				
				
			return model,nil
		}
	

	
		func(this *UserService)Find (ctx context.Context,req *pb.UserPageQuery)([]*ent.User,int,error){
			reqUnit := &pb.User{}
			bytes,err := jsonMarshal(req)
			if err != nil {
				return nil,0,err
			}
			err = jsonUnmarshal(bytes,reqUnit)
			if err != nil{
				return nil,0,err
			}

			filter := buildUserQueryFilter(reqUnit)
			if filter == nil {
				filter = predicate.And()
			}
			builder :=  this.cli.User.Query().Where(filter).Limit(int(req.PageSize)).Offset(int(req.PageSize*req.PageIndex))

			var count int
			if req.PageDataCount.GetValue(){
				countInt32,err := builder.Count(ctx)
				if err != nil{
					return nil,0,err
				}
				count = int(countInt32)
			}
			res,err := builder.All(ctx)
			return res,count,err
		}
	

	
		func(this *UserService)Count (ctx context.Context,req *pb.User)(int32,error){
			filter := buildUserQueryFilter(req)
			c,err := this.cli.User.Query().Where(filter).Count(ctx)
			if err != nil{
				return 0,err
			}
			return int32(c),nil
		}

	

		
			func (this *UserService) AddCar (ctx context.Context,req  *pb.Car)(id string,err error){
				createBuilder := this.cli.Car.Create()
				mutation := createBuilder.Mutation()
				mutationSetCar(mutation,req)
				model,err := createBuilder.Save(ctx)
				if err != nil{
					return "",err
				}
				return model.ID,nil
			}

	
		func(this *UserService)AddCarById (ctx context.Context,req *pb.UserIdAndCarId)(error){
			
							// TODO handle o2m ,m2m
							err := this.cli.User.Update().Where(
								predicate.IDEQ(req.UserId),
							).AddCarIDs(req.CarId).Exec(ctx)
							if err != nil{
								var targetErr *sqlgraph.ConstraintError
								if errors.As(err,&targetErr){
									return nil
								}
								return err
							}
							return nil
		}

	
		func(this *UserService)RemoveCar (ctx context.Context,req *pb.UserRemoveCarReq)(error){
			
							queryEntity := &pb.Car{}
							bytes,err := jsonMarshal(req)
							if err != nil{
								return err
							}
							err = jsonUnmarshal(bytes,queryEntity)
							if err != nil{
								return err
							}
							filter := buildCarQueryFilter(queryEntity)
							if filter == nil {
								//return fmt.Errorf("TODO request error, filter is zero") // TODO error handle
								// filter = preCar.IDEQ(req.CarId)
								// filter = preCar.IDEQ(req.CarId)
								filter = preCar.HasUserWith(
									predicate.IDEQ(req.UserId),
								)
							}
							edgeEntities,err := this.cli.Car.Query().Where(filter).All(ctx)
							if err != nil{
								return err
							}
							if len(edgeEntities) == 0{
								return nil
							}
							edgeEntityIds := make([]string,len(edgeEntities))
							for i,e  := range edgeEntities{
								 edgeEntityIds[i]= e.ID
							}
							// TODO handle o2m ,m2m
							return this.cli.User.Update().Where(
								predicate.IDEQ(req.UserId),
							).RemoveCarIDs(edgeEntityIds...).Exec(ctx)
		}

	
		
		
		func(this *UserService)FindCars (ctx context.Context,req *pb.UserFindCarsReq)(data []*ent.Car,total int ,err error){
			queryEntity := &pb.Car{}
			bytes,err := jsonMarshal(req)
			if err != nil {
				return nil,0,err
			}
			err = jsonUnmarshal(bytes,queryEntity)
			if err != nil{
				return nil,0,err
			}
			filter := buildCarQueryFilter(queryEntity)
			if filter == nil{
				filter = preCar.And()
			}
			query := this.cli.Car.Query().Where(
				filter,
				preCar.HasUserWith(
					// TODO
					predicate.IDEQ(req.UserId),
				),
			)
			total,err = query.Count(ctx)
			if err != nil{
				return nil,0,err
			}

			data , err = query.Offset(int(req.PageSize)*int(req.PageIndex)).Limit(int(req.PageSize)).All(ctx)
			if err != nil {
				return nil,total,err
			}
			return data,total,nil
		}

	
		
		
		func(this *UserService)FindDepts (ctx context.Context,req *pb.UserFindDeptsReq)(data []*ent.Dept,total int ,err error){
			queryEntity := &pb.Dept{}
			bytes,err := jsonMarshal(req)
			if err != nil {
				return nil,0,err
			}
			err = jsonUnmarshal(bytes,queryEntity)
			if err != nil{
				return nil,0,err
			}
			filter := buildDeptQueryFilter(queryEntity)
			if filter == nil{
				filter = preDept.And()
			}
			query := this.cli.Dept.Query().Where(
				filter,
				preDept.HasUserWith(
					// TODO
					predicate.IDEQ(req.UserId),
				),
			)
			total,err = query.Count(ctx)
			if err != nil{
				return nil,0,err
			}

			data , err = query.Offset(int(req.PageSize)*int(req.PageIndex)).Limit(int(req.PageSize)).All(ctx)
			if err != nil {
				return nil,total,err
			}
			return data,total,nil
		}





