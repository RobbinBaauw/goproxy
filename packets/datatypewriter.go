package packets

func WriteVarInt(value int) []byte {
	var result []byte

	for ok := true; ok; ok = value != 0 {
		temp := byte(value & 0b01111111)
		value = int(uint(value) >> 7)

		if value != 0 {
			temp |= 0b10000000;
		}

		result = append(result, temp);
	}

	return result
}

func WriteString(value string) []byte {
	bytes := []byte(value)
	return WriteStringBytes(bytes)
}

func WriteStringBytes(bytes []byte) []byte {
	lengthBytes := WriteVarInt(len(bytes))
	return append(lengthBytes, bytes...)
}
