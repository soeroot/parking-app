package service

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	constant "soeroot/parking-app/const"
	"soeroot/parking-app/models"
)

type ParkingService struct {
	parking  *models.Parking
	PathFile string
}

func NewParkingService(pathFile string) *ParkingService {
	return &ParkingService{
		PathFile: pathFile,
	}
}

func (service *ParkingService) BatchCommandFromFile() {
	if service.PathFile != "" {
		var inputReader io.Reader

		file, err := os.Open(service.PathFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file %s: %v", service.PathFile, err)
			return
		}
		defer file.Close()
		inputReader = file
		fmt.Printf("reading from file: %s\n", service.PathFile)

		scanner := bufio.NewScanner(inputReader)
		for scanner.Scan() {
			strSplit := strings.Split(scanner.Text(), " ")
			cmd := strings.ToLower(strSplit[0])

			var valStr string
			switch cmd {
			case constant.CreateParkingLotArg:
				valStr = strSplit[1]
				val, err := strconv.Atoi(valStr)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to create parking lot, err: %v", err)
					return
				}
				service.parking = models.NewParking(val)
			case constant.ParkArg:
				valStr = strSplit[1]
				slot, err := service.parking.Park(valStr)
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n\n", err.Error())
				} else {
					fmt.Fprintf(os.Stdout, "Allocated slot number: %v\n", slot)
				}

			case constant.LeaveArg:
				valStr = strSplit[1]
				if len(strSplit) < 3 {
					fmt.Fprintf(os.Stderr, "required 3 parameter")
					return
				}

				hourStr := strSplit[2]
				hour, err := strconv.Atoi(hourStr)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to free parking slot, err: %v", err)
					return
				}

				charge := 10
				if hour-2 != 0 {
					charge += charge * (hour - 2)
				}

				slot, err := service.parking.Leave(valStr)
				if err != nil {
					fmt.Fprintf(os.Stdout, "%v\n", err.Error())
				}

				fmt.Printf("Registration Number %v with slot number %v is free with charge %v\n", valStr, slot, charge)
			case constant.StatusArg:
				fmt.Printf("Slot No.\t Registration No.\n")

				keys := make([]int, 0, len(service.parking.Occupied))
				for k := range service.parking.Occupied {
					keys = append(keys, k)
				}

				sort.Ints(keys)

				for _, key := range keys {
					fmt.Printf("%v\t\t %v\n", key, service.parking.Occupied[key].PlateNumber)
				}
			default:
				fmt.Fprintf(os.Stdout, "argument %s in the file is invalid\n", cmd)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to run parking service, err: %v", err)
			return
		}
	}
}
