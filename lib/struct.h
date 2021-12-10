#ifndef STRUCT_H
#define STRUCT_H

#include <stdlib.h>
#include "string.h"
#include "stdio.h"
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

#endif