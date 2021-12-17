package gobacnet

type ReadM_para struct {
	ObjType      string
	ObjInstance  uint32
	PropAndIndex string
}

type WriteM_para struct {
	ObjType      string
	ObjInstance  uint32
	PropAndIndex string
	Priority     uint32
	Tag          string
	Value        string
}
