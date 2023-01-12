package main

func convertToTitle(columnNumber int) string {
	var res []byte
	for columnNumber > 0 {
		res = append([]byte{byte((columnNumber-1)%26) + 'A'}, res...)
		columnNumber = (columnNumber - 1) / 26
	}
	return string(res)
}
