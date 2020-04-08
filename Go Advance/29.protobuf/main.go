package main

import (
	// sesuaikan dengan struktuk folder projek masing2
	"bytes"
	"fmt"
	"os"

	"../29.protobuf/model"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	// more code here ...
	var user1 = &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   model.UserGender_FEMALE,
	}

	var userList = &model.UserList{
		List: []*model.User{
			user1,
		},
	}

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	var garageListByUser = &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	jsonString := objProtoToJSON(garageList)
	userListJSON := objProtoToJSON(userList)
	garageListByUserJSON := objProtoToJSON(garageListByUser)

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)
	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", user1.String())
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)
	fmt.Printf("# ==== As JSON String\n       %v \n", userListJSON)
	fmt.Printf("# ==== As JSON String\n       %v \n", garageListByUserJSON)
}

func objProtoToJSON(obj proto.Message) string {
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, obj)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(0)
	}
	jsonString := buf.String()
	return jsonString
}
