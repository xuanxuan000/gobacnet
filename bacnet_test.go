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
	readp, err := Readprop(200000, OBJECT_BINARY_VALUE, 1, PROP_PRESENT_VALUE, -1) // -1表示非array
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadp:%+v\n", readp)
	}
	// para: devid, objType, objID, prop, priority, index, tag, value
	err = Writeprop(200000, OBJECT_BINARY_VALUE, 1,
		PROP_PRESENT_VALUE, 16, -1, BACNET_APPLICATION_TAG_ENUMERATED, "1")
	if err != nil {
		fmt.Println(err)
	}
	readp, err = Readprop(200000, OBJECT_BINARY_VALUE, 1, PROP_PRESENT_VALUE, -1) // -1表示非array
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadp:%+v\n", readp)
	}

	// para: devid, objType, objID, prop, index
	tmp0 := ReadM_para{OBJECT_ANALOG_INPUT, 1, PROP_PRESENT_VALUE}
	tmp1 := ReadM_para{OBJECT_BINARY_VALUE, 1, PROP_PRESENT_VALUE}
	tmp2 := ReadM_para{OBJECT_MULTI_STATE_VALUE, 1, PROP_PRESENT_VALUE}
	tmp3 := ReadM_para{OBJECT_ANALOG_VALUE, 1, PROP_PRESENT_VALUE}
	tmp4 := ReadM_para{OBJECT_MULTI_STATE_VALUE, 2, PROP_PRESENT_VALUE}
	var parasli []ReadM_para
	parasli = append(parasli, tmp0)
	parasli = append(parasli, tmp1)
	parasli = append(parasli, tmp2)
	parasli = append(parasli, tmp3)
	parasli = append(parasli, tmp4)
	readp, err = ReadpropM(200000, parasli)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadpm:%+v\n", readp)
	}

	w_tmp0 := WriteM_para{OBJECT_BINARY_VALUE, 1, PROP_PRESENT_VALUE, 16, BACNET_APPLICATION_TAG_ENUMERATED, "12"}
	w_tmp1 := WriteM_para{OBJECT_MULTI_STATE_VALUE, 1, PROP_PRESENT_VALUE, 16, BACNET_APPLICATION_TAG_UNSIGNED_INT, "12"}
	w_tmp2 := WriteM_para{OBJECT_ANALOG_VALUE, 1, PROP_PRESENT_VALUE, 16, BACNET_APPLICATION_TAG_REAL, "12"}
	w_tmp3 := WriteM_para{OBJECT_MULTI_STATE_VALUE, 2, PROP_PRESENT_VALUE, 16, BACNET_APPLICATION_TAG_UNSIGNED_INT, "12"}
	var parasli2 []WriteM_para
	parasli2 = append(parasli2, w_tmp0)
	parasli2 = append(parasli2, w_tmp1)
	parasli2 = append(parasli2, w_tmp2)
	parasli2 = append(parasli2, w_tmp3)
	err = WritepropM(200000, parasli2)
	if err != nil {
		fmt.Println(err)
	}

	readp, err = ReadpropM(200000, parasli)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadpm:%+v\n", readp)
	}
}
