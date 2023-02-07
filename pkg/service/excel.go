package service

import (
	"errors"
	"github.com/xuri/excelize/v2"
	"os"
)

func (service *ServiceT) writeExcel(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		os.Remove(path)
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()
	index, err := f.NewSheet("Altitudes")
	if err != nil {
		return err
	}
	for ir, row := range service.Area {
		for ic, col := range row {
			cell, err := excelize.CoordinatesToCellName(ic+1, ir+1)
			if err != nil {
				return err
			}
			err = f.SetCellValue("Altitudes", cell, col.Altitude)
			if err != nil {
				return err
			}
		}
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(path); err != nil {
		return err
	}
	return nil
}
