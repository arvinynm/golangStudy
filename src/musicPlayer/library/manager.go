package library

import "musicPlayer/model"

type MusicManager struct {
	musics []model.MusicEntry
}

func NewMusicManager() *MusicManager{
	return &MusicManager{make([]model.MusicEntry, 0)}
}