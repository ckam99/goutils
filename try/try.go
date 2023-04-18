package try

import "fmt"


type Try struct {
	catch func()
  do func()
}

func Catch(f func(ex error)) *Try {
  return &Try{
    catch: func() {
		if r := recover(); r != nil {
			f(fmt.Errorf("%v", r))
		}
	 },
  }
}

func(t *Try) Do(f func()) *Try {
  if t.catch != nil {
    defer t.catch()
    f()
  }
	return t
}


func Do(f func()) *Try {
  return &Try{
    do: f,
  }
}

func(t *Try) Catch(f func(ex error)) *Try {
  t.catch = func() {
  		if r := recover(); r != nil {
  			f(fmt.Errorf("%v", r))
  		}
	}
  defer t.catch()
  
  if t.do != nil {
    t.do()
  }
	return t
}
