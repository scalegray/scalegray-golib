package elasticsearch


import (
  "fmt"
  "github.com/packetbeat/elastigo/api"
 "github.com/packetbeat/elastigo/core"
 )

 //need to configure it from conf file.
//api.Port

func Index(data []byte){

  api.Domain = "localhost"
  fmt.Println("Entered inside elasticgo file...lets do this")
  response, _ := core.Index("scalegraydata", "sg-type", "3", nil, data)
   fmt.Println(response)
  fmt.Println("Done pushing the data into elastic search..woohoo!")
}
