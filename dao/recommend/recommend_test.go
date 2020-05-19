package recommendDao

import "testing"

func TestGetAllMovieCount(t *testing.T) {
	data, err := GetAllPostRecord()
	if err != nil {
		t.Error("err")
	}

	t.Log(data)
}


func TestGetUserPostRecord(t *testing.T) {
	data, err := GetUserPostRecord(1,103539870875648)
	if err != nil {
		t.Error("err")
	}

	t.Log(data)
}
