package main

import "./minimesos"

func main() {
  var mc minimesos.MesosCluster
  mc.Create()
  mc.Start()
  mc.Destroy()
}

