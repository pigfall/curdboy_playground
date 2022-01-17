package http_api_contacts

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/pigfall/curdboy_playground/api/v1/contacts"
	"github.com/pigfall/tzzGoUtil/http"
	rf "github.com/pigfall/tzzGoUtil/reflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"reflect"
	"strings"
	// "google.golang.org/protobuf/encoding/protojson"
	// "google.golang.org/protobuf/proto"
	"github.com/go-kratos/kratos/v2/errors"
)

type CarCreate struct {
	reqEntity *pb.Car
}

func CarCreateNew(entity *pb.Car) *CarCreate {
	return &CarCreate{
		reqEntity: entity,
	}
}

func (this *CarCreate) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.CarId, error) {
	res := &pb.CarId{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/cars"

	resBodyBytes, err := cli.SendReq(ctx, reqBuilder, apiPath, res)
	if err != nil {
		// try parse error
		errEntity := &errors.Error{}
		errForDecode := jsonUnmarshal(resBodyBytes, errEntity)
		if errForDecode == nil {
			return nil, errEntity
		}
		//
		err = fmt.Errorf("%w %s", err, string(resBodyBytes))
	}
	return res, err
}

type CarUpdate struct {
	reqEntity *pb.Car
}

func CarUpdateNew(entity *pb.Car) *CarUpdate {
	return &CarUpdate{
		reqEntity: entity,
	}
}

func (this *CarUpdate) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPatch()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/cars/{car_id}"

	// [car_id]

	if len(this.reqEntity.CarId) == 0 {
		panic(fmt.Errorf("car_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{car_id}", this.reqEntity.CarId)
	resBodyBytes, err := cli.SendReq(ctx, reqBuilder, apiPath, res)
	if err != nil {
		// try parse error
		errEntity := &errors.Error{}
		errForDecode := jsonUnmarshal(resBodyBytes, errEntity)
		if errForDecode == nil {
			return nil, errEntity
		}
		//
		err = fmt.Errorf("%w %s", err, string(resBodyBytes))
	}
	return res, err
}

type CarDelete struct {
	reqEntity *pb.CarId
}

func CarDeleteNew(entity *pb.CarId) *CarDelete {
	return &CarDelete{
		reqEntity: entity,
	}
}

func (this *CarDelete) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodDelete()
	var apiPath = "/cars/{car_id}"

	// [car_id]

	if len(this.reqEntity.CarId) == 0 {
		panic(fmt.Errorf("car_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{car_id}", this.reqEntity.CarId)
	resBodyBytes, err := cli.SendReq(ctx, reqBuilder, apiPath, res)
	if err != nil {
		// try parse error
		errEntity := &errors.Error{}
		errForDecode := jsonUnmarshal(resBodyBytes, errEntity)
		if errForDecode == nil {
			return nil, errEntity
		}
		//
		err = fmt.Errorf("%w %s", err, string(resBodyBytes))
	}
	return res, err
}

type CarFindById struct {
	reqEntity *pb.CarId
}

func CarFindByIdNew(entity *pb.CarId) *CarFindById {
	return &CarFindById{
		reqEntity: entity,
	}
}

func (this *CarFindById) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.Car, error) {
	res := &pb.Car{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodGet()
	urlQueryParamMap := make(map[string]string)
	var reqEntity = this.reqEntity
	rtReqEnt := reflect.TypeOf(reqEntity).Elem()
	jsonFieldToStuctField := map[string]reflect.StructField{}
	for i := 0; i < rtReqEnt.NumField(); i++ {
		jsonValue := rtReqEnt.Field(i).Tag.Get("json")
		if len(jsonValue) > 0 {
			jsonFieldToStuctField[strings.Split(jsonValue, ",")[0]] = rtReqEnt.Field(i)
		}
	}
	reqEntityJSONMap := make(map[string]interface{})
	bytes, err := jsonMarshal(reqEntity)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &reqEntityJSONMap)
	if err != nil {
		return nil, err
	}
	//for k,v := range reqEntityJSONMap{
	//	urlQueryParamMap [k]= strings.TrimSuffix(rf.ToString(reflect.ValueOf(v)),".00")
	//}
	for k, v := range reqEntityJSONMap {
		if f := jsonFieldToStuctField[k]; len(f.Name) > 0 {
			if !(f.Type.Kind() == reflect.String || f.Type.Kind() == reflect.Float32 || f.Type.Kind() == reflect.Float64 || (f.Type.Kind() == reflect.Ptr && (f.Type.Elem().Name() == "DoubleValue" || f.Type.Elem().Name() == "StringValue"))) {
				urlQueryParamMap[k] = strings.TrimSuffix(rf.ToString(reflect.ValueOf(v)), ".00")
			} else {
				urlQueryParamMap[k] = rf.ToString(reflect.ValueOf(v))
			}
		} else {
			panic(fmt.Errorf("Unreachable TODO 详细说明"))
		}
	}
	reqBuilder.PutParamsToUrl(urlQueryParamMap)
	var apiPath = "/cars/{car_id}"

	// [car_id]

	if len(this.reqEntity.CarId) == 0 {
		panic(fmt.Errorf("car_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{car_id}", this.reqEntity.CarId)
	resBodyBytes, err := cli.SendReq(ctx, reqBuilder, apiPath, res)
	if err != nil {
		// try parse error
		errEntity := &errors.Error{}
		errForDecode := jsonUnmarshal(resBodyBytes, errEntity)
		if errForDecode == nil {
			return nil, errEntity
		}
		//
		err = fmt.Errorf("%w %s", err, string(resBodyBytes))
	}
	return res, err
}

type CarFind struct {
	reqEntity *pb.CarPageQuery
}

func CarFindNew(entity *pb.CarPageQuery) *CarFind {
	return &CarFind{
		reqEntity: entity,
	}
}

func (this *CarFind) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.CarWithCount, error) {
	res := &pb.CarWithCount{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodGet()
	urlQueryParamMap := make(map[string]string)
	var reqEntity = this.reqEntity
	rtReqEnt := reflect.TypeOf(reqEntity).Elem()
	jsonFieldToStuctField := map[string]reflect.StructField{}
	for i := 0; i < rtReqEnt.NumField(); i++ {
		jsonValue := rtReqEnt.Field(i).Tag.Get("json")
		if len(jsonValue) > 0 {
			jsonFieldToStuctField[strings.Split(jsonValue, ",")[0]] = rtReqEnt.Field(i)
		}
	}
	reqEntityJSONMap := make(map[string]interface{})
	bytes, err := jsonMarshal(reqEntity)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &reqEntityJSONMap)
	if err != nil {
		return nil, err
	}
	//for k,v := range reqEntityJSONMap{
	//	urlQueryParamMap [k]= strings.TrimSuffix(rf.ToString(reflect.ValueOf(v)),".00")
	//}
	for k, v := range reqEntityJSONMap {
		if f := jsonFieldToStuctField[k]; len(f.Name) > 0 {
			if !(f.Type.Kind() == reflect.String || f.Type.Kind() == reflect.Float32 || f.Type.Kind() == reflect.Float64 || (f.Type.Kind() == reflect.Ptr && (f.Type.Elem().Name() == "DoubleValue" || f.Type.Elem().Name() == "StringValue"))) {
				urlQueryParamMap[k] = strings.TrimSuffix(rf.ToString(reflect.ValueOf(v)), ".00")
			} else {
				urlQueryParamMap[k] = rf.ToString(reflect.ValueOf(v))
			}
		} else {
			panic(fmt.Errorf("Unreachable TODO 详细说明"))
		}
	}
	reqBuilder.PutParamsToUrl(urlQueryParamMap)
	var apiPath = "/cars"

	resBodyBytes, err := cli.SendReq(ctx, reqBuilder, apiPath, res)
	if err != nil {
		// try parse error
		errEntity := &errors.Error{}
		errForDecode := jsonUnmarshal(resBodyBytes, errEntity)
		if errForDecode == nil {
			return nil, errEntity
		}
		//
		err = fmt.Errorf("%w %s", err, string(resBodyBytes))
	}
	return res, err
}

type CarCount struct {
	reqEntity *pb.Car
}

func CarCountNew(entity *pb.Car) *CarCount {
	return &CarCount{
		reqEntity: entity,
	}
}

func (this *CarCount) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.CountNumber, error) {
	res := &pb.CountNumber{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodGet()
	urlQueryParamMap := make(map[string]string)
	var reqEntity = this.reqEntity
	rtReqEnt := reflect.TypeOf(reqEntity).Elem()
	jsonFieldToStuctField := map[string]reflect.StructField{}
	for i := 0; i < rtReqEnt.NumField(); i++ {
		jsonValue := rtReqEnt.Field(i).Tag.Get("json")
		if len(jsonValue) > 0 {
			jsonFieldToStuctField[strings.Split(jsonValue, ",")[0]] = rtReqEnt.Field(i)
		}
	}
	reqEntityJSONMap := make(map[string]interface{})
	bytes, err := jsonMarshal(reqEntity)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &reqEntityJSONMap)
	if err != nil {
		return nil, err
	}
	//for k,v := range reqEntityJSONMap{
	//	urlQueryParamMap [k]= strings.TrimSuffix(rf.ToString(reflect.ValueOf(v)),".00")
	//}
	for k, v := range reqEntityJSONMap {
		if f := jsonFieldToStuctField[k]; len(f.Name) > 0 {
			if !(f.Type.Kind() == reflect.String || f.Type.Kind() == reflect.Float32 || f.Type.Kind() == reflect.Float64 || (f.Type.Kind() == reflect.Ptr && (f.Type.Elem().Name() == "DoubleValue" || f.Type.Elem().Name() == "StringValue"))) {
				urlQueryParamMap[k] = strings.TrimSuffix(rf.ToString(reflect.ValueOf(v)), ".00")
			} else {
				urlQueryParamMap[k] = rf.ToString(reflect.ValueOf(v))
			}
		} else {
			panic(fmt.Errorf("Unreachable TODO 详细说明"))
		}
	}
	reqBuilder.PutParamsToUrl(urlQueryParamMap)
	var apiPath = "/cars/cars/count"

	resBodyBytes, err := cli.SendReq(ctx, reqBuilder, apiPath, res)
	if err != nil {
		// try parse error
		errEntity := &errors.Error{}
		errForDecode := jsonUnmarshal(resBodyBytes, errEntity)
		if errForDecode == nil {
			return nil, errEntity
		}
		//
		err = fmt.Errorf("%w %s", err, string(resBodyBytes))
	}
	return res, err
}

func ForEachCar(ctx context.Context, apiCli ApiClientIfce, do func(ctx context.Context, entity *pb.Car) error) error {
	var pageIndex = 0
	for {
		res, err := CarFindNew(&pb.CarPageQuery{
			PageIndex: int32(pageIndex),
			PageSize:  int32(10),
		}).ExecBy(ctx, apiCli)
		if err != nil {
			return err
		}

		if len(res.Cars) == 0 {
			return nil
		}
		pageIndex++
		for _, entity := range res.Cars {
			err := do(ctx, entity)
			if err != nil {
				return err
			}
		}
	}
}
