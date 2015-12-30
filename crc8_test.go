package crc8

import (
	"testing"
)

var testData = []byte("123456789")

func testSelectedCRC8(params Params, t *testing.T) {
	table := MakeTable(params)
	if table == nil {
		t.Errorf("Failed to make %q table", params.Name)
	}

	crc := Checksum(testData, table)

	if crc != table.params.Check {
		t.Errorf("Invalid %q sample calculation, expected: %X, actual: %X", table.params.Name, table.params.Check, crc)
	}
}

func TestCRC8(t *testing.T) {
	testSelectedCRC8(CRC8, t)
}

func TestCRC8_CDMA2000(t *testing.T) {
	testSelectedCRC8(CRC8_CDMA2000, t)
}

func TestCRC8_DARC(t *testing.T) {
	testSelectedCRC8(CRC8_DARC, t)
}

func TestCRC8_DVB_S2(t *testing.T) {
	testSelectedCRC8(CRC8_DVB_S2, t)
}

func TestCRC8_EBU(t *testing.T) {
	testSelectedCRC8(CRC8_EBU, t)
}

func TestCRC8_I_CODE(t *testing.T) {
	testSelectedCRC8(CRC8_I_CODE, t)
}

func TestCRC8_ITU(t *testing.T) {
	testSelectedCRC8(CRC8_ITU, t)
}

func TestCRC8_MAXIM(t *testing.T) {
	testSelectedCRC8(CRC8_MAXIM, t)
}

func TestCRC8_ROHC(t *testing.T) {
	testSelectedCRC8(CRC8_ROHC, t)
}

func TestCRC8_WCDMA(t *testing.T) {
	testSelectedCRC8(CRC8_WCDMA, t)
}
