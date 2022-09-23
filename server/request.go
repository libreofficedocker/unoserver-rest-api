package server

import (
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/socheatsok78/unoserver-rest-api/deport"
	"github.com/socheatsok78/unoserver-rest-api/unoconvert"
)

type RequestForm struct {
	Name        string                `form:"name"`
	Orientation string                `form:"orientation"`
	ConvertTo   string                `form:"convert-to" binding:"required"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
}

func RequestHandler(c *gin.Context) {
	var err error
	var form RequestForm

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if form.Name != "" {
		form.Name = form.File.Filename
	}

	inFile, _ := os.CreateTemp(deport.WorkDir, "*-"+form.Name)
	filePath := inFile.Name()

	// Save file to working directory
	err = c.SaveUploadedFile(form.File, filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "unknown error")
		return
	}

	// Prepare output file path
	outFile, _ := os.CreateTemp(deport.WorkDir, "*-"+form.Name+"."+form.ConvertTo)

	// Prepare unoconvert comamnd flags
	args := []string{}
	if form.Orientation == "landscape" {
		args = append(args, "--landscape")
	}

	// Run unoconvert command with options
	err = unoconvert.Run(inFile.Name(), outFile.Name(), args...)
	if err != nil {
		c.String(http.StatusInternalServerError, "unoconvert unknown error")
		return
	}

	// Send the converted file to body stream
	c.File(outFile.Name())
}
