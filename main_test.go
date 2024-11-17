package main

import (
	"testing"
)

func TestReadConfigFromConfigFile(t *testing.T) {

	ReadConfigs(true, "")

}

// this test verifies the correct processing of the durationUnits `.d` and `.min`
func TestReadConfigFromText(t *testing.T) {
	text := `
	doc_confs = new Mapping {
    ["test"] = new Mapping {
        ["cap1"] {
            configs = new Listing {
            new {
                name = "PRE_PROCESS"
                val = true
                type = "boolean"
            }
            new {
                name = "DOC_MAX_SIZE"
                val = 80.mib
                type = "size"
            }
            new {
                name = "DOC_MIMES"
                val = "pdf,docx,jpg,heic"
                type = "csv"
            }
            new {
                name = "doc_timeout"
                val = 15.min
                type = "duration"
            }
            }
        }
        ["cap2"] {
            configs = new Listing {
            new {
                name = "PRE_PROCESS"
                val = true
                type = "boolean"
            }
            new {
                name = "DOC_MAX_SIZE"
                val = 20.mb
                type = "size"
            }
            new {
                name = "DOC_MIMES"
                val = "pdf,docx,jpg,heic"
                type = "csv"
            }
            }
        }
		 ["cap3"] {
            configs = new Listing {
            new {
                name = "PRE_PROCESS"
                val = true
                type = "boolean"
            }
            new {
                name = "DOC_MAX_SIZE"
                val = 20.gib
                type = "size"
            }
            new {
                name = "DOC_MIMES"
                val = "pdf,docx,jpg,heic"
                type = "csv"
            }
			new {
                name = "doc_timeout"
                val = 15.d
                type = "duration"
            }
            }
        }
    }
}
foo = "FOO"

output {
  renderer = new JsonRenderer {
    converters {
      [DataSize] = (size) -> "\(size.value)\(size.unit)"
      [Duration] = (dur) -> "\(dur.value)\(dur.unit)"
    }
  }
}
`

	ReadConfigs(false, text)

}

func TestReadConfigFromInvalidPkl(t *testing.T) {
	text := `
	doc_confs = new Mapping {
    ["test"] = new Mapping {
        ["cap1"] {
            configs = new Listing {
            new {
                name = "PRE_PROCESS"
                val = true
                type = "boolean"
            }
            new {
                name = "DOC_MAX_SIZE"
                val = 80.ppp
                type = "size"
            }
            new {
                name = "DOC_MIMES"
                val = "pdf,docx,jpg,heic"
                type = "csv"
            }
            new {
                name = "doc_timeout"
                val = 15.min
                type = "duration"
            }
            }
        }
        ["cap2"] {
            configs = new Listing {
            new {
                name = "PRE_PROCESS"
                val = true
                type = "boolean"
            }
            new {
                name = "DOC_MAX_SIZE"
                val = 20.mib
                type = "size"
            }
            new {
                name = "DOC_MIMES"
                val = "pdf,docx,jpg,heic"
                type = "csv"
            }
            }
        }
		 ["cap3"] {
            configs = new Listing {
            new {
                name = "PRE_PROCESS"
                val = true
                type = "boolean"
            }
            new {
                name = "DOC_MAX_SIZE"
                val = 20.gib
                type = "size"
            }
            new {
                name = "DOC_MIMES"
                val = "pdf,docx,jpg,heic"
                type = "csv"
            }
			new {
                name = "doc_timeout"
                val = 15.d
                type = "duration"
            }
            }
        }
    }
}
foo = "FOO"

output {
  renderer = new JsonRenderer {
    converters {
      [DataSize] = (size) -> "\(size.value)\(size.unit)"
      [Duration] = (dur) -> "\(dur.value)\(dur.unit)"
    }
  }
}
`
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	ReadConfigs(false, text)

}
