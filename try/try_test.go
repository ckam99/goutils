package try

import (
	"testing"
)

const (
  trycatmsg = "I tried: try/catch"
  catchtrymsg = "I tried: catch/try"
  errorInfo = "error occured!"
)

func TestTryCatch(t testing.T) {
  var msg string
  var err error
  
  Do(func(){
    msg = trycatmsg
    
    panic("error occured!")
  }).Catch(func(e error) {
       err = e
  })
  if msg != trycatmsg {
    t.Errorf("expected: %s, got: %s", trycatmsg, msg)
  }
  if err.Error() != errorInfo {
    t.Errorf("expected: %s, got: %s", errorInfo, err.Error())
  }
}

func TestCatchTry(t testing.T) {
  var msg string
  var err error
  
  Catch(func(e error) {
       err = e
  }).Do(func(){
    msg = catchtrymsg
    panic("error occured!")
  })
  if msg != catchtrymsg {
    t.Errorf("expected: %s, got: %s", catchtrymsg, msg)
  }
  if err.Error() != errorInfo {
    t.Errorf("expected: %s, got: %s", errorInfo, err.Error())
  }
}
