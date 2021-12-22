package gobacnet

type ReadM_para struct {
	ObjType      string // analog-input
	ObjInstance  uint32
	PropAndIndex string // 属性以及下标 85[1]
}

// Priority于设置写的优先级。 如果指定优先级0，则不发送优先级。
// BACnet标准规定，如果object属性支持优先级，而不发送优先级，则以最低优先级(16)写入值。
type WriteM_para struct {
	ObjType      string // analog-input
	ObjInstance  uint32
	PropAndIndex string // 属性以及下标 85[1]
	Priority     uint32 // 优先级 1-16
	Tag          string // 参数类型
	Value        string // 值
}
