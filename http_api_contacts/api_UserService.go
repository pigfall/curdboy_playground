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

type UserCreate struct {
	reqEntity *pb.User
}

func UserCreateNew(entity *pb.User) *UserCreate {
	return &UserCreate{
		reqEntity: entity,
	}
}

func (this *UserCreate) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.UserId, error) {
	res := &pb.UserId{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/users"

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

type UserUpdate struct {
	reqEntity *pb.User
}

func UserUpdateNew(entity *pb.User) *UserUpdate {
	return &UserUpdate{
		reqEntity: entity,
	}
}

func (this *UserUpdate) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPatch()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/users/{user_id}"

	// [user_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)
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

type UserDelete struct {
	reqEntity *pb.UserId
}

func UserDeleteNew(entity *pb.UserId) *UserDelete {
	return &UserDelete{
		reqEntity: entity,
	}
}

func (this *UserDelete) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodDelete()
	var apiPath = "/users/{user_id}"

	// [user_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)
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

type UserFindById struct {
	reqEntity *pb.UserId
}

func UserFindByIdNew(entity *pb.UserId) *UserFindById {
	return &UserFindById{
		reqEntity: entity,
	}
}

func (this *UserFindById) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.User, error) {
	res := &pb.User{}
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
	var apiPath = "/users/{user_id}"

	// [user_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)
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

type UserFind struct {
	reqEntity *pb.UserPageQuery
}

func UserFindNew(entity *pb.UserPageQuery) *UserFind {
	return &UserFind{
		reqEntity: entity,
	}
}

func (this *UserFind) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.UserWithCount, error) {
	res := &pb.UserWithCount{}
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
	var apiPath = "/users"

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

type UserCount struct {
	reqEntity *pb.User
}

func UserCountNew(entity *pb.User) *UserCount {
	return &UserCount{
		reqEntity: entity,
	}
}

func (this *UserCount) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.CountNumber, error) {
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
	var apiPath = "/users/users/count"

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

type UserAddCar struct {
	reqEntity *pb.Car
}

func UserAddCarNew(entity *pb.Car) *UserAddCar {
	return &UserAddCar{
		reqEntity: entity,
	}
}

func (this *UserAddCar) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.CarId, error) {
	res := &pb.CarId{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/users/{user_id}/cars"

	// [user_id]

	if len(this.reqEntity.UserId.GetValue()) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId.GetValue())
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

type UserAddCarById struct {
	reqEntity *pb.UserIdAndCarId
}

func UserAddCarByIdNew(entity *pb.UserIdAndCarId) *UserAddCarById {
	return &UserAddCarById{
		reqEntity: entity,
	}
}

func (this *UserAddCarById) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/users/{user_id}/cars/{car_id}"

	// [user_id car_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)

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

type UserRemoveCar struct {
	reqEntity *pb.UserRemoveCarReq
}

func UserRemoveCarNew(entity *pb.UserRemoveCarReq) *UserRemoveCar {
	return &UserRemoveCar{
		reqEntity: entity,
	}
}

func (this *UserRemoveCar) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodDelete()
	var apiPath = "/users/{user_id}/cars"

	// [user_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)
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

type UserFindCars struct {
	reqEntity *pb.UserFindCarsReq
}

func UserFindCarsNew(entity *pb.UserFindCarsReq) *UserFindCars {
	return &UserFindCars{
		reqEntity: entity,
	}
}

func (this *UserFindCars) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.UserFindCarsRes, error) {
	res := &pb.UserFindCarsRes{}
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
	var apiPath = "/users/{user_id}/cars"

	// [user_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)
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

type UserFindDepts struct {
	reqEntity *pb.UserFindDeptsReq
}

func UserFindDeptsNew(entity *pb.UserFindDeptsReq) *UserFindDepts {
	return &UserFindDepts{
		reqEntity: entity,
	}
}

func (this *UserFindDepts) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.UserFindDeptsRes, error) {
	res := &pb.UserFindDeptsRes{}
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
	var apiPath = "/users/{user_id}/depts"

	// [user_id]

	if len(this.reqEntity.UserId) == 0 {
		panic(fmt.Errorf("user_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{user_id}", this.reqEntity.UserId)
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

func ForEachUser(ctx context.Context, apiCli ApiClientIfce, do func(ctx context.Context, entity *pb.User) error) error {
	var pageIndex = 0
	for {
		res, err := UserFindNew(&pb.UserPageQuery{
			PageIndex: int32(pageIndex),
			PageSize:  int32(10),
		}).ExecBy(ctx, apiCli)
		if err != nil {
			return err
		}

		if len(res.Users) == 0 {
			return nil
		}
		pageIndex++
		for _, entity := range res.Users {
			err := do(ctx, entity)
			if err != nil {
				return err
			}
		}
	}
}
