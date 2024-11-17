# PKL Configuration Loader
This is a simple POC to see how feasible the Go language bindings are 
to use for configuration parsing.

## How to run
```bash
➜  pkl-conf-loader git:(main) ✗ make run
go run main.go
{
  "doc_confs": {
    "test": {
      "cap1": {
        "configs": [
          {
            "name": "PRE_PROCESS",
            "val": true,
            "type": "boolean"
          },
          {
            "name": "DOC_MAX_SIZE",
            "val": "80mib",
            "type": "size"
          },
          {
            "name": "DOC_MIMES",
            "val": "pdf,docx,jpg,heic",
            "type": "csv"
          },
          {
            "name": "doc_timeout",
            "val": "15min",
            "type": "duration"
          }
        ]
      },
      "cap2": {
        "configs": [
          {
            "name": "PRE_PROCESS",
            "val": true,
            "type": "boolean"
          },
          {
            "name": "DOC_MAX_SIZE",
            "val": "20mib",
            "type": "size"
          },
          {
            "name": "DOC_MIMES",
            "val": "pdf,docx,jpg,heic",
            "type": "csv"
          }
        ]
      }
    }
  },
  "foo": "FOO"
}

doc config:  {Foo:FOO DocConfigs:map[test:map[cap1:{Configs:[{Name:PRE_PROCESS Value:true ElemType:boolean} {Name:DOC_MAX_SIZE Value:80mib ElemType:size} {Name:DOC_MIMES Value:pdf,docx,jpg,heic ElemType:csv} {Name:doc_timeout Value:15min ElemType:duration}]} cap2:{Configs:[{Name:PRE_PROCESS Value:true ElemType:boolean} {Name:DOC_MAX_SIZE Value:20mib ElemType:size} {Name:DOC_MIMES Value:pdf,docx,jpg,heic ElemType:csv}]}]]}
true
from human size: 80000000
from human size to bytes:  83886080
80mib
pdf,docx,jpg,heic
duration: 15m0s
```
