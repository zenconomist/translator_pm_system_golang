package readfiles

import "os"

func GetFileContents(filename string) ([]byte, error) {
	bytes, errReadSql := os.ReadFile("./packages/readfiles/" + filename)
	if errReadSql != nil {
		// retry if called from service testing
		bytes, errReadSql = os.ReadFile("./../readfiles/" + filename)
		if errReadSql != nil {
			return bytes, errReadSql
		}
	}
	return bytes, nil
}
