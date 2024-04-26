// This package contains the operations for managing elements in a system
// It provides
package operations

import (
        "fmt"

        "net/http"
        "hash/fnv"
)

const STDINTENTS_SHA256 = "c53093c061d9fb514c414771f4a7036a6d18f44cfc029f7d3a09108ac4d77c1b"
const STDINTENTS_LOCATION = "https://gitlab.jyu.fi/ijoliver/jane/-/tree/main/etc/standardintents.json"

func calculateHash(b []byte) uint32 {
        h := fnv.New32a()
        h.Write(b)
        return h.Sum32()
}

func LoadStandardIntents() (string,string,error) { 

   // read the file

   resp, err := http.Get( STDINTENTS_LOCATION )
   if err != nil {
        return err.Error(), "1", err 
   }
   defer resp.Body.Close()

   contents := make([]byte, resp.ContentLength)
   _,err = resp.Body.Read(contents)
   if err != nil {
        return err.Error(), "2", err 
   }

   // calculate the HASH

   h := calculateHash( contents )

   fmt.Printf("Body is %v\n",contents)
   fmt.Printf("Hash is %v\n",h)

   return "yes","maybe",nil

}