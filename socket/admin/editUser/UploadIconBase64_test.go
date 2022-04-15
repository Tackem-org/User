package editUser_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/User/model"
	"github.com/Tackem-org/User/socket/admin/editUser"
	"github.com/stretchr/testify/assert"
)

var tgif = "data:image/gif;base64,R0lGODlhQABAAIABAP8AAP///yH+EUNyZWF0ZWQgd2l0aCBHSU1QACwAAAAAQABAAAACRYSPqcvtD6OctNqLs968+w+G4kiW5omm6sq27gvH8kzX9o3n+s73/g8MCofEovGITCqXzKbzCY1Kp9Sq9YrNarfcrhdQAAA7"
var tjpg = "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEBLAEsAAD//gATQ3JlYXRlZCB3aXRoIEdJTVD/4gKwSUNDX1BST0ZJTEUAAQEAAAKgbGNtcwQwAABtbnRyUkdCIFhZWiAH5gAEAA4AFwAeABNhY3NwQVBQTAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA9tYAAQAAAADTLWxjbXMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA1kZXNjAAABIAAAAEBjcHJ0AAABYAAAADZ3dHB0AAABmAAAABRjaGFkAAABrAAAACxyWFlaAAAB2AAAABRiWFlaAAAB7AAAABRnWFlaAAACAAAAABRyVFJDAAACFAAAACBnVFJDAAACFAAAACBiVFJDAAACFAAAACBjaHJtAAACNAAAACRkbW5kAAACWAAAACRkbWRkAAACfAAAACRtbHVjAAAAAAAAAAEAAAAMZW5VUwAAACQAAAAcAEcASQBNAFAAIABiAHUAaQBsAHQALQBpAG4AIABzAFIARwBCbWx1YwAAAAAAAAABAAAADGVuVVMAAAAaAAAAHABQAHUAYgBsAGkAYwAgAEQAbwBtAGEAaQBuAABYWVogAAAAAAAA9tYAAQAAAADTLXNmMzIAAAAAAAEMQgAABd7///MlAAAHkwAA/ZD///uh///9ogAAA9wAAMBuWFlaIAAAAAAAAG+gAAA49QAAA5BYWVogAAAAAAAAJJ8AAA+EAAC2xFhZWiAAAAAAAABilwAAt4cAABjZcGFyYQAAAAAAAwAAAAJmZgAA8qcAAA1ZAAAT0AAACltjaHJtAAAAAAADAAAAAKPXAABUfAAATM0AAJmaAAAmZwAAD1xtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAAgAAAAcAEcASQBNAFBtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAAgAAAAcAHMAUgBHAEL/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wgARCABAAEADAREAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAWAQEBAQAAAAAAAAAAAAAAAAAABgj/2gAMAwEAAhADEAAAAZzC6pAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAH//xAAUEAEAAAAAAAAAAAAAAAAAAABg/9oACAEBAAEFAgH/xAAUEQEAAAAAAAAAAAAAAAAAAABg/9oACAEDAQE/AQH/xAAUEQEAAAAAAAAAAAAAAAAAAABg/9oACAECAQE/AQH/xAAUEAEAAAAAAAAAAAAAAAAAAABg/9oACAEBAAY/AgH/xAAUEAEAAAAAAAAAAAAAAAAAAABg/9oACAEBAAE/IQH/2gAMAwEAAgADAAAAEP8A/wD/AP8A/wD/AP8A/wD/AP8A/wD/AP8A/wD/AP8A/wD/AP8A/wD/AP8A/wD/AP/EABQRAQAAAAAAAAAAAAAAAAAAAGD/2gAIAQMBAT8QAf/EABQRAQAAAAAAAAAAAAAAAAAAAGD/2gAIAQIBAT8QAf/EABQQAQAAAAAAAAAAAAAAAAAAAGD/2gAIAQEAAT8QAf/Z"
var tpng = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAIAAAAlC+aJAAABhWlDQ1BJQ0MgcHJvZmlsZQAAKJF9kT1Iw0AcxV8/pCJVhxYUEclQnSyIijhqFYpQIdQKrTqYXPohNGlIUlwcBdeCgx+LVQcXZ10dXAVB8APE0clJ0UVK/F9SaBHjwXE/3t173L0D/PUyU83gGKBqlpFOJoRsbkUIvSKIHvRhCBGJmfqsKKbgOb7u4ePrXZxneZ/7c3QreZMBPoF4humGRbxOPLVp6Zz3iaOsJCnE58SjBl2Q+JHrsstvnIsO+3lm1Mik54ijxEKxjeU2ZiVDJZ4kjimqRvn+rMsK5y3OarnKmvfkLwznteUlrtMcRBILWIQIATKq2EAZFuK0aqSYSNN+wsM/4PhFcsnk2gAjxzwqUCE5fvA/+N2tWZgYd5PCCaDjxbY/hoHQLtCo2fb3sW03ToDAM3CltfyVOjD9SXqtpcWOgN5t4OK6pcl7wOUO0P+kS4bkSAGa/kIBeD+jb8oBkVuga9XtrbmP0wcgQ12lboCDQ2CkSNlrHu/ubO/t3zPN/n4AUCtymWrqe9cAAAAJcEhZcwAALiMAAC4jAXilP3YAAAAHdElNRQfmBA4XIBZhr6fKAAAAGXRFWHRDb21tZW50AENyZWF0ZWQgd2l0aCBHSU1QV4EOFwAAAExJREFUaN7tz0ENAAAIBKDT/p01gm83aEBNfusICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBwWUoUBf8cBNPQAAAAASUVORK5CYII="

func TestUploadIconBase64(t *testing.T) {
	assert.NotPanics(t, func() { model.Setup("testUploadIconBase64.db") })
	defer os.Remove("testUploadIconBase64.db")

	r1, err1 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{},
	})
	assert.IsType(t, &structs.SocketReturn{}, r1)
	assert.Nil(t, err1)
	assert.Equal(t, http.StatusNotAcceptable, int(r1.StatusCode))
	assert.Equal(t, "userid missing", r1.ErrorMessage)

	r2, err2 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": "fail",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r2)
	assert.Nil(t, err2)
	assert.Equal(t, http.StatusNotAcceptable, int(r2.StatusCode))
	assert.Equal(t, "userid not an int", r2.ErrorMessage)

	r3, err3 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 30,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r3)
	assert.Nil(t, err3)
	assert.Equal(t, http.StatusNotFound, int(r3.StatusCode))
	assert.Equal(t, "user not found", r3.ErrorMessage)

	r4, err4 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r4)
	assert.Nil(t, err4)
	assert.Equal(t, http.StatusNotAcceptable, int(r4.StatusCode))
	assert.Equal(t, "icon missing", r4.ErrorMessage)

	r5, err5 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
			"icon":   1,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r5)
	assert.Nil(t, err5)
	assert.Equal(t, http.StatusBadRequest, int(r5.StatusCode))
	assert.Equal(t, "icon not a string", r5.ErrorMessage)

	r6, err6 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
			"icon":   "notbase64",
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r6)
	assert.Nil(t, err6)
	assert.Equal(t, http.StatusBadRequest, int(r6.StatusCode))
	assert.Equal(t, "icon not base64 encoded", r6.ErrorMessage)

	r7, err7 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
			"icon":   tpng,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r7)
	assert.Nil(t, err7)
	assert.Equal(t, http.StatusOK, int(r7.StatusCode))
	assert.Equal(t, "", r7.ErrorMessage)

	r8, err8 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
			"icon":   tjpg,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r8)
	assert.Nil(t, err8)
	assert.Equal(t, http.StatusOK, int(r8.StatusCode))
	assert.Equal(t, "", r8.ErrorMessage)

	r9, err9 := editUser.UploadIconBase64(&structs.SocketRequest{
		Data: map[string]interface{}{
			"userid": 2,
			"icon":   tgif,
		},
	})
	assert.IsType(t, &structs.SocketReturn{}, r9)
	assert.Nil(t, err9)
	assert.Equal(t, http.StatusOK, int(r9.StatusCode))
	assert.Equal(t, "", r9.ErrorMessage)

}
