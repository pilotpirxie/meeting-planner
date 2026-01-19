import { createSlice, type PayloadAction } from "@reduxjs/toolkit";

export type Theme = "light" | "dark";

interface ConfigState {
  theme: Theme;
}

const initialState: ConfigState = {
  theme: "light",
};

const configSlice = createSlice({
  name: "config",
  initialState,
  reducers: {
    setTheme(state, action: PayloadAction<Theme>) {
      state.theme = action.payload;
    },
  },
});

export const {
  setTheme
} = configSlice.actions;

export default configSlice.reducer;