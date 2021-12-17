#ifndef BAC_H
#define BAC_H

#include "struct.h"

// size of read device property result
#define READ_RES_SIZE 32
#define READ_RES_SIZE_M 128

//
Wl_array Whois(char* type, char* val);

//
char* Readprop(char* devID, char* objType, char* objID, char* prop, char* index);
//
char* ReadPropM(char* deviceID, char* size, char* args);

//
int Writeprop(char* devID, char* objType, char* objID, char* prop, char* priority, char* index, char* tag, char* value);
//
int WritePropM(char* argc, char *args);

#endif