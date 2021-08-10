package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech"
   "net/http"
   "net/http/httputil"
   "os"
)

func main() {
   var (
      https bool
      indent, output string
   )
   flag.BoolVar(&https, "s", false, "HTTPS")
   flag.StringVar(&indent, "i", "", "indent")
   flag.StringVar(&output, "o", "", "output file")
   flag.Parse()
   if flag.NArg() != 1 {
      fmt.Println("mech [flags] [request file]")
      flag.PrintDefaults()
      return
   }
   file := flag.Arg(0)
   rd, err := os.Open(file)
   if err != nil {
      panic(err)
   }
   defer rd.Close()
   req, err := mech.ReadRequest(rd)
   if err != nil {
      panic(err)
   }
   if https {
      req.URL.Scheme = "https"
   } else {
      req.URL.Scheme = "http"
   }
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   d, err := httputil.DumpResponse(res, false)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(d)
   // stdout
   if indent == "" && output == "" {
      os.Stdout.ReadFrom(res.Body)
      return
   }
   // indent stdout
   if output == "" {
      e := mech.NewEncoder(os.Stdout)
      e.SetIndent(indent)
      e.Encode(res.Body)
      return
   }
   // file
   wr, err := os.Create(output)
   if err != nil {
      panic(err)
   }
   defer wr.Close()
   if indent == "" {
      wr.ReadFrom(res.Body)
      return
   }
   // indent file
   e := mech.NewEncoder(wr)
   e.SetIndent(indent)
   e.Encode(res.Body)
}
