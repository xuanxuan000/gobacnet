package gobacnet

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	list, _ := Whois("devid", "111")
	// list, _ := Whois("", "")
	fmt.Printf("\nlist:%+v\n", list)

	// para: devid, obj_tag, obj_index,
	readp, err := Readprop("111", OBJECT_ANALOG_OUTPUT, "101", "85", "1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nreadp:%+v\n", readp)
	}

	err = Writeprop("111", "1", "101", "85", "16", "-1", "3", "34")
	if err != nil {
		fmt.Println(err)
	}
}
