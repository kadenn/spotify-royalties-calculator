import { createSlice } from '@reduxjs/toolkit';
import { apiCallBegan } from './api';

const initialState = {
  data: [],
  errorMessage: "",
  loading: false,
};

const playlists = createSlice({
  name: "playlists",
  initialState,
  reducers: {
    requested: (state) => ({
      ...state,
      loading: true,
    }),
    success: (state, action) => ({
      ...state,
      loading: false,
      data: action.payload,
    }),
    failed: (state, action) => ({
      ...state,
      loading: false,
      errorMessage: action.payload,
    }),
  },
});

export default playlists.reducer;

export const { requested, success, failed } = playlists.actions;

export const loadPlaylists = () =>
  apiCallBegan({
    url: `${process.env.REACT_APP_API_URL}/featured_playlists`,
    onStart: requested.type,
    onSuccess: success.type,
    onError: failed.type,
  });

