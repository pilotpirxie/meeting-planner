import { createBrowserRouter, RouterProvider } from "react-router";

import { OutletContainer } from "./components/OutletContainer";

const router = createBrowserRouter([
  {
    path: "/",
    element: <OutletContainer />,
    children: [
      {
        path: "/",
        element: <div>Home Page</div>,
      }
    ],
  },
]);

export function App() {
  return <RouterProvider router={router} />;
}
