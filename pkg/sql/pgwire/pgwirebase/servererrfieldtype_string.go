// Code generated by "stringer -type=ServerErrFieldType"; DO NOT EDIT.

package pgwirebase

import "strconv"

const (
	_ServerErrFieldType_name_0 = "ServerErrFieldSQLStateServerErrFileldDetail"
	_ServerErrFieldType_name_1 = "ServerErrFieldSrcFile"
	_ServerErrFieldType_name_2 = "ServerErrFileldHint"
	_ServerErrFieldType_name_3 = "ServerErrFieldSrcLineServerErrFieldMsgPrimary"
	_ServerErrFieldType_name_4 = "ServerErrFieldSrcFunctionServerErrFieldSeverity"
)

var (
	_ServerErrFieldType_index_0 = [...]uint8{0, 22, 43}
	_ServerErrFieldType_index_3 = [...]uint8{0, 21, 45}
	_ServerErrFieldType_index_4 = [...]uint8{0, 25, 47}
)

func (i ServerErrFieldType) String() string {
	switch {
	case 67 <= i && i <= 68:
		i -= 67
		return _ServerErrFieldType_name_0[_ServerErrFieldType_index_0[i]:_ServerErrFieldType_index_0[i+1]]
	case i == 70:
		return _ServerErrFieldType_name_1
	case i == 72:
		return _ServerErrFieldType_name_2
	case 76 <= i && i <= 77:
		i -= 76
		return _ServerErrFieldType_name_3[_ServerErrFieldType_index_3[i]:_ServerErrFieldType_index_3[i+1]]
	case 82 <= i && i <= 83:
		i -= 82
		return _ServerErrFieldType_name_4[_ServerErrFieldType_index_4[i]:_ServerErrFieldType_index_4[i+1]]
	default:
		return "ServerErrFieldType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
