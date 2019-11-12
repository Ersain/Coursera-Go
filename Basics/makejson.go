package main

import(
  "os"
  "fmt"
  "bufio"
  "encoding/json"
)


func readString(reader *bufio.Reader) string{
  text, _ := reader.ReadString('\n') 
  return text   
}

func main(){
   reader := bufio.NewReader(os.Stdin)
   var name, address string   
   var input map[string]string 
   input = make(map[string]string)

   fmt.Println("Enter name:")
   name = readString(reader)
   fmt.Println("Enter address")
   address = readString(reader)

   input["name"] = name[:len(name)-1] 
   input["address"] = address[:len(address)-1]
 
  jsonBytes, _ := json.Marshal(input)
  jsonStr := string(jsonBytes)

  fmt.Println(jsonStr)
   
}
