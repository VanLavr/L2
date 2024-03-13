package main

import "fmt"

func main() {
	facade := new(ImageManager)

	facade.fileServer = new(FileServer)
	facade.metadataManager = new(MetadataManager)
	facade.dataagregator = new(DataAgregator)

	facade.SaveImage()
	facade.GetImage()
	facade.DeleteImage()
}

type MetadataManager struct{}

func (m MetadataManager) SaveMetadata() {
	fmt.Println("Saving metadata...")
}
func (m MetadataManager) DeleteMetadata() {
	fmt.Println("Deleting metadata...")
}
func (m MetadataManager) GetMetadata() {
	fmt.Println("Getting metadata...")
}

type FileServer struct{}

func (f FileServer) SaveImage() {
	fmt.Println("Saving image")
}
func (f FileServer) DeleteImage() {
	fmt.Println("Deleting image")
}
func (f FileServer) GetImage() {
	fmt.Println("Getting image")
}

type ImageManager struct {
	fileServer      *FileServer
	metadataManager *MetadataManager
	dataagregator   *DataAgregator
}

func (i ImageManager) SaveImage() {
	i.metadataManager.SaveMetadata()
	i.fileServer.SaveImage()
	i.dataagregator.SendTimeStamp()
}
func (i ImageManager) DeleteImage() {
	i.metadataManager.DeleteMetadata()
	i.fileServer.DeleteImage()
}
func (i ImageManager) GetImage() {
	i.metadataManager.GetMetadata()
	i.fileServer.GetImage()
	i.dataagregator.UpdateViewStatistics()
}

type DataAgregator struct{}

func (d DataAgregator) SendTimeStamp() {
	fmt.Println("Sending timestamp...")
}
func (d DataAgregator) UpdateViewStatistics() {
	fmt.Println("Updating view statistics...")
}
