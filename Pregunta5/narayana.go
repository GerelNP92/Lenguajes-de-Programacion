package main
import(."fmt";."math";."math/big")
func c(n,k int)*Int{r:=NewInt(1);for i:=0;i<k;i++{r.Mul(r,NewInt(int64(n-i)));r.Div(r,NewInt(int64(i+1)))};return r}
func main(){var n int;Scan(&n);k:=int(Log2(float64(n)));N:=c(n,k);N.Mul(N,c(n,k-1));N.Div(N,NewInt(int64(n)));a,b,d:=0,1,2;for i:=3;i<=int(Log2(float64(N.Int64())))+1;i++{a,b,d=b,d,a+b+d};Println(d)}
