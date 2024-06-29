module github.com/torbenish/pos-golang/05-Packaging/3-Replace/system

go 1.22.4

replace github.com/torbenish/pos-golang/05-Packaging/3-Replace/math => ../math

require github.com/torbenish/pos-golang/05-Packaging/3-Replace/math v0.0.0-00010101000000-000000000000
