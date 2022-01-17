
// auto generated not edit
package service

import(
	
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	
	"context"
	
	"github.com/pigfall/curdboy_playground/internal/biz"
	
	"google.golang.org/protobuf/types/known/emptypb"
	
	"github.com/go-kratos/kratos/v2/log"
)

type DeptService struct {
	pb.UnimplementedDeptServiceServer
	bizIns *biz.DeptService
	logger *log.Helper
}

func NewDeptService(rawLogger log.Logger ,bizIns *biz.DeptService)*DeptService{
	return &DeptService{
		bizIns:bizIns,
		logger:log.NewHelper(rawLogger),
	}
}




	
	
		func (this *DeptService) Create (ctx context.Context, req *pb.Dept) (*pb.DeptId,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Create(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	func (this *DeptService) Update (ctx context.Context, req *pb.Dept) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.Update(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	func (this *DeptService) Delete (ctx context.Context, req *pb.DeptId) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.Delete(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	
		func (this *DeptService) FindById (ctx context.Context, req *pb.DeptId) (*pb.Dept,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.FindById(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *DeptService) Find (ctx context.Context, req *pb.DeptPageQuery) (*pb.DeptWithCount,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Find(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *DeptService) Count (ctx context.Context, req *pb.Dept) (*pb.CountNumber,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.Count(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	
		func (this *DeptService) AddUser (ctx context.Context, req *pb.DeptAddUserReq) (*pb.UserId,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.AddUser(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	

	
	func (this *DeptService) AddUserById (ctx context.Context, req *pb.DeptIdAndUserId) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.AddUserById(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	func (this *DeptService) RemoveUser (ctx context.Context, req *pb.DeptRemoveUserReq) (*emptypb.Empty,error){
			this.logger.Debugf("request param: %+v\n",req)
			res,err := this.bizIns.RemoveUser(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}
			return res,err
		}
	

	
	
		func (this *DeptService) FindUsers (ctx context.Context, req *pb.DeptFindUsersReq) (*pb.DeptFindUsersRes,error){
			this.logger.Debugf("Request param %+v\n",req)
			res,err := this.bizIns.FindUsers(ctx,req)
			if err != nil{
				this.logger.Debug(err)
			}else{
				this.logger.Debugf("Service response: %+v\n",res)
			}
			return res,err
		}
	



