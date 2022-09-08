package rizingo

// Operation is the representation of an assemble operation.
type Operation struct {
	Offset    int      `json:"offset"`
	Esil      string   `json:"esil"`
	Refptr    bool     `json:"refptr"`
	FcnAddr   int      `json:"fcn_addr"`
	FcnLast   int      `json:"fcn_last"`
	Size      int      `json:"size"`
	Opcode    string   `json:"opcode"`
	Disasm    string   `json:"disasm"`
	Bytes     string   `json:"bytes"`
	Family    string   `json:"family"`
	Type      string   `json:"type"`
	Reloc     bool     `json:"reloc"`
	TypeNum   int      `json:"type_num"`
	Type2Num  int      `json:"type2_num"`
	Flags     []string `json:"flags"`
	Comment   string   `json:"comment"`
	XrefsTo   []XRef   `json:"xrefs_to"`
	Val       int      `json:"val,omitempty"`
	Ptr       int      `json:"ptr,omitempty"`
	Jump      int      `json:"jump,omitempty"`
	XrefsFrom []XRef   `json:"xrefs_from,omitempty"`
}

// XRef is a cross references of an address
type XRef struct {
	Addr int    `json:"addr"`
	Type string `json:"type"`
}

// Function is the representation of a function identified by Rizin
type Function struct {
	Name string      `json:"name"`
	Size int         `json:"size"`
	Addr int         `json:"addr"`
	Ops  []Operation `json:"ops"`
}
