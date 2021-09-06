import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit';
import api from './middleware/api';
import playlists from './playlists';
import tracks from './tracks';

export const store = configureStore({
  reducer: {
    playlists: playlists,
    tracks: tracks,
  },
  middleware: [...getDefaultMiddleware(), api],
});