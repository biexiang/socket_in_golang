package util

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

/**
 * 加密 包头添加长度
 * @param  {[type]} input string)       (bytes, err [description]
 * @return {[type]}       [description]
 */
func Encode(input string) ([]byte, error) {
	var len = int32(len(input))
	var pkg *bytes.Buffer = new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, len)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(input))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(input *bufio.Reader) (string, error) {
	//读取长度
	slice, _ := input.Peek(4)
	sliceData := bytes.NewBuffer(slice)
	var len int32
	err := binary.Read(sliceData, binary.LittleEndian, &len)
	if err != nil {
		return "", err
	}

	if int32(input.Buffered()) < len+4 {
		//should be error
		return "", nil
	}

	pack := make([]byte, len+4)
	_, err = input.Read(pack)
	if err != nil {
		return "", err
	}

	return string(pack[4:]), nil
}
