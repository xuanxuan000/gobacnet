# `gobacnet`

*BACnet for go*

### Whois:

**Send BACnet WhoIs service request to a device or multiple devices, and wait for responses. Displays any devices found and their network information.**

#### parameters:

type:

support: mac, dnet, range

#### *examples:*

Send a WhoIs request to DNET 123:

Whois("dnet", "123")

Send a WhoIs request to MAC 10.1.2.3:47808:

Whois("mac", "10.1.2.3:4780812")

Send a WhoIs request to Device 123:

Whois("range", "123-123")

Send a WhoIs request to Devices from 1000 to 9000:

Whois("range", "1000-9000")

Send a WhoIs request to all devices:

Whois("", "")



## Readprop:

**Send BACnet WhoIs service request to a device or multiple devices, and wait for responses. Displays any devices found and their network information.**

#### *examples:*

Send a WhoIs request to DNET 123:

Whois("dnet", "123")

Send a WhoIs request to MAC 10.1.2.3:47808:

Whois("mac", "10.1.2.3:4780812")

Send a WhoIs request to Device 123:

Whois("range", "123-123")

Send a WhoIs request to Devices from 1000 to 9000:

Whois("range", "1000-9000")

Send a WhoIs request to all devices:

Whois("", "")
