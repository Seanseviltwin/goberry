package web

import (
	"github.com/alecthomas/gobundle"
)

var GoberryBundle = gobundle.NewBuilder("web").Compressed().UncompressOnInit().Add(
	"api.raml", []byte("x\xdat\x8e=\x8b\xc30\f@w\xfd\n\xc1q\xe3\xc5w\xdb\xe1\xad\xd0n\xfd\x80\x94vw\x12\xd5\x11\xd8Q\xb0\x95\x94\xfc\xfb:\xb4\x1d;\t\x9e\x9ex\xfa\xfa\xae7\x87=\xfeV\xff\xa0\xac\x81,\xfaD7\x98)e\x96\xc1\xe2\x1f@\xe32]\x12[\xecUGk\x8cg\xed\xa7\xa6j%\x9a]+\x83D\u03ba\xe5B]8K\x98\xb4\xdce\u3961\x94\x16\u891d\"\r\xeaVl\x01\xf1\a\xf1\x15:\x95\xc8\xcct\a0\xeb\u0093\xae\x03\xb1\xe3<\x06\xb7\x1c],R-\xa2`\xde\xef|\u052eO\xe1\x11\x00\x00\xff\xff\x93\xc6D\x0e"),
).Build()