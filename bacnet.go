package gobacnet

//#cgo CFLAGS: -I./lib
//#cgo LDFLAGS: -L${SRCDIR}/lib -lbac
//
//#include "bac.h"
//#include "struct.h"
//#include "stdlib.h"
import "C"
import (
	"errors"
	"unsafe"
)

func Whois(Wtype string, Wvalue string) (res []Who, err error) {
	var list C.struct_Whois_list_array
	if Wtype == "" {
		list = C.Whois(nil, nil)
	} else {
		arg0 := C.CString(Wtype)
		arg1 := C.CString(Wvalue)
		defer func() {
			C.free(unsafe.Pointer(arg0))
			C.free(unsafe.Pointer(arg1))
		}()
		list = C.Whois(arg0, arg1)
	}
	res = CArrayToGoArray_Who(unsafe.Pointer(list.array), int(list.size))
	return
}

func Readprop(devID string, objType string, objID string, prop string, index string) (res string, err error) {
	var para []*C.char
	para = append(para, C.CString(devID), C.CString(objType), C.CString(objID), C.CString(prop), C.CString(index))
	defer func() {
		for i := 0; i < len(para); i++ {
			C.free(unsafe.Pointer(para[i]))
		}
	}()
	read_res := C.Readprop(para[0], para[1], para[2], para[3], para[4])
	goread := CArrayToGoArray_char(unsafe.Pointer(read_res), 32)
	res = string(goread)
	if res == "&err" {
		err = errors.New("readprop err")
	}
	return
}

func Writeprop(devID string, objType string, objID string, prop string,
	priority string, index string, tag string, value string) (err error) {
	var para []*C.char
	para = append(para, C.CString(devID), C.CString(objType), C.CString(objID),
		C.CString(prop), C.CString(priority), C.CString(index), C.CString(tag), C.CString(value))
	defer func() {
		for i := 0; i < len(para); i++ {
			C.free(unsafe.Pointer(para[i]))
		}
	}()
	res := C.Writeprop(para[0], para[1], para[2], para[3], para[4], para[5], para[6], para[7])
	if res != 0 {
		err = errors.New("writeprop err")
	}
	return
}

// func Whois(Wtype string, Wvalue string) (res []Who, err error) {
// 	arg0 := (*reflect.StringHeader)(unsafe.Pointer(&Wtype))
// 	arg1 := (*reflect.StringHeader)(unsafe.Pointer(&Wvalue))
// 	cStruct := C.Whois((*C.char)(unsafe.Pointer(arg0.Data)), (*C.char)(unsafe.Pointer(arg1.Data)))
// 	res = CArrayToGoArray_Who(unsafe.Pointer(cStruct.array), int(cStruct.size))
// 	if len(res) == 0 {
// 		err = errors.New("Whois err")
// 		fmt.Println(err)
// 	}
// 	return
// }
// func Readprop(devID string, objType string, objID string, prop string, index string) (res string, err error) {
// 	arg0 := (*reflect.StringHeader)(unsafe.Pointer(&devID))
// 	arg1 := (*reflect.StringHeader)(unsafe.Pointer(&objType))
// 	arg2 := (*reflect.StringHeader)(unsafe.Pointer(&objID))
// 	arg3 := (*reflect.StringHeader)(unsafe.Pointer(&prop))
// 	arg4 := (*reflect.StringHeader)(unsafe.Pointer(&index))
// 	fmt.Println(arg0, arg1, arg2, arg3, arg4)
// 	cRes := C.Readprop((*C.char)(unsafe.Pointer(arg0.Data)),
// 		(*C.char)(unsafe.Pointer(arg1.Data)),
// 		(*C.char)(unsafe.Pointer(arg2.Data)),
// 		(*C.char)(unsafe.Pointer(arg3.Data)),
// 		(*C.char)(unsafe.Pointer(arg4.Data)))
// 	res = string(CArrayToGoArray_char(unsafe.Pointer(cRes), 32))
// 	if len(res) == 0 {
// 		err = errors.New("Readprop err")
// 		fmt.Println(err)
// 	}
// 	return
// }

func CArrayToGoArray_Who(cArray unsafe.Pointer, size int) (goArray []Who) {
	p := uintptr(cArray)
	for i := 0; i < size; i++ {
		j := *(*Who)(unsafe.Pointer(p))
		goArray = append(goArray, j)
		p += unsafe.Sizeof(j)
	}
	return
}
func CArrayToGoArray_char(cArray unsafe.Pointer, size int) (goArray []byte) {
	p := uintptr(cArray)
	for i := 0; i < size; i++ {
		j := *(*byte)(unsafe.Pointer(p))
		goArray = append(goArray, j)
		p += unsafe.Sizeof(j)
	}
	return
}
