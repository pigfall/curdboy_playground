package http_api_contacts

import (
	"context"
	"fmt"
	"github.com/pigfall/tzzGoUtil/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http/httputil"
	"strings"
)

type ApiClient struct {
	ApiBase string
}

type ApiClientIfce interface {
	SendReq(ctx context.Context, httpReqBuilder *http.RequestBuilder, apiPath string, responseObj interface{}) (httpBodyByte []byte, err error)
}

func ApiClientNew(apiBasePath string) ApiClientIfce {
	return &ApiClient{
		ApiBase: apiBasePath,
	}
}

func (this *ApiClient) SendReq(ctx context.Context, httpReqBuilder *http.RequestBuilder, apiPath string, responseObj interface{}) (httpBodyBytes []byte, err error) {
	fullApiPath := fmt.Sprintf("%s/%s", this.ApiBase, strings.TrimLeft(apiPath, "/"))
	httpReq, err := httpReqBuilder.URL(fullApiPath).Build(ctx)
	if err != nil {
		return nil, err
	}
	bytes, err := httputil.DumpRequest(httpReq, true)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bytes))

	unmarshalOption := protojson.UnmarshalOptions{}

	unmarshalWrapper := func(data []byte, v interface{}) error {
		msg := v.(proto.Message)
		return unmarshalOption.Unmarshal(data, msg)
	}
	resBodyBytes, err := http.DoRequestThenJsonUnMarshalAntReturnResBodyDataWithUnmarshal(ctx, unmarshalWrapper, httpReq, responseObj, nil, false)
	if err != nil {
		return resBodyBytes, err
	}
	return resBodyBytes, nil
}

func jsonMarshal(v interface{}) ([]byte, error) {
	marshalOption := protojson.MarshalOptions{}
	msg := v.(proto.Message)
	return marshalOption.Marshal(msg)
}

func jsonUnmarshal(dataBytes []byte, v interface{}) error {
	unmarshalOptions := protojson.UnmarshalOptions{}
	pbMsg := v.(proto.Message)
	return unmarshalOptions.Unmarshal(dataBytes, pbMsg)
}
