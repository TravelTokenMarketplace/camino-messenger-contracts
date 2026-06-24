import { Route, Routes } from "react-router-dom";
import { Layout } from "./components/Layout";
import { Dashboard } from "./pages/Dashboard";
import { CreateAccount } from "./pages/CreateAccount";
import { AccountWorkspace } from "./pages/AccountWorkspace";
import { ManagerWorkspace } from "./pages/ManagerWorkspace";

export default function App() {
  return (
    <Routes>
      <Route element={<Layout />}>
        <Route index element={<Dashboard />} />
        <Route path="create" element={<CreateAccount />} />
        <Route path="manager" element={<ManagerWorkspace />} />
        <Route path="account/:address" element={<AccountWorkspace />} />
      </Route>
    </Routes>
  );
}
