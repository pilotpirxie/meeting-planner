import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const emptyApi = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: import.meta.env.DEV
      ? "http://localhost:8080/api"
      : "https://api.example.com",
  }),
  endpoints: () => ({}),
});