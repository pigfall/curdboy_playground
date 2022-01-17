package api_inquisitor

import (
	"context"
	"fmt"
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	api "github.com/pigfall/curdboy_playground/http_api_contacts"
	"log"
)

type DeptServiceInquisitor struct {
	client api.ApiClientIfce
}

func DeptServiceInquisitorNew(client api.ApiClientIfce) ApiInquisitor {
	return &DeptServiceInquisitor{
		client: client,
	}
}

func (this *DeptServiceInquisitor) Check(ctx context.Context) error {

	// DeptCreate

	// DeptFind

	// DeptFindById

	// DeptUpdate

	// DeptDelete

	//reqEntity := &pb.Dept{}
	reqEntity := RandomDept()
	// { insert data
	log.Printf("开始插入测试数据 %+v \n", reqEntity)
	created, err := api.DeptCreateNew(reqEntity).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("插入测试数据成功")
	// }

	// { find data by id
	log.Printf("开始根据 id %+v 查询上一步插入的数据", created)
	findedData, err := api.DeptFindByIdNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("根据id 查询数据成功")
	// }

	toCmp := reqEntity
	toCmp.DeptId = created.DeptId

	{

		toCmp.DeptId = created.DeptId
		var cmpFunc func() bool

		cmpFunc = func() bool {
			return toCmp.DeptId == findedData.DeptId
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "DeptId", toCmp.DeptId, findedData.DeptId)
			log.Println(err)
			return err
		}

	}

	// { update
	reqEntityToUpdate := RandomDept()
	reqEntityToUpdate.DeptId = created.DeptId
	log.Printf("开始更新 数据 %+v \n", reqEntityToUpdate)
	_, err = api.DeptUpdateNew(reqEntityToUpdate).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	// }

	// { check data after update
	log.Println("😈 check data after update")
	findedData, err = api.DeptFindByIdNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	toCmp = reqEntityToUpdate

	{

		toCmp.DeptId = created.DeptId
		var cmpFunc func() bool

		cmpFunc = func() bool {
			return toCmp.DeptId == findedData.DeptId
		}
		if !cmpFunc() {
			err = fmt.Errorf("Field %s not matched %v %v", "DeptId", toCmp.DeptId, findedData.DeptId)
			log.Println(err)
			return err
		}

	}
	// }

	// { delete data
	_, err = api.DeptDeleteNew(created).ExecBy(ctx, this.client)
	if err != nil {
		log.Println(err)
		return err
	}
	// }

	// { check deleted data TODO
	log.Println("Checking data if delete")
	findedData, err = api.DeptFindByIdNew(created).ExecBy(ctx, this.client)
	if err == nil {
		log.Println("deleted data failed")
	}

	// }

	return nil
}

func RandomDept() *pb.Dept {
	output := &pb.Dept{}

	output.DeptId = randomStringValue()
	return output
}
