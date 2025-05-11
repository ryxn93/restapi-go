package productcontrollers

import (
	"gin/database"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	namaProduct := c.PostForm("nama_product")
	deskripsi := c.PostForm("deskripsi")

	// Ambil file gambar dari form
	file, err := c.FormFile("gambar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "File wajib diunggah"})
		return
	}

	//Menyimpan gambar ke direktori lokal
	uploadPath := "/uploads" + file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan gambar"})
		return
	}

	//Menyimpan data produk ke database
	product := models.Product{
		NamaProduct: namaProduct,
		Deskripsi:   deskripsi,
		Gambar:      uploadPath,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan produk "})
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
