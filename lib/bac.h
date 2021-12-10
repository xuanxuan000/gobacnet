#ifndef BAC_H
#define BAC_H

#include "struct.h"

// size of read device property result
#define READ_RES_SIZE 32

//
Wl_array Whois(char* type, char* val);

//
char* Readprop(char* devID, char* objType, char* objID, char* prop, char* index);

//
int Writeprop(char* devID, char* objType, char* objID, char* prop, char* priority, char* index, char* tag, char* value);

#endif