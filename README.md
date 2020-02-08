# :memo: zlog
> You don't know what it's like, you don't have a clue
If you did you'd find yourselves doing the same thing too

Sugared logger for console with request_id.

![](image/zlog_gopher.png)

## :traffic_light: Disclaimer

This package heavily uses [tylerb/gls](https://github.com/tylerb/gls) as Goroutin Local Storage implementation which may not seem good practice to some.

## :rocket: Installation

```sh
go get github.com/DiscoFighter47/zlog
```

## :eyeglasses: Examples
**Middleware**
```
package main

import (
    "net/http"

    "github.com/DiscoFighter47/zlog"
)

func main() {
    http.Handle("/", zlog.GenReqID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        zlog.Info("Hello World!")
    })))
    http.ListenAndServe(":3000", nil)
}

// Output: 2020-02-08T14:28:35+06:00 [info ] Hello World! request_id:d9b1e20e-ef83-4aaa-a12a-d696f96e4929
```

**Task Log**
```
package main

import (
    "github.com/DiscoFighter47/zlog"
)

func main() {
    zlog.SetReqID()
    zlog.Info("Hello World!")
}

// Output: 2020-02-08T14:33:54+06:00 [info ] Hello World! request_id:5668c859-5685-40e0-9ccb-99ece2cf7b84
```

## :unlock: License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE.md](LICENSE.md) file for details

## :question: FAQ

1. Why not just use context?
Passing context as the first perimeter is a common norm at [Google](https://about.google/) (and many other tech giants). I'll start using context for logging when Google hires me :smile:.