package main

func returnString() interface{} {
	return "string"
}

func returnInt() interface{} {
	return 100
}

type testType string

func returnTestType() interface{} {
	var a testType
	a = "this is test type"
	return a
}

type testStruct struct {
	ID   int
	Name string
}

type anotherTestStruct testStruct

func returnAnotherTestStruct() interface{} {
	return testStruct{
		ID:   100,
		Name: "struct name",
	}
}
