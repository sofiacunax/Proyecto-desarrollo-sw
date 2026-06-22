import { Routes, Route, Navigate } from "react-router-dom"
import Dashboard from "./pages/Dashboard"
import ProtectedRoute from "./components/ProtectedRoute"
import CompraEntrada from "./pages/CompraEntrada"

import Login from "./pages/Login"
import Register from "./pages/Register"
import MisEntradas from "./pages/MisEntradas"
import AdminEventos from "./pages/AdminEventos";
import AdminRoute from "./components/AdminRoute"
import CrearEvento from "./pages/CrearEvento"
import EditarEvento from "./pages/EditarEvento"
import AdminUsuarios from "./pages/AdminUsuarios"

function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route
  path="/dashboard"
  element={
    <ProtectedRoute>
      <Dashboard />
    </ProtectedRoute>
  }
/>
<Route
  path="/mis-entradas"
  element={
    <ProtectedRoute>
      <MisEntradas />
    </ProtectedRoute>
  }
/>
<Route
  path="/admin/eventos"
  element={
    <AdminRoute>
      <AdminEventos />
    </AdminRoute>
  }
/>
<Route
  path="/admin/usuarios"
  element={
    <AdminRoute>
      <AdminUsuarios />
    </AdminRoute>
  }
/>
<Route
  path="/admin/eventos/nuevo"
  element={
    <AdminRoute>
      <CrearEvento />
    </AdminRoute>
  }
/>
<Route
  path="/admin/eventos/editar/:id"
  element={
    <AdminRoute>
      <EditarEvento />
    </AdminRoute>
  }
/>
<Route
  path="/comprar/:id"
  element={
    <ProtectedRoute>
      <CompraEntrada />
    </ProtectedRoute>
  }
/>
<Route
  path="*"
  element={<Navigate to="/dashboard" />}
/>

    </Routes>
  )
  
}


export default App