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

type DeptCreate struct {
	reqEntity *pb.Dept
}

func DeptCreateNew(entity *pb.Dept) *DeptCreate {
	return &DeptCreate{
		reqEntity: entity,
	}
}

func (this *DeptCreate) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.DeptId, error) {
	res := &pb.DeptId{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/depts"

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

type DeptUpdate struct {
	reqEntity *pb.Dept
}

func DeptUpdateNew(entity *pb.Dept) *DeptUpdate {
	return &DeptUpdate{
		reqEntity: entity,
	}
}

func (this *DeptUpdate) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPatch()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/depts/{dept_id}"

	// [dept_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)
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

type DeptDelete struct {
	reqEntity *pb.DeptId
}

func DeptDeleteNew(entity *pb.DeptId) *DeptDelete {
	return &DeptDelete{
		reqEntity: entity,
	}
}

func (this *DeptDelete) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodDelete()
	var apiPath = "/depts/{dept_id}"

	// [dept_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)
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

type DeptFindById struct {
	reqEntity *pb.DeptId
}

func DeptFindByIdNew(entity *pb.DeptId) *DeptFindById {
	return &DeptFindById{
		reqEntity: entity,
	}
}

func (this *DeptFindById) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.Dept, error) {
	res := &pb.Dept{}
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
	var apiPath = "/depts/{dept_id}"

	// [dept_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)
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

type DeptFind struct {
	reqEntity *pb.DeptPageQuery
}

func DeptFindNew(entity *pb.DeptPageQuery) *DeptFind {
	return &DeptFind{
		reqEntity: entity,
	}
}

func (this *DeptFind) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.DeptWithCount, error) {
	res := &pb.DeptWithCount{}
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
	var apiPath = "/depts"

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

type DeptCount struct {
	reqEntity *pb.Dept
}

func DeptCountNew(entity *pb.Dept) *DeptCount {
	return &DeptCount{
		reqEntity: entity,
	}
}

func (this *DeptCount) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.CountNumber, error) {
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
	var apiPath = "/depts/depts/count"

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

type DeptAddUser struct {
	reqEntity *pb.DeptAddUserReq
}

func DeptAddUserNew(entity *pb.DeptAddUserReq) *DeptAddUser {
	return &DeptAddUser{
		reqEntity: entity,
	}
}

func (this *DeptAddUser) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.UserId, error) {
	res := &pb.UserId{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/depts/{dept_id}/users"

	// [dept_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)
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

type DeptAddUserById struct {
	reqEntity *pb.DeptIdAndUserId
}

func DeptAddUserByIdNew(entity *pb.DeptIdAndUserId) *DeptAddUserById {
	return &DeptAddUserById{
		reqEntity: entity,
	}
}

func (this *DeptAddUserById) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodPost()
	reqBuilder.JsonParamWithMarshalFunc(this.reqEntity, jsonMarshal)
	var apiPath = "/depts/{dept_id}/users/{user_id}"

	// [dept_id user_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)

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

type DeptRemoveUser struct {
	reqEntity *pb.DeptRemoveUserReq
}

func DeptRemoveUserNew(entity *pb.DeptRemoveUserReq) *DeptRemoveUser {
	return &DeptRemoveUser{
		reqEntity: entity,
	}
}

func (this *DeptRemoveUser) ExecBy(ctx context.Context, cli ApiClientIfce) (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	reqBuilder := http.NewRequestBuilder()
	var err error
	reqBuilder.MethodDelete()
	var apiPath = "/depts/{dept_id}/users"

	// [dept_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)
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

type DeptFindUsers struct {
	reqEntity *pb.DeptFindUsersReq
}

func DeptFindUsersNew(entity *pb.DeptFindUsersReq) *DeptFindUsers {
	return &DeptFindUsers{
		reqEntity: entity,
	}
}

func (this *DeptFindUsers) ExecBy(ctx context.Context, cli ApiClientIfce) (*pb.DeptFindUsersRes, error) {
	res := &pb.DeptFindUsersRes{}
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
	var apiPath = "/depts/{dept_id}/users"

	// [dept_id]

	if len(this.reqEntity.DeptId) == 0 {
		panic(fmt.Errorf("dept_id can not be nil"))
	}
	apiPath = strings.ReplaceAll(apiPath, "{dept_id}", this.reqEntity.DeptId)
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

func ForEachDept(ctx context.Context, apiCli ApiClientIfce, do func(ctx context.Context, entity *pb.Dept) error) error {
	var pageIndex = 0
	for {
		res, err := DeptFindNew(&pb.DeptPageQuery{
			PageIndex: int32(pageIndex),
			PageSize:  int32(10),
		}).ExecBy(ctx, apiCli)
		if err != nil {
			return err
		}

		if len(res.Depts) == 0 {
			return nil
		}
		pageIndex++
		for _, entity := range res.Depts {
			err := do(ctx, entity)
			if err != nil {
				return err
			}
		}
	}
}
