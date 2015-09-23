package main

import (
   "os"
   "fmt"
   "log" 
   "extractor/parser"
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
          
   		  bytesGHL := parser.SplitFileToSlice(fileGHL, 48)  		
		  for i := range bytesGHL {
		          parser.ParseGHL(bytesGHL[i]) 
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
          
		  parser.ParseGHT(bytesGHT)	
          
		  // GHP 
		  fmt.Println("\nGHP file...")
          
		  fileGHP, err := os.Open(inputFilePath + ".GHP")
		  if err != nil {
		      log.Fatal("Error while opening file", err)
		  }
          
		  defer fileGHP.Close() 
          
		  bytesGHP := parser.SplitFileToSlice(fileGHP, 20)  		
		  for i := range bytesGHP {
		          parser.ParseGHP(bytesGHP[i]) 
		  		fmt.Println("---")
		  }				
		} else {
   	        log.Fatal("you must specify the path to an input GHL/GHP/GHT file") 
			os.Exit(1)
		}  	     		
} 