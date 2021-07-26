package handleFile

import (
	"fmt"
	"os"
)

func WriteToFile(filename string, newcontent string) {
	// nếu filename tồn tại hàm sẽ ghi thêm vào, nếu không hàm sẽ tạo file mới
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err == nil {
		//File tồn tại và mở thành công
		_, err = fmt.Fprintln(file, newcontent)
		if err == nil {
			fmt.Println("\nFile was found and content was written to file !")
		} else {
			fmt.Println("\nFile was found but can't write to newfile !")
		}
		file.Close()
	} else {
		//File không tồn tại, hàm sẽ tạo file mới
		file.Close()

		file2, err2 := os.Create(filename)
		if err2 != nil {
			fmt.Println("\nFile not found and Can't create newfile !")
			file2.Close()
			return
		}

		_, err3 := file2.WriteString(newcontent)
		if err3 != nil {
			fmt.Println("\nCreated newfile but can't write to newfile !")
			file2.Close()
			return
		} else {
			fmt.Println("\nCreated newfile and content was written to file !")
			file2.Close()
		}
	}
}
