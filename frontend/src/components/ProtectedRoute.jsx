export default function ProtectedRoute({ children }) {
  const isAuthenticated = true

  if (!isAuthenticated) {
    return <h1>No autorizado</h1>
  }

  return children
}