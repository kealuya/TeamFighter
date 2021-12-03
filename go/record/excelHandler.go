package main

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"path/filepath"
	"strconv"
)
/*
操作excel的案例
 */
func excelHandler() {
	myPath := "F:\\YiwanDownloads\\"
	err := filepath.Walk(myPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".xlsx" {
			fmt.Println(path)
			xlsx, _ := excelize.OpenFile(path)
			cell_slice := xlsx.GetRows(xlsx.GetSheetName(1))

			for row := 0; row < len(cell_slice); row++ {
				for col := 0; col < len(cell_slice[row]); col++ {
					value := cell_slice[row][col]
					if value != "" {
						var buf bytes.Buffer
						buf.WriteString("   ")
						buf.WriteString("range[")
						buf.WriteString(strconv.Itoa(row + 1))
						buf.WriteString(",")
						buf.WriteString(strconv.Itoa(col + 1))
						buf.WriteString("] = ")
						buf.WriteString(value)
						fmt.Println(buf.String())
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
}
