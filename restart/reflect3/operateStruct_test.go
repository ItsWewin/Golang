package operateStruct

import "testing"

func TestReflectUsed(t *testing.T) {
	user1 := user{"wewin", 26}
	if age := reflectUsed(user1); age == 26 {
		t.Log("ok")
	} else {
		t.Errorf("expect: %v, get %v", 26, age)
	}
}
