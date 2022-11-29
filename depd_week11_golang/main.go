package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type mahasiswa struct {
	Nim   string `json:"nim"`
	Name  string `json:"name"`
	Prodi string `json:"prodi"`
}

var initData = []mahasiswa{
	{Nim: "07062010001", Name: "Agus", Prodi: "IMT"},
	{Nim: "07062010002", Name: "Ucok", Prodi: "IMT"},
	{Nim: "07062010003", Name: "Cahyo", Prodi: "IMT"},
}

func getMahasiswa(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, initData)
}

func getMhsByNim(nim string) (*mahasiswa, error) {
	for i, t := range initData {
		if t.Nim == nim {
			return &initData[i], nil
		}
	}
	return nil, errors.New("Data not found!")
}

func getMahasiswa2(context *gin.Context) {
	nim := context.Param("nim")
	mahasiswa, err := getMhsByNim(nim)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound,
			gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, mahasiswa)
}

func main() {
	router := gin.Default()
	router.GET("/mahasiswa", getMahasiswa)
	router.GET("/mahasiswa/:nim", getMahasiswa2)
	router.Run("localhost:9090")
}