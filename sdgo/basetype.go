package sdgo

import (
	"fmt"
	"encoding/hex"
	"strings"
	"strconv"
)


// for json encode
type HexByte []byte
func (m HexByte) MarshalJSON() ([]byte, error) {
	if m == nil || len(m) == 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%02X"`, m)), nil
}
func (m *HexByte) UnmarshalJSON(in []byte) error {
	buf, err := hex.DecodeString(strings.Trim(string(in), `"`))
	if err != nil {
		return err
	}
	*m = buf
	return nil
}

type HexUint16 uint16
func (m HexUint16) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%X"`, m)), nil
}
func (m *HexUint16) UnmarshalJSON(in []byte) error {
	tmp, err :=  strconv.ParseUint(strings.Trim(string(in), `"`), 16, 16)
	if err != nil {
		return err
	}
	*m = HexUint16(tmp)
	return nil
}

type HexUint32 uint32
func (m HexUint32) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%X"`, m)), nil
}
func (m *HexUint32) UnmarshalJSON(in []byte) error {
	tmp, err :=  strconv.ParseUint(strings.Trim(string(in), `"`), 16, 32)
	if err != nil {
		return err
	}
	*m = HexUint32(tmp)
	return nil
}

type HexUint64 uint64
func (m HexUint64) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%X"`, m)), nil
}
func (m *HexUint64) UnmarshalJSON(in []byte) error {
	tmp, err :=  strconv.ParseUint(strings.Trim(string(in), `"`), 16, 64)
	if err != nil {
		return err
	}
	*m = HexUint64(tmp)
	return nil
}

type HexColor16 uint16
func (m *HexColor16) UnmarshalJSON(in []byte) error {
	return m.ParseColor(strings.Trim(string(in), `"`), uint16(*m))
}
func (m *HexColor16) ParseColor(str string, def uint16) error {
	tmp, err := strconv.ParseUint(str, 16, 32)
	if err != nil {
		*m = HexColor16(def)
		return err
	}
	out := uint16((tmp & 0xFF) >> 3) // B
	out |= uint16(((tmp >> 8) & 0xFF) >> 3) << 5 // G
	out |= uint16(((tmp >> 16) & 0xFF) >> 3) << 10 // R
	*m = HexColor16(out)
	return nil
}
func (m HexColor16) DumpColor() (out uint32) {
	out = uint32((m & 0x1F) << 3) // B
	out |= uint32(((m >> 5) & 0x1F) << 3) << 8 // G
	out |= uint32(((m >> 10) & 0x1F) << 3) << 16 // R
	return out
}

// for print order
type HexBotID uint16
func (m *HexBotID) UnmarshalJSON(in []byte) error {
	tmp, err := hex.DecodeString(strings.Trim(string(in), `"`))
	if len(tmp) == 2 && err == nil {
		*m = HexBotID((uint16(tmp[1]) << 8) | uint16(tmp[0]))
	}
	return err
}
func (m HexBotID) String() string {
	id := uint16((m >> 8) | ((m & 0xFF) << 8))
	return fmt.Sprintf("%04X", id)
}


