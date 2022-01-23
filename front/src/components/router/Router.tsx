import { memo, VFC } from "react";
import { Route, Routes } from "react-router-dom";

import { CreateWorkspace } from "../pages/CreateWorkspace";
import { Login } from "../pages/Login";
import { LoginSuccess } from "../pages/LoginSuccess";
import { Page404 } from "../pages/Page404";

export const Router: VFC = memo(() => {
  return (
    <Routes>
      <Route
        path="/create_workspace"
        element={
          <>
            <CreateWorkspace />
          </>
        }
      />
      <Route
        path="/login"
        element={
          <>
            <Login />
          </>
        }
      />
      <Route
        path="/login_success"
        element={
          <>
            <LoginSuccess />
          </>
        }
      />
      <Route path="/*" element={<Page404 />} />
    </Routes>
  );
});
