const API_URL = "http://localhost:8080"

export async function getEventos(search = "") {
  const response = await fetch(
    `${API_URL}/eventos?search=${encodeURIComponent(search)}`
  )

  if (!response.ok) {
    throw new Error("No se pudieron cargar los eventos")
  }

  return response.json()
}

export async function getRankingEventos() {
  const response = await fetch(`${API_URL}/eventos/ranking`)

  if (!response.ok) {
    throw new Error("No se pudo cargar el ranking")
  }

  return response.json()
}
export async function getEventoPorId(id) {
  const response = await fetch(
    `http://localhost:8080/eventos/${id}`
  )

  return await response.json()
}