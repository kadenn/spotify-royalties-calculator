import { createSlice } from '@reduxjs/toolkit';
import { apiCallBegan } from './api';

const initialState = {
  data: [],
  errorMessage: "",
  loading: false,
};

const tracks = createSlice({
  name: "tracks",
  initialState,
  reducers: {
    requested: (state) => ({
      ...state,
      loading: true,
    }),
    success: (state, action) => ({
      ...state,
      loading: false,
      data: state.data.concat(action.payload),
    }),
    failed: (state, action) => ({
      ...state,
      loading: false,
      errorMessage: action.payload,
    }),
  },
});

export default tracks.reducer;

export const { requested, success, failed } = tracks.actions;

export const appendTrack = (id) =>
  apiCallBegan({
    url: `${process.env.REACT_APP_API_URL}/track/${id}`,
    onStart: requested.type,
    onSuccess: success.type,
    onError: failed.type,
  });
