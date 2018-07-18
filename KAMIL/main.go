package main
import."fmt"
func main(){a:=""
for{if _,e:=Scan(&a);e!=nil{break}
r:=1
for _,v:=range a{switch v{case 68,70,76,84:r*=2}}
Println(r)}}