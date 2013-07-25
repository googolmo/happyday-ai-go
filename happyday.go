package main

import (
    "github.com/hoisie/web"
    "encoding/json"
    "fmt"
    "strings"
)


func index(ctx *web.Context) {
    ctx.WriteString(`<html>
    <head>
    <title>HappyDay AI</title>
    <style>
    body {font-family: 'Source Code Pro', sans-serif;}
    </style>
    </head>
    <body>
    <h1>HappyDay AI player Go</h1>
    <script type="text/javascript">
    WebFontConfig = {
        google: { families: [ 'Source+Code+Pro:600:latin' ] }
    };
    (function() {
        var wf = document.createElement('script');
        wf.src = ('https:' == document.location.protocol ? 'https' : 'http') +
        '://ajax.googleapis.com/ajax/libs/webfont/1/webfont.js';
        wf.type = 'text/javascript';
        wf.async = 'true';
        var s = document.getElementsByTagName('script')[0];
        s.parentNode.insertBefore(wf, s);
    })(); </script>
    </body>
    </html>`)
}

func begin(ctx *web.Context) {
    for k, v := range ctx.Params {
        println(k, v)
    }
    ctx.SetHeader("Content-Type", "application/json", true)
    ctx.WriteString("")
}

func next(ctx *web.Context) []byte{
    for k, v := range ctx.Params {
        fmt.Println(k, v)
        println(k, v)
    }
    grid := ctx.Params["grid"]
    step := ctx.Params["step"]
    println(grid, step)
    resultMap := make(map[string]interface{})

    resultMap["direction"] = cal(grid, step)
    b , _ := json.Marshal(resultMap)
    ctx.SetHeader("Content-Type", "application/json", true)
    return b
}


func result(ctx *web.Context, val string) {
    for k, v := range ctx.Params {
        println(k, v)
    }
    ctx.SetHeader("Content-Type", "application/json", true)
    ctx.WriteString("")
}

func cal(grid string, step string) int{
    fmt.Println(step)
    ss := strings.Split(grid, ",")
    fmt.Println(ss)
    s0 := strings.Split(ss[0], "")
    s1 := strings.Split(ss[1], "")
    s2 := strings.Split(ss[2], "")
    up := s0[1]
    left := s1[0]
    right := s1[2]
    down := s2[1]
    fmt.Println(up)
    fmt.Println(down)
    fmt.Println(left)
    fmt.Println(right)
    dd := []string{up, down, right, left}
    fmt.Println(dd)
    if (left == "2" || left == "3") {
        return 0
    }
    if (up == "2" || up == "3") {
        return 1
    }
    if (right == "2" || right == "3") {
        return 2
    }
    if (down == "2" || down == "3") {
        return 3
    }
    if (up == "4" || up == "9"){
        return getNotFour(dd, 1)
    }
    if (down == "4" || down == "9"){
        return getNotFour(dd, 3)
    }
    if (left == "4" || left == "9"){
		return getNotFour(dd, 0)
    }
    if (right == "4" || right == "9") {
        return getNotFour(dd, 2)
    }
    if (up == "0"){
        return 1
    }
    if (down == "0" ){
        return 3
    }
    if (left == "0"){
		return 0
    }
    if (right == "0") {
        return 2
    }
    return 1
}

func getNotFour(direct []string, d int) int {
    if (direct[1] != "4" && direct[1] != "9") {
        return 3
    }
    if (direct[3] != "4" && direct[3] != "9") {
        return 0
    }
    if (direct[2] != "4" && direct[2] != "9" ) {
        return 2
    }
    if (direct[0] != "4" && direct[0] != "9") {
        return 1
    }
    return d
}

func main() {
    web.Get("/", index)
    web.Post("/begin", begin)
    web.Post("/next", next)
    web.Post("/result", result)
    web.Run("0.0.0.0:9999")
}
