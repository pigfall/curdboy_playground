package api_inquisitor

import (
	"context"
	"fmt"
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	api "github.com/pigfall/curdboy_playground/http_api_contacts"
	"log"

	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type CarServiceInquisitor struct {
	client api.ApiClientIfce
}

func CarServiceInquisitorNew(client api.ApiClientIfce) ApiInquisitor {
	return &CarServiceInquisitor{
		client: client,
	}
}

func (this *CarServiceInquisitor) Check(ctx context.Context) error {

	// CarCreate

	// CarFind

	// CarFindById

	// CarUpdate

	// CarDelete

	//reqEntity := &pb.Car{}
	reqEntity := RandomCar()
	// { insert data
	log.Printf("开始插入测试数据 %+v \n", reqEntity)
	created, err := api.CarCreateNew(reqEntity).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("插入测试数据成功")
	// }

	// { find data by id
	log.Printf("开始根据 id %+v 查询上一步插入的数据", created)
	findedData, err := api.CarFindByIdNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("根据id 查询数据成功")
	// }

	toCmp := reqEntity
	toCmp.CarId = created.CarId

	{

		toCmp.CarId = created.CarId
		var cmpFunc func() bool

		cmpFunc = func() bool {
			return toCmp.CarId == findedData.CarId
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "CarId", toCmp.CarId, findedData.CarId)
			log.Println(err)
			return err
		}

		cmpFunc = func() bool {
			return toCmp.UserId.GetValue() == findedData.UserId.GetValue()
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "UserId", toCmp.UserId, findedData.UserId)
			log.Println(err)
			return err
		}

	}

	// { update
	reqEntityToUpdate := RandomCar()
	reqEntityToUpdate.CarId = created.CarId
	log.Printf("开始更新 数据 %+v \n", reqEntityToUpdate)
	_, err = api.CarUpdateNew(reqEntityToUpdate).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	// }

	// { check data after update
	log.Println("😈 check data after update")
	findedData, err = api.CarFindByIdNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	toCmp = reqEntityToUpdate

	{

		toCmp.CarId = created.CarId
		var cmpFunc func() bool

		cmpFunc = func() bool {
			return toCmp.CarId == findedData.CarId
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "CarId", toCmp.CarId, findedData.CarId)
			log.Println(err)
			return err
		}

		cmpFunc = func() bool {
			return toCmp.UserId.GetValue() == findedData.UserId.GetValue()
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "UserId", toCmp.UserId, findedData.UserId)
			log.Println(err)
			return err
		}

	}
	// }

	// { delete data
	_, err = api.CarDeleteNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	// }

	// { check deleted data TODO
	log.Println("Checking data if delete")
	findedData, err = api.CarFindByIdNew(created).ExecBy(ctx, this.client)
	if err == nil {
		log.Println("deleted data failed")
	}

	// }

	return nil
}

func RandomCar() *pb.Car {
	output := &pb.Car{}

	output.CarId = randomStringValue()

	// compatiable import path
	_ = wrapperspb.String("")
	output.UserId = nil
	// TODO handle edge UserId
	return output
}
