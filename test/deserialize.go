package test

import (
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/httpx"
	"io"
	"testing"
)

func Deserialize[E ErrorHandler, T any](gotBody, wantBody io.Reader, t *testing.T) (gotT, wantT T, success bool) {
	var e E

	gotStatus := aspect.StatusOK()
	gotT, gotStatus = httpx.Content[T](gotBody)
	if !gotStatus.OK() && !gotStatus.NoContent() {
		//t.Errorf("Deserialize() %v err = %v", "got", gotStatus.Err)
		e.Handle(gotStatus, t, "got")
		return
	}

	wantStatus := aspect.StatusOK()
	wantT, wantStatus = httpx.Content[T](wantBody)
	if !wantStatus.OK() && !wantStatus.NoContent() {
		//t.Errorf("Deserialize() %v err = %v", "want", wantStatus.Err)
		e.Handle(wantStatus, t, "want")
		return
	}

	if gotStatus.Code != wantStatus.Code {
		t.Errorf("Deserialize() got status code = %v, want status code = %v", gotStatus.Code, wantStatus.Code)
		return
	}
	return gotT, wantT, true
}
