
// auto generated not edit
package service

import(
	
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	
	"context"
	
	"github.com/pigfall/curdboy_playground/internal/biz"
	
	"google.golang.org/protobuf/types/known/emptypb"
	
	"github.com/go-kratos/kratos/v2/log"
)

type CarService struct {
	pb.UnimplementedCarServiceServer
	bizIns *biz.CarService
	logger *log.Helper
}

func NewCarService(rawLogger log.Logger ,bizIns *biz.CarService)*CarService{
	return &CarService{
		bizIns:bizIns,
		logger:log.NewHelper(rawLogger),
	}
}




	
	
		func (this *CarService) Create (ctx context.Context, req *pb.Car) (*pb.CarId,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Create(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	func (this *CarService) Update (ctx context.Context, req *pb.Car) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.Update(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	func (this *CarService) Delete (ctx context.Context, req *pb.CarId) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.Delete(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	
		func (this *CarService) FindById (ctx context.Context, req *pb.CarId) (*pb.Car,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.FindById(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *CarService) Find (ctx context.Context, req *pb.CarPageQuery) (*pb.CarWithCount,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Find(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *CarService) Count (ctx context.Context, req *pb.Car) (*pb.CountNumber,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Count(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	



