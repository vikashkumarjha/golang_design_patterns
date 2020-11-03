package singleton

import "testing"


func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Errorf("expected pointer of instance not nil")
	}
}

func TestSingleton_AddOne(t *testing.T) {

	inst := GetInstance()
	inst.AddOne()
	inst.AddOne()

	inst = GetInstance()

	inst.AddOne()
	inst.AddOne()

	expected := 4
	if inst.GetCount() != expected {
		t.Errorf("expected (%v) does not match actual(%v) ", expected, inst.GetCount())
	}


}




