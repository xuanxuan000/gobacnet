package gobacnet

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	// list, _ := Whois("devid", "111")
	list, _ := Whois("", "")
	fmt.Printf("\nlist:%+v\n", list)

	// para: devid, objType, objID, prop, index
	readp, err := Readprop(192000, OBJECT_ANALOG_INPUT, 1, PROP_PRESENT_VALUE, -1) // -1表示非array
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadp:%+v\n", readp)
	}
	// readp, err := Readprop(111, OBJECT_ANALOG_INPUT, 5, PROP_PRESENT_VALUE, -1) // -1表示非array
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("\nreadp:%+v\n", readp)
	// }

	// para: devid, objType, objID, prop, priority, index, tag, value
	err = Writeprop(192000, OBJECT_BINARY_VALUE, 5,
		PROP_PRESENT_VALUE, 16, -1, BACNET_APPLICATION_TAG_SIGNED_INT, "34")
	if err != nil {
		fmt.Println(err)
	}

	// para: devid, objType, objID, prop, index
	tmp1 := ReadM_para{OBJECT_ANALOG_INPUT, 1, PROP_PRESENT_VALUE}
	tmp2 := ReadM_para{OBJECT_MULTI_STATE_VALUE, 2, PROP_PRESENT_VALUE}
	var parasli []ReadM_para
	parasli = append(parasli, tmp1)
	parasli = append(parasli, tmp2)
	readp, err = ReadpropM(192000, parasli)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadpm:%+v\n", readp)
	}

	tmp3 := WriteM_para{OBJECT_ANALOG_INPUT, 5, PROP_PRESENT_VALUE, 16, BACNET_APPLICATION_TAG_SIGNED_INT, "12"}
	tmp4 := WriteM_para{OBJECT_ANALOG_INPUT, 5, PROP_PRESENT_VALUE, 16, BACNET_APPLICATION_TAG_SIGNED_INT, "12"}
	var parasli2 []WriteM_para
	parasli2 = append(parasli2, tmp3)
	parasli2 = append(parasli2, tmp4)
	err = WritepropM(111, parasli2)
	if err != nil {
		fmt.Println(err)
	}
}
