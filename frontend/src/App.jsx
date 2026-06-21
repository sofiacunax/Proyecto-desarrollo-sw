import { Routes, Route } from "react-router-dom"
import Dashboard from "./pages/Dashboard"
import ProtectedRoute from "./components/ProtectedRoute"
import CompraEntrada from "./pages/CompraEntrada"

import Login from "./pages/Login"
import Register from "./pages/Register"
import MisEntradas from "./pages/MisEntradas"

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
  path="/comprar/:id"
  element={
    <ProtectedRoute>
      <CompraEntrada />
    </ProtectedRoute>
  }
/>

    </Routes>
  )
  
}


export default App