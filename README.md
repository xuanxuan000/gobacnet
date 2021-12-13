# `gobacnet`

*BACnet for go*

---

### Whois:

**Send BACnet WhoIs service request to a device or multiple devices, and wait for responses. Displays any devices found and their network information.**

#### parameters:

`type`: support: devid, mac, dnet, range

`val`: value

#### *examples:*

Send a WhoIs request to DNET 123:

```
Whois("dnet", "123")
```

Send a WhoIs request to MAC 10.1.2.3:47808:

```
Whois("mac", "10.1.2.3:4780812")
```

Send a WhoIs request to Device 123:

```
Whois("range", "123-123")
```

Send a WhoIs request to Devices from 1000 to 9000:

```
Whois("range", "1000-9000")
```

Send a WhoIs request to all devices:

```
Whois("", "")
```

## Readprop:

**Reads the value of a property of a device**

#### parameters:

`devID`: BACnet Device Object Instance number that you are trying to communicate to.  This number will be used to try and bind with the device using 			Who-Is and I-Am services.  For example, if you were reading Device Object 123, the device-instance would be 123.

`objType`: The object type is object that you are reading. It can be defined either as the object-type name string as defined in the BACnet specification, or as the integer value of the enumeration BACNET_OBJECT_TYPE in bacenum.h. For example if you were reading Analog Output 2, the object-type would be analog-output or 1.

`objID`: This is the object instance number of the object that you are reading.  For example, if you were reading Analog Output 2, the object-instance would be 2.

`prop`: The property of the object that you are reading. It can be defined either as the property name string as defined in the BACnet specification, or as an integer value of the enumeration BACNET_PROPERTY_ID in bacenum.h. For example, if you were reading the Present Value property, use present-value or 85 as the property.

`index`: This integer parameter is the index number of an array. If the property is an array, individual elements can be read.  If this parameter is missing and the property is an array, the entire array will be read.

#### *examples:*

Read the Present-Value of Analog Output 101 in Device 123, you could use either of the following

```
Readprop("123", "analog_output", "101", "present-value", "1")
```

```
Readprop("123", "1", "101", "85", "1")
```

## Writeprop:

**Writes the value of a property of a device**

#### parameters:

`devID`: BACnet Device Object Instance number that you are trying to communicate to.  This number will be used to try and bind with the device using Who-Is and I-Am services.  For example, if you were writing to Device Object 123, the device-instance would be 123.

`objType`: The object type is object that you are reading. It can be defined either as the object-type name string as defined in the BACnet specification, or as the integer value of the enumeration BACNET_OBJECT_TYPE in bacenum.h. For example if you were reading Analog Output 2, the object-type would be analog-output or 1.

`objID`: This is the object instance number of the object that you are writing to.  For example, if you were writing to Analog Output 2, the object-instance would be 2.

`prop`: The property of the object that you are reading. It can be defined either as the property name string as defined in the BACnet specification, or as an integer value of the enumeration BACNET_PROPERTY_ID in bacenum.h. For example, if you were reading the Present Value property, use present-value or 85 as the property.

`priority`: This parameter is used for setting the priority of the write. If Priority 0 is given, no priority is sent.  The BACnet standard states that the value is written at the lowest priority (16) if the object property supports priorities when no priority is sent.

`index`: This integer parameter is the index number of an array. If the property is an array, individual elements can be written to if supported.  If this parameter is -1, the index is ignored.

`tag`: Tag is the integer value of the enumeration BACNET_APPLICATION_TAG in bacenum.h.  It is the data type of the value that you are writing.  For example, if you were writing a REAL value, you would use a tag of 4. Context tags are created using two tags in a row.  The context tag is preceded by a C.  Ctag tag. C2 4 creates a context 2 tagged REAL.

`value`: The value is an ASCII representation of some type of data that you are writing.  It is encoded using the tag information provided.  For example, if you were writing a REAL value of 100.0, you would use 100.0 as the value.

#### *examples:*

Send a value of 100 to the Present-Value in Analog Output 0 of Device 123 at priority 16, use the one of following commands:

```
Writeprop("123", OBJECT_ANALOG_OUTPUT, "0", PROP_PRESENT_VALUE, "16", "-1", BACNET_APPLICATION_TAG_SIGNED_INT, "100")
```

```
Writeprop("111", OBJECT_ANALOG_OUTPUT, "101", "85", "16", "-1", "3", "100")
```
