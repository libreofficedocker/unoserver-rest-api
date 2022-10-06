package api

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/socheatsok78/unoserver-rest-api/deport"
	"github.com/socheatsok78/unoserver-rest-api/unoconvert"
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

	inFile, _ := os.CreateTemp(deport.WorkDir, tempFilename)
	filePath := inFile.Name()

	// Save file to working directory
	err = c.SaveUploadedFile(form.File, filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	// Prepare output file path
	outFile, _ := os.CreateTemp(deport.WorkDir, tempFilename+"."+form.ConvertTo)

	// Run unoconvert command with optionsq
	err = unoconvert.RunContext(context.Background(), inFile.Name(), outFile.Name(), form.Options...)

	if err != nil {
		log.Printf("unoconvert error: %s", err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("unoconvert error: %s", err))
		return
	}

	// Send the converted file to body stream
	c.File(outFile.Name())
}
