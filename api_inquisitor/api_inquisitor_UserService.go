package api_inquisitor

import (
	"context"
	"fmt"
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	api "github.com/pigfall/curdboy_playground/http_api_contacts"
	"log"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"math/rand"
)

type UserServiceInquisitor struct {
	client api.ApiClientIfce
}

func UserServiceInquisitorNew(client api.ApiClientIfce) ApiInquisitor {
	return &UserServiceInquisitor{
		client: client,
	}
}

func (this *UserServiceInquisitor) Check(ctx context.Context) error {

	// UserCreate

	// UserFind

	// UserFindById

	// UserUpdate

	// UserDelete

	//reqEntity := &pb.User{}
	reqEntity := RandomUser()
	// { insert data
	log.Printf("开始插入测试数据 %+v \n", reqEntity)
	created, err := api.UserCreateNew(reqEntity).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("插入测试数据成功")
	// }

	// { find data by id
	log.Printf("开始根据 id %+v 查询上一步插入的数据", created)
	findedData, err := api.UserFindByIdNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("根据id 查询数据成功")
	// }

	toCmp := reqEntity
	toCmp.UserId = created.UserId

	{

		toCmp.UserId = created.UserId
		var cmpFunc func() bool

		cmpFunc = func() bool {
			return toCmp.UserId == findedData.UserId
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "UserId", toCmp.UserId, findedData.UserId)
			log.Println(err)
			return err
		}

		cmpFunc = func() bool {
			return toCmp.Phone.GetValue() == findedData.Phone.GetValue()
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "Phone", toCmp.Phone, findedData.Phone)
			log.Println(err)
			return err
		}

		cmpFunc = func() bool {
			return toCmp.Race == findedData.Race
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "Race", toCmp.Race, findedData.Race)
			log.Println(err)
			return err
		}

	}

	// { update
	reqEntityToUpdate := RandomUser()
	reqEntityToUpdate.UserId = created.UserId
	log.Printf("开始更新 数据 %+v \n", reqEntityToUpdate)
	_, err = api.UserUpdateNew(reqEntityToUpdate).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	// }

	// { check data after update
	log.Println("😈 check data after update")
	findedData, err = api.UserFindByIdNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	toCmp = reqEntityToUpdate

	{

		toCmp.UserId = created.UserId
		var cmpFunc func() bool

		cmpFunc = func() bool {
			return toCmp.UserId == findedData.UserId
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "UserId", toCmp.UserId, findedData.UserId)
			log.Println(err)
			return err
		}

		cmpFunc = func() bool {
			return toCmp.Phone.GetValue() == findedData.Phone.GetValue()
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "Phone", toCmp.Phone, findedData.Phone)
			log.Println(err)
			return err
		}

		cmpFunc = func() bool {
			return toCmp.Race == findedData.Race
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "Race", toCmp.Race, findedData.Race)
			log.Println(err)
			return err
		}

	}
	// }

	// { delete data
	_, err = api.UserDeleteNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	// }

	// { check deleted data TODO
	log.Println("Checking data if delete")
	findedData, err = api.UserFindByIdNew(created).ExecBy(ctx, this.client)
	if err == nil {
		log.Println("deleted data failed")
	}

	// }

	return nil
}

func RandomUser() *pb.User {
	output := &pb.User{}

	output.UserId = randomStringValue()

	// handle field Phone
	output.Phone = wrapperspb.String(randomStringValue())

	// handle field Race

	lengthrace := (3)
	randrace := rand.Int31n(int32(lengthrace))
	enumsrace := []string{"YELLOW", "WHITE", "BLACK"}

	output.Race = pb.User_Race(pb.User_Race_value[enumsrace[randrace]])
	return output
}
