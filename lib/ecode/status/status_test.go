package status

import (
	"com.youyu.api/lib/ecode"
	"github.com/pkg/errors"
	"testing"
)

func TestStatus(t *testing.T) {
	err := Error(ecode.UserDuplicate, ecode.UserDuplicate.Message())
	if err != nil {
		t.Log(FromError(err))
	}
	err = errors.New("hello world")
	_,bl := FromError(err)
	if bl {
		t.Log(err)
	}
	st,bl := FromError(nil)
	if bl {
		t.Log(st.Code)
	}
}
