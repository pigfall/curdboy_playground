
package data






import(
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	"context"
	"github.com/pigfall/curdboy_playground/internal/biz"
	"github.com/pigfall/curdboy_playground/ent"
	predicate "github.com/pigfall/curdboy_playground/ent/dept"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"errors"
			
			
			preUser "github.com/pigfall/curdboy_playground/ent/user"
)

// 实现接口


type DeptService struct{
	cli *ent.ClientWrapper
}

// 定义构造方法
func NewDeptService(cli *ent.ClientWrapper) biz.DataDeptServiceIfce{
	return &DeptService{
		cli:cli,
	}
}



	
	func(this *DeptService)Create (ctx context.Context,req *pb.Dept)(id string,err error){
		createBuilder := this.cli.Dept.Create()
		mutation := createBuilder.Mutation()
		mutationSetDept(mutation,req)
		model,err := createBuilder.Save(ctx)
		if err != nil{
			return "",err
		}
		return model.ID,nil
	}
	

	
	func(this *DeptService)Update (ctx context.Context,req *pb.Dept)(err error){
		updateBuilder := this.cli.Dept.Update()
		mutation := updateBuilder.Mutation()
		mutationSetDept(mutation,req)
		_,err = updateBuilder.Save(ctx)
		if err != nil{
			return err
		}
		return nil
	}
	

	
	func(this *DeptService)Delete (ctx context.Context,id string)(error){
		return this.cli.Dept.DeleteOneID(id).Exec(ctx)
	}
	

	
		func(this *DeptService)FindById (ctx context.Context,id string)(*ent.Dept,error){
			query := this.cli.Dept.Query().Where(predicate.IDEQ(id))
				
			model,err := query.Only(ctx)
			if err != nil{
				return nil,err
			}
				
			return model,nil
		}
	

	
		func(this *DeptService)Find (ctx context.Context,req *pb.DeptPageQuery)([]*ent.Dept,int,error){
			reqUnit := &pb.Dept{}
			bytes,err := jsonMarshal(req)
			if err != nil {
				return nil,0,err
			}
			err = jsonUnmarshal(bytes,reqUnit)
			if err != nil{
				return nil,0,err
			}

			filter := buildDeptQueryFilter(reqUnit)
			if filter == nil {
				filter = predicate.And()
			}
			builder :=  this.cli.Dept.Query().Where(filter).Limit(int(req.PageSize)).Offset(int(req.PageSize*req.PageIndex))

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
	

	
		func(this *DeptService)Count (ctx context.Context,req *pb.Dept)(int32,error){
			filter := buildDeptQueryFilter(req)
			c,err := this.cli.Dept.Query().Where(filter).Count(ctx)
			if err != nil{
				return 0,err
			}
			return int32(c),nil
		}

	

		// m2m add TODO
			func (this *DeptService) AddUser (ctx context.Context,req  *pb.DeptAddUserReq)(id string,err error){
				
				mutationParam := &pb.User{}
				bytes,err := jsonMarshal(req)
				if err != nil {
					return "",err
				}
				err = jsonUnmarshal(bytes,mutationParam)
				if err != nil{
					return "",err
				}
				createBuilder := this.cli.User.Create()
				mutation := createBuilder.Mutation()
				mutationSetUser(mutation,mutationParam)
				mutation.AddDeptIDs(req.DeptId)
				model,err := createBuilder.Save(ctx)
				if err != nil{
					return "",err
				}
				return model.ID,nil
			}

	
		func(this *DeptService)AddUserById (ctx context.Context,req *pb.DeptIdAndUserId)(error){
			
							// TODO handle o2m ,m2m
							err := this.cli.Dept.Update().Where(
								predicate.IDEQ(req.DeptId),
							).AddUserIDs(req.UserId).Exec(ctx)
							if err != nil{
								var targetErr *sqlgraph.ConstraintError
								if errors.As(err,&targetErr){
									return nil
								}
								return err
							}
							return nil
		}

	
		func(this *DeptService)RemoveUser (ctx context.Context,req *pb.DeptRemoveUserReq)(error){
			
							queryEntity := &pb.User{}
							bytes,err := jsonMarshal(req)
							if err != nil{
								return err
							}
							err = jsonUnmarshal(bytes,queryEntity)
							if err != nil{
								return err
							}
							filter := buildUserQueryFilter(queryEntity)
							if filter == nil {
								//return fmt.Errorf("TODO request error, filter is zero") // TODO error handle
								// filter = preUser.IDEQ(req.UserId)
								// filter = preUser.IDEQ(req.UserId)
								filter = preUser.HasDeptWith(
									predicate.IDEQ(req.DeptId),
								)
							}
							edgeEntities,err := this.cli.User.Query().Where(filter).All(ctx)
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
							return this.cli.Dept.Update().Where(
								predicate.IDEQ(req.DeptId),
							).RemoveUserIDs(edgeEntityIds...).Exec(ctx)
		}

	
		
		
		func(this *DeptService)FindUsers (ctx context.Context,req *pb.DeptFindUsersReq)(data []*ent.User,total int ,err error){
			queryEntity := &pb.User{}
			bytes,err := jsonMarshal(req)
			if err != nil {
				return nil,0,err
			}
			err = jsonUnmarshal(bytes,queryEntity)
			if err != nil{
				return nil,0,err
			}
			filter := buildUserQueryFilter(queryEntity)
			if filter == nil{
				filter = preUser.And()
			}
			query := this.cli.User.Query().Where(
				filter,
				preUser.HasDeptWith(
					// TODO
					predicate.IDEQ(req.DeptId),
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





