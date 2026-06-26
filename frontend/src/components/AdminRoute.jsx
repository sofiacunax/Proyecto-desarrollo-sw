import { Navigate } from "react-router-dom"

export default function AdminRoute({ children }) {

  const token = localStorage.getItem("token")

  if (!token) {
    return <Navigate to="/" />
  }

  let payload

  try {
    payload = JSON.parse(
      atob(token.split(".")[1])
    )
  } catch {
    return <Navigate to="/" />
  }

  if (payload.rol !== "ADMIN") {
    return <Navigate to="/dashboard" />
  }

  return children
}