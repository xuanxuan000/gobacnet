package gobacnet

/*
#include <stdint.h>
typedef struct WhoIs
{
	uint32_t Device;
	uint8_t Mac[6];
}Who;

typedef struct Whois_list_array{
	int size;
	int capacity;
	Who* array;
}Wl_array;

typedef struct ReadPropM_Para
{
	char* objType;
	char* objInstance;
	char* propAndIndex;
}ReadM_para;
*/
import "C"

type Who C.Who
type Wl_array C.Wl_array
type ReadM_para C.ReadM_para
