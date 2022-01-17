
package data






import(
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	"context"
	"github.com/pigfall/curdboy_playground/internal/biz"
	"github.com/pigfall/curdboy_playground/ent"
	predicate "github.com/pigfall/curdboy_playground/ent/car"
	"errors"
)

// 实现接口


type CarService struct{
	cli *ent.ClientWrapper
}

// 定义构造方法
func NewCarService(cli *ent.ClientWrapper) biz.DataCarServiceIfce{
	return &CarService{
		cli:cli,
	}
}



	
	func(this *CarService)Create (ctx context.Context,req *pb.Car)(id string,err error){
		createBuilder := this.cli.Car.Create()
		mutation := createBuilder.Mutation()
		mutationSetCar(mutation,req)
		model,err := createBuilder.Save(ctx)
		if err != nil{
			return "",err
		}
		return model.ID,nil
	}
	

	
	func(this *CarService)Update (ctx context.Context,req *pb.Car)(err error){
		updateBuilder := this.cli.Car.Update()
		mutation := updateBuilder.Mutation()
		mutationSetCar(mutation,req)
		_,err = updateBuilder.Save(ctx)
		if err != nil{
			return err
		}
		return nil
	}
	

	
	func(this *CarService)Delete (ctx context.Context,id string)(error){
		return this.cli.Car.DeleteOneID(id).Exec(ctx)
	}
	

	
		func(this *CarService)FindById (ctx context.Context,id string)(*ent.Car,error){
			query := this.cli.Car.Query().Where(predicate.IDEQ(id))
				
					query = query.WithUser()
			model,err := query.Only(ctx)
			if err != nil{
				return nil,err
			}
				
					// TODO ent issue: 现在的ent 的 eager load 有问题，实际上的edge 没有真正load出来
					// 所以我们判断下 edge 是不是真的有 load
					_,edgeLoadedUserErr := model.Edges.UserOrErr()
					if edgeLoadedUserErr != nil{
							 var errNotFound *ent.NotFoundError
							 edgeEntity,err := model.QueryUser().Only(ctx)
							 if err != nil && !errors.As(err,&errNotFound){
									return nil,err
							 }
							 model.Edges.User = edgeEntity
						 
					}
			return model,nil
		}
	

	
		func(this *CarService)Find (ctx context.Context,req *pb.CarPageQuery)([]*ent.Car,int,error){
			reqUnit := &pb.Car{}
			bytes,err := jsonMarshal(req)
			if err != nil {
				return nil,0,err
			}
			err = jsonUnmarshal(bytes,reqUnit)
			if err != nil{
				return nil,0,err
			}

			filter := buildCarQueryFilter(reqUnit)
			if filter == nil {
				filter = predicate.And()
			}
			builder :=  this.cli.Car.Query().Where(filter).Limit(int(req.PageSize)).Offset(int(req.PageSize*req.PageIndex))

			var count int
			if req.PageDataCount.GetValue(){
				countInt32,err := builder.Count(ctx)
				if err != nil{
					return nil,0,err
				}
				count = int(countInt32)
			}
				builder.WithUser()
			res,err := builder.All(ctx)
			return res,count,err
		}
	

	
		func(this *CarService)Count (ctx context.Context,req *pb.Car)(int32,error){
			filter := buildCarQueryFilter(req)
			c,err := this.cli.Car.Query().Where(filter).Count(ctx)
			if err != nil{
				return 0,err
			}
			return int32(c),nil
		}





