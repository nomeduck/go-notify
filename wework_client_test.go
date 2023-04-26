//go:build integration
// +build integration

package notify_test

import (
	"github.com/pkg6/go-notify"
	"github.com/pkg6/go-notify/message/wework"
	"reflect"
	"testing"
)

func TestWeWorkClient_Send(t *testing.T) {
	client := &notify.WeWorkClient{
		Key: "90645cb6-86d9-4a3f-85d1-ea21a2afba2f",
	}
	n := notify.New(client)
	var sender map[string]notify.IResult
	text := wework.NewTextMessage("微信测试 text")
	sender = n.Send(text)
	result := sender[notify.Wework]
	wantResult := notify.WeWorkResponse{
		Errcode: 0,
		Errmsg:  "ok",
	}
	wantStatus := 0
	wantError := ""
	if !reflect.DeepEqual(result.Result(), wantResult) {
		t.Fatalf("result.Result \nwant: %v\n got: %v", wantResult, result.Result())
	}
	if result.Status() != wantStatus {
		t.Fatalf("result.Status \nwant: %v\n got: %v", wantStatus, result.Status())
	}
	if result.Error() != wantError {
		t.Fatalf("result.Error \nwant: %v\n got: %v", wantError, result.Error())
	}
}
