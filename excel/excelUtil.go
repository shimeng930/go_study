package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadExcel(filePath string) ([][]string, error) {
	xlsx, err := excelize.OpenFile(filePath)
	if err !=nil{
		return nil, err
	}

	return xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex())), nil
}


func ExcelDo()  {
	//writeExcel()
	logicNew()
}

func getOutletQrTokenMap(filePath string) map[int64]string {
	outletTokenMap := make(map[int64]string)
	excelData, err := ReadExcel(filePath)
	if err != nil {
		log.Printf("ERROR|MakeFinanceInfo|read excel error|err=%v", err)
	}

	if excelData == nil || len(excelData) == 0 {
		log.Printf("ERROR|MakeFinanceInfo|empty excel")
	}
	for i, row := range excelData {
		if i == 0 {
			continue
		}
		outletId, _ := strconv.ParseInt(row[0], 10, 64)
		outletTokenMap[outletId] = row[1]
	}
	return outletTokenMap
}

func getOutletQrIdMap(filePath string) map[int64]int64 {
	outletQrMap := make(map[int64]int64)
	excelData, err := ReadExcel(filePath)
	if err != nil {
		log.Printf("ERROR|MakeFinanceInfo|read excel error|err=%v", err)
	}

	if excelData == nil || len(excelData) == 0 {
		log.Printf("ERROR|MakeFinanceInfo|empty excel")
	}
	for _, row := range excelData {
		outletId, _ := strconv.ParseInt(row[0], 10, 64)
		qrId, _ := strconv.ParseInt(row[1], 10, 64)
		outletQrMap[outletId] = qrId
	}
	return outletQrMap
}

func getQrNameMap(filePath string) map[int64]string {
	outletQrName := make(map[int64]string)
	excelData, err := ReadExcel(filePath)
	if err != nil {
		log.Printf("ERROR|MakeFinanceInfo|read excel error|err=%v", err)
	}

	if excelData == nil || len(excelData) == 0 {
		log.Printf("ERROR|MakeFinanceInfo|empty excel")
	}
	for _, row := range excelData {
		outletId, _ := strconv.ParseInt(row[0], 10, 64)
		outletQrName[outletId] = row[1]
	}
	return outletQrName
}

func getQrTokenMap(filePath string) map[int64]string {
	qrTokenMap := make(map[int64]string)
	excelData, err := ReadExcel(filePath)
	if err != nil {
		log.Printf("ERROR|MakeFinanceInfo|read excel error|err=%v", err)
	}

	if excelData == nil || len(excelData) == 0 {
		log.Printf("ERROR|MakeFinanceInfo|empty excel")
	}
	for _, row := range excelData {
		qrId, _ := strconv.ParseInt(row[0], 10, 64)
		qrTokenMap[qrId] = row[1]
	}
	return qrTokenMap
}

func getOutletNameMap(filePath string) map[int64]string {
	outletNameMap := make(map[int64]string)
	excelData, err := ReadExcel(filePath)
	if err != nil {
		log.Printf("ERROR|MakeFinanceInfo|read excel error|err=%v", err)
	}

	if excelData == nil || len(excelData) == 0 {
		log.Printf("ERROR|MakeFinanceInfo|empty excel")
	}
	for i, row := range excelData {
		if i == 0 {
			continue
		}
		outletId, _ := strconv.ParseInt(row[0], 10, 64)
		outletNameMap[outletId] = row[1]
	}
	return outletNameMap
}

func logic() {
	outletTokenMap := getOutletQrTokenMap("./static/old_qr_token.xlsx")
	qrTokenMap := getQrTokenMap("./static/old_qr_and_token.xlsx")
	outletQRIdMap := getOutletQrIdMap("./static/old_out_qr_id.xlsx")

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "outlet")
	xlsx.SetCellValue("Sheet1", "B1", "qr_token")
	rowIndex := 2
	for outletId, _ := range outletTokenMap {
		rowA := fmt.Sprintf("A%d", rowIndex)
		rowB := fmt.Sprintf("B%d", rowIndex)
		qrId, ok := outletQRIdMap[outletId]
		if ok {
			outletTokenMap[outletId] = qrTokenMap[qrId]
		}

		xlsx.SetCellValue("Sheet1", rowA, outletId)
		xlsx.SetCellValue("Sheet1", rowB, outletTokenMap[outletId])
		rowIndex++

	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func logicNew() {
	outletTokenMap := getOutletQrTokenMap("./static/new_qr_token.xlsx")
	qrTokenMap := getQrTokenMap("./static/new_qr_id_token.xlsx")
	outletQRIdMap := getOutletQrIdMap("./static/new_outlet_qr_id.xlsx")
	qrNameMap := getQrNameMap("./static/new_qr_name.xlsx")

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "outlet")
	xlsx.SetCellValue("Sheet1", "B1", "qr_token")
	xlsx.SetCellValue("Sheet1", "C1", "qr_name")
	rowIndex := 2
	ai := 0
	arr := []int64{}
	for outletId, _ := range outletTokenMap {
		rowA := fmt.Sprintf("A%d", rowIndex)
		rowB := fmt.Sprintf("B%d", rowIndex)
		rowC := fmt.Sprintf("C%d", rowIndex)
		qrId, ok := outletQRIdMap[outletId]
		if ok {
			outletTokenMap[outletId] = qrTokenMap[qrId]
		}

		if ai % 20 != 0 {
			arr = append(arr, outletId)
		} else {
			arr = []int64{}
		}
		ai++

		xlsx.SetCellValue("Sheet1", rowA, outletId)
		xlsx.SetCellValue("Sheet1", rowB, outletTokenMap[outletId])
		xlsx.SetCellValue("Sheet1", rowC, qrNameMap[outletId])
		rowIndex++

	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func ReadDir() map[int64]bool {
	outletQRMap := make(map[int64]bool)
	rd1, _ := ioutil.ReadDir("./static/old_qr_1")
	for _, fi := range rd1 {
		pngName := fi.Name()
		names := strings.Split(pngName, "_")
		names = strings.Split(names[1], ".")
		oid, _ := strconv.ParseInt(names[0], 10, 64)
		outletQRMap[oid] = true
	}
	return outletQRMap
}

func CheckUndoQr() {
	outletTokenMap := getOutletQrTokenMap("./static/old_need.xlsx")
	actualMap := ReadDir()

	needFixMap := make(map[int64]string)
	for outletId, token := range outletTokenMap {
		if !actualMap[outletId] {
			needFixMap[outletId] = token
		}
	}

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "outlet")
	xlsx.SetCellValue("Sheet1", "B1", "qr_token")
	rowIndex := 2
	for outletId, _ := range needFixMap {
		rowA := fmt.Sprintf("A%d", rowIndex)
		rowB := fmt.Sprintf("B%d", rowIndex)
		xlsx.SetCellValue("Sheet1", rowA, outletId)
		xlsx.SetCellValue("Sheet1", rowB, needFixMap[outletId])
		rowIndex++

	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("test_need.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func CheckOutletId() {
	outletNameMap := getOutletNameMap("./static/db_outlet_name.xlsx")
	outletQRData := getOutletQrTokenMap("./static/old_outlet_qr_token_final.xlsx")

	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetCellValue("Sheet1", "A1", "outlet")
	xlsx.SetCellValue("Sheet1", "B1", "qr_token")
	xlsx.SetCellValue("Sheet1", "C1", "qr_token")
	rowIndex := 2
	for outletId, name := range outletNameMap {
		rowA := fmt.Sprintf("A%d", rowIndex)
		rowB := fmt.Sprintf("B%d", rowIndex)
		rowC := fmt.Sprintf("C%d", rowIndex)

		xlsx.SetCellValue("Sheet1", rowA, outletId)
		xlsx.SetCellValue("Sheet1", rowB, outletQRData[outletId])
		xlsx.SetCellValue("Sheet1", rowC, name)
		rowIndex++

	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("actual_outlet.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}