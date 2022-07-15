package main

import("fmt")


type app struct{}

func(app *app) Run()error{
   fmt.Println("setting up")
   return nil

}

func main(){
   app:= app{}
   if err:= app.Run;err!=nil{
      return fmt.Println("error")
   }
}
