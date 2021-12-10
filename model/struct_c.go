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
*/
import "C"

type Who C.Who
type Wl_array C.Wl_array
