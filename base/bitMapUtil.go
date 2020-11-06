package base

type FunctionBitType uint64

func (set FunctionBitType) __maxCapacity() uint32 {
	return 64 // uint64
}

// ToFunctionCodes 解析出存储的 FunctionCodeType 列表
func (set FunctionBitType) ToFunctionCodes() []FunctionCodeType {
	var array []FunctionCodeType
	code := uint64(0x01)
	for idx := uint32(1); idx <= set.__maxCapacity(); idx++ {
		if uint64(set)&code == code {
			array = append(array, FunctionCodeType(idx))
		}
		code <<= 1
	}
	return array
}

// HasFunction 判断是否存在FunctionCode
func (set FunctionBitType) HasFunction(code FunctionCodeType) bool {
	bit := code.ToFunctionBitType()
	return bit != 0 && (set&bit == bit)
}

// AddFunction 将AddFunction添加(压缩)到set中
func (set *FunctionBitType) AddFunction(code FunctionCodeType) bool {
	if code == FunctionCode_NONE {
		return true
	}
	if uint32(code) > set.__maxCapacity() {
		// out of bit-set capacity
		return false
	}
	*set |= FunctionBitType(0x01 << (code - 1))
	return true
}

// AddFunctionList 将AddFunction list添加(压缩)到set中
func (set *FunctionBitType) AddFunctionList(codes []FunctionCodeType) bool {
	for _, code := range codes {
		if !set.AddFunction(code) {
			return false
		}
	}
	return true
}

type FunctionCodeType uint32

const (
	FunctionCode_NONE                  FunctionCodeType = 0
	FunctionCode_SDK_Tokenized_Payment FunctionCodeType = 2
	FunctionCode_Payment_Gateway       FunctionCodeType = 3
	FunctionCode_Oauth                 FunctionCodeType = 5
	FunctionCode_Scan_QR               FunctionCodeType = 7 // B scan C
	FunctionCode_Static_QR             FunctionCodeType = 8 // C scan B (static QR)
	FunctionCode_Auth_Direct_Payment   FunctionCodeType = 9
	FunctionCode_Dynamic_QR            FunctionCodeType = 11 // C scan B (dynamic QR)
	FunctionCode_Offline_Topup         FunctionCodeType = 16
	FunctionCode_Deal                  FunctionCodeType = 20 // deal
	// 添加功能需同步修改 IsValidFunctionCode 范围
)

func (code FunctionCodeType) IsValidFunctionCode(allowEmpty bool) bool {
	if code == FunctionCode_NONE && allowEmpty {
		return true
	}
	return true
}

func (code FunctionCodeType) ToFunctionBitType() FunctionBitType {
	if code == 0 || code > 64 {
		return 0
	}
	return FunctionBitType(0x01 << (code - 1))
}
