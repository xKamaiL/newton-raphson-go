# newton-raphson
newton raphson theory implement in golang


## How to find the value of x ?

$$0=sin(x-0.361) + sin(0.361)*exp(-x/0.377)$$

```go
math.Sin(x-0.361) + (math.Sin(0.361) * math.Exp(-x/0.377))
```
