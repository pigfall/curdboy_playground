
// auto generated not edit
package service

import(
	
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	
	"context"
	
	"github.com/pigfall/curdboy_playground/internal/biz"
	
	"google.golang.org/protobuf/types/known/emptypb"
	
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	bizIns *biz.UserService
	logger *log.Helper
}

func NewUserService(rawLogger log.Logger ,bizIns *biz.UserService)*UserService{
	return &UserService{
		bizIns:bizIns,
		logger:log.NewHelper(rawLogger),
	}
}




	
	
		func (this *UserService) Create (ctx context.Context, req *pb.User) (*pb.UserId,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Create(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	func (this *UserService) Update (ctx context.Context, req *pb.User) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.Update(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	func (this *UserService) Delete (ctx context.Context, req *pb.UserId) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.Delete(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	
		func (this *UserService) FindById (ctx context.Context, req *pb.UserId) (*pb.User,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.FindById(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *UserService) Find (ctx context.Context, req *pb.UserPageQuery) (*pb.UserWithCount,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Find(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *UserService) Count (ctx context.Context, req *pb.User) (*pb.CountNumber,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Count(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *UserService) AddCar (ctx context.Context, req *pb.Car) (*pb.CarId,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.AddCar(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	func (this *UserService) AddCarById (ctx context.Context, req *pb.UserIdAndCarId) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.AddCarById(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	func (this *UserService) RemoveCar (ctx context.Context, req *pb.UserRemoveCarReq) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.RemoveCar(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	
		func (this *UserService) FindCars (ctx context.Context, req *pb.UserFindCarsReq) (*pb.UserFindCarsRes,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.FindCars(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *UserService) FindDepts (ctx context.Context, req *pb.UserFindDeptsReq) (*pb.UserFindDeptsRes,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.FindDepts(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	



