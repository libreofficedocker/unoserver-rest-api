package api

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/libreoffice-docker/unoserver-rest-api/depot"
	"github.com/libreoffice-docker/unoserver-rest-api/unoconvert"
)

type RequestForm struct {
	Name      string                `form:"name"`
	Options   []string              `form:"opts[]"`
	ConvertTo string                `form:"convert-to" binding:"required"`
	File      *multipart.FileHeader `form:"file" binding:"required"`
}

func RequestHandler(c *gin.Context) {
	var err error
	var form RequestForm

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var tempFilename = "*"

	if form.Name == "" {
		form.Name = form.File.Filename
	}

	tempFilename += "-" + form.Name

	inFile, err := os.CreateTemp(depot.WorkDir, tempFilename)
	if err != nil {
		log.Println("Create temp file failed", err)
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}
	filePath := inFile.Name()
	defer func() {
		err := os.Remove(filePath)
		if err != nil {
			log.Println("Delege temp file failed", err)
		}
	}()

	// Save file to working directory
	err = c.SaveUploadedFile(form.File, filePath)
	if err != nil {
		log.Println("Convert failed", err)
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	// Prepare output file path
	outFile, err := os.CreateTemp(depot.WorkDir, tempFilename+"."+form.ConvertTo)
	if err != nil {
		log.Println("Create temp file failed", err)
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}
	defer func() {
		err := os.Remove(outFile.Name())
		if err != nil {
			log.Println("Delege temp file failed", err)
		}
	}()

	// Run unoconvert command with options
	// If context timeout is 0s run without timeout
	if unoconvert.ContextTimeout == 0 {
		err = unoconvert.Run(inFile.Name(), outFile.Name(), form.Options...)
	} else {
		err = unoconvert.RunContext(context.Background(), inFile.Name(), outFile.Name(), form.Options...)
	}

	if err != nil {
		log.Printf("unoconvert error: %s", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("unoconvert error: %s", err))
		return
	}

	// Send the converted file to body stream
	c.File(outFile.Name())
}
