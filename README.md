身份证号码解析库，判断号码格式是否正确，并返回包含信息（地区、出生年月日、性别）。

```go
func main() {
    info, err := idnumber.Parse("11010519430606229X") // 由身份证号码生成器生成的号码，仅作示例
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", info)
    /* info 的内容如下：
    {
        Region:{ // 地区
            Code:110105  // 地区行政代码
            Name:朝阳区  // 地区名
            ParentRegions:[ // 上级地区
                {Code:110000 Name:北京市 ParentRegions:[]}
            ]
        } 
        Year:1943  // 年
        Month:6  // 月
        Day:6 // 日
        Gender:1 // 性别 1:男 0:女
    }
    */
}
```