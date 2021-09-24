# :memo: zlog
> You don't know what it's like, you don't have a clue  
If you did you'd find yourselves doing the same thing too

Sugared logger for console with request id.

![](image/zlog_gopher.png)  
*source: [gopherize.me](https://gopherize.me)*

## :traffic_light: Disclaimer

This package heavily uses [tylerb/gls](https://github.com/tylerb/gls) as Goroutin Local Storage implementation which may not seem good practice to some.

## :rocket: Installation

```sh
go get github.com/DiscoFighter47/zlog
```

## :eyeglasses: Examples
```
package main

import (
	"net/http"

	"github.com/DiscoFighter47/zlog"
)

func call() {
	zlog.Info("Hello Universe!")
}

func main() {
	http.Handle("/", zlog.GenReqID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zlog.Info("Hello World!")
		call()
	})))
	http.ListenAndServe(":3000", nil)
}

// Output: 2020-02-08T19:54:10+06:00 [info ] Hello World! request_id:92fb565f-3856-4a51-a0e9-744a5fad12fc
// Output: 2020-02-08T19:54:10+06:00 [info ] Hello Universe! request_id:92fb565f-3856-4a51-a0e9-744a5fad12fc
```

## :unlock: License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE.md](LICENSE) file for details

## :question: FAQ

1. Why not just use context?  
Passing context as the first parameter is a common norm at [Google](https://about.google) (and many other tech giants). I'll start using context for logging when Google hires me :smile:.
