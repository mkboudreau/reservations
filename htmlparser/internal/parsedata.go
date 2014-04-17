package internal

import "strings"

type ParseData struct {
	Index int
	Data  string
	//TODO: maybe put in io.Reader
}

func (data *ParseData) HasNext() bool {
	return (data.Index) < len(data.Data)
}

func (data *ParseData) Next() string {
	str := string(data.Data[data.Index])
	data.Index++
	return str
}
func (data *ParseData) PeekCurrent() string {
	return data.PeekBy(0)
}
func (data *ParseData) PeekNext() string {
	return data.PeekBy(1)
}
func (data *ParseData) PeekBy(num int) string {
	if data.Index+num >= len(data.Data) {
		return ""
	}
	return string(data.Data[data.Index+num])
}
func (data *ParseData) PeekTo(length int) string {
	start := data.Index
	requestedEndIndex := data.Index + length
	if requestedEndIndex >= len(data.Data) {
		requestedEndIndex = len(data.Data)
	}
	return string(data.Data[start:requestedEndIndex])
}
func (data *ParseData) SeekTo(length int) string {
	str := string(data.Data[data.Index : data.Index+length])
	data.Index += length
	return str
}
func (data *ParseData) IsNext(someString string) bool {
	return data.PeekTo(len(someString)) == someString
}
func (data *ParseData) IsNextCaseInsensitive(someString string) bool {
	str := strings.ToUpper(someString)
	return strings.ToUpper(data.PeekTo(len(str))) == str
}
