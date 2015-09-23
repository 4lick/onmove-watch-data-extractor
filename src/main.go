package main

import (
   "os"
   "fmt"
   "log"
   "encoding/binary" 
   "strconv"
   "time"
)

func main() { 
				
		// input file 
		if len(os.Args) == 2 {
		  inputFilePath := os.Args[1] 
		   
		  // GHL 
		  fmt.Println("\nGHL file...")
		
	      fileGHL, err := os.Open(inputFilePath + ".GHL")	
	      if err != nil {
	          log.Fatal("Error while opening file")
	      }
          
	      defer fileGHL.Close() 
          
   		  bytesGHL := splitFileToSlice(fileGHL, 48)  		
		  for i := range bytesGHL {
		          parseGHL(bytesGHL[i]) 
		  		fmt.Println("---")
		  } 
          
		  // GHT 
		  fmt.Println("\nGHT file...")
          
	      fileGHT, err := os.Open(inputFilePath + ".GHT")
	      if err != nil {
	          log.Fatal("Error while opening file", err)
	      }
          
	      defer fileGHT.Close() 
          
		  infoGHT, _ := fileGHT.Stat()                       
		  var sizeGHT int64 = infoGHT.Size() 
          
		  bytesGHT := make([]byte, sizeGHT)             
          _, errGHT := fileGHT.Read(bytesGHT) 
    	  if errGHT != nil {
	          log.Fatal("Error while opening file", errGHT)
	      }
          
		  parseGHT(bytesGHT)	
          
		  // GHP 
		  fmt.Println("\nGHP file...")
          
		  fileGHP, err := os.Open(inputFilePath + ".GHP")
		  if err != nil {
		      log.Fatal("Error while opening file", err)
		  }
          
		  defer fileGHP.Close() 
          
		  bytesGHP := splitFileToSlice(fileGHP, 20)  		
		  for i := range bytesGHP {
		          parseGHP(bytesGHP[i]) 
		  		fmt.Println("---")
		  }				
		} else {
   	        log.Fatal("you must specify the path to an input GHL/GHP/GHT file") 
			os.Exit(1)
		}  	     		
} 

func parseFilename(file_name string) time.Time {
	year, _ := strconv.ParseInt(file_name[0:1], 36, 32)
    month, _ := strconv.ParseInt(file_name[1:2], 36, 32)
    day, _ := strconv.ParseInt(file_name[2:3], 36, 32)
    hour, _ := strconv.ParseInt(file_name[3:4], 36, 32)
	minute, _ := strconv.Atoi(file_name[4:6])
	second, _ := strconv.Atoi(file_name[6:])
	
    time_session := time.Date(int(year), time.Month(month), int(day), int(hour), minute, second, 0, time.Local)
	return time_session
} 

func splitFileToSlice(file *os.File, off int) [][]byte {
	// calculate the bytes size                
	info, _ := file.Stat()                   
	var size int64 = info.Size()               
                                               
	s := [][]byte{}                            
	var i int64 = 0                            
	                                           
	for i < size {                             
	    bytes := make([]byte, off)             
        _, err := file.ReadAt(bytes, i)        
        if err != nil {                        
            log.Fatal(err)                     
        }                                      
                                               
	    s = append(s, bytes)                   
	    i += int64(off)                               
	}                                          

    return s 
} 

func parseGHL(data []byte) {
	 // totalTime
		totalTime := binary.LittleEndian.Uint32(data[4:8])
		fmt.Printf("=> totalTime :%f\n", float64(totalTime) / 10)
		
		// totalDistance
		totalDistance := binary.LittleEndian.Uint32(data[8:12])
		fmt.Printf("=> totalTime :%d\n", int(totalDistance))
		
		// maxSpeed
		maxSpeed := binary.LittleEndian.Uint16(data[12:14])
		fmt.Printf("=> maxSpeed :%f\n", float64(maxSpeed) / 100.0)
		
		// averageSpeed
		averageSpeed := binary.LittleEndian.Uint16(data[14:16])
		fmt.Printf("=> averageSpeed :%f\n", float64(averageSpeed) / 100.0)
				
		// maxHeartRate
		maxHeartRate := data[20] 
		fmt.Printf("=> maxHeartRate :%d\n", int(maxHeartRate))
		
		// averageHeartRate
		averageHeartRate := data[21] 
		fmt.Printf("=> averageHeartRate :%d\n", int(averageHeartRate))
				
		// averageCalory
		averageCalory := binary.LittleEndian.Uint16(data[22:24])
		fmt.Printf("=> averageCalory :%d\n", int(averageCalory))
				
		// weightLoss
		weightLoss := binary.LittleEndian.Uint16(data[28:30])
		fmt.Printf("=> weightLoss :%d\n", int(weightLoss))
				
		// averageAscent
		averageAscent := binary.LittleEndian.Uint16(data[44:46])
		fmt.Printf("=> averageAscent :%d\n", int(averageAscent))
				
		// averageDescent
		averageDescent := binary.LittleEndian.Uint16(data[46:48])
		fmt.Printf("=> averageDescent :%d\n", int(averageDescent))
		
		// startPoint					 
		startPoint := binary.LittleEndian.Uint16(data[40:42])
		fmt.Println("=> startPoint :", startPoint)
		
		// endPoint
		endPoint := binary.LittleEndian.Uint16(data[42:44])
		fmt.Println("=> endPoint :", endPoint)	
} 

func parseGHT(data []byte) { 
 		//totalPoint 
		totalPoint := binary.LittleEndian.Uint16(data[6:8])  
 		fmt.Printf("=> totalPoint :%d\n", int(totalPoint))

		//totalTime
		totalTime := binary.LittleEndian.Uint32(data[8:12]) 
		fmt.Printf("=> totalTime :%f\n", float64(totalTime) / 10)

		//totalDistance 
		totalDistance := binary.LittleEndian.Uint32(data[12:16]) 
		fmt.Printf("=> totalDistance :%d\n", int(totalDistance))

		//lapCount
		lapCount := binary.LittleEndian.Uint16(data[16:18]) 
		fmt.Printf("=> lapCount :%d\n", int(lapCount))

		//challengeName 
		challengeName := string(data[24:41])
		fmt.Printf("=> challengeName : %s\n", challengeName)

		//maxSpeed 
		maxSpeed := binary.LittleEndian.Uint16(data[52:54]) 
		fmt.Printf("=> maxSpeed :%f\n", float64(maxSpeed) / 100)

		//averageSpeed 
		averageSpeed := binary.LittleEndian.Uint16(data[54:56]) 
		fmt.Printf("=> averageSpeed :%f\n", float64(averageSpeed) / 100)

		//maxHeartRate
		maxHeartRate := int(data[60])
		fmt.Printf("=> maxHeartRate :%d\n", maxHeartRate)

		//averageHeartRate 
		averageHeartRate := int(data[61]) 
		fmt.Printf("=> averageHeartRate :%d\n", averageHeartRate)

		//totalCalory   
		totalCalory := binary.LittleEndian.Uint16(data[66:68])
		fmt.Printf("=> totalCalory :%d\n", int(totalCalory))
		
		//weightLoss   
		weightLoss := binary.LittleEndian.Uint16(data[68:70]) 
		fmt.Printf("=> weightLoss :%d\n", int(weightLoss))
		
		//ascent 
		ascent := binary.LittleEndian.Uint16(data[70:72]) 
		fmt.Printf("=> ascent :%d\n", int(ascent))
		
		//descent 
		descent := binary.LittleEndian.Uint16(data[72:74]) 
		fmt.Printf("=> descent :%d\n", int(descent))  
		
		//trainingName
		trainingName := string(data[76:91])		 
		fmt.Printf("=> trainingName : %s\n", trainingName)
}  

func parseGHP(data []byte) { 
		//latitude 
		latitude := binary.LittleEndian.Uint32(data[0:4]) 
		fmt.Printf("=> latitude :%f\n", float64(latitude) / 1000000)
		
		//longitude 
		longitude := binary.LittleEndian.Uint32(data[4:8]) 
		fmt.Printf("=> longitude :%f\n", float64(longitude) / 1000000)
		
		//altitude 
		altitude := binary.LittleEndian.Uint16(data[8:10]) 
		fmt.Printf("=> altitude :%d\n", int(altitude))
		
		//speed  
		speed := binary.LittleEndian.Uint16(data[10:12]) 
		fmt.Printf("=> speed :%f\n", float64(speed) / 100)
		
		//heartRate 
		heartRate := data[12]
		fmt.Printf("=> heartRate :%d\n", heartRate)
		
		//status 
		status := data[13]	
		fmt.Printf("=> status :%d\n", status)
} 
