import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { FiArrowLeft, FiMail, FiRefreshCw, FiShield, FiUser, FiUsers } from "react-icons/fi";
import { getUsuarios, cambiarRol } from "../services/usuariosService";
import "../styles/AdminUsuarios.css";

export default function AdminUsuarios() {
  const [usuarios, setUsuarios] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  const handleCambiarRol = async (usuario) => {
    const nuevoRol = usuario.rol === "ADMIN" ? "CLIENTE" : "ADMIN";
    try {
      const token = localStorage.getItem("token");
      await cambiarRol(usuario.id, nuevoRol, token);
      setUsuarios(usuarios.map((u) => u.id === usuario.id ? { ...u, rol: nuevoRol } : u));
      alert("Rol actualizado correctamente");
    } catch (error) {
      console.error(error);
      alert(error.message);
    }
  };

  useEffect(() => {
    async function cargarUsuarios() {
      try {
        const token = localStorage.getItem("token");
        const data = await getUsuarios(token);
        setUsuarios(data);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    }
    cargarUsuarios();
  }, []);

  return (
    <main className="admin-users-page">
      <div className="admin-users-container">
        <header className="admin-users-header">
          <div>
            <span className="users-eyebrow">Panel administrativo</span>
            <h1>Administración de Usuarios</h1>
            <p>Gestioná roles y permisos de los usuarios registrados</p>
          </div>
          <span className="users-count"><FiUsers aria-hidden="true" /> {usuarios.length} usuarios</span>
        </header>
        <nav className="users-actions" aria-label="Acciones de administración">
          <button className="users-back-button" onClick={() => navigate("/admin/eventos")}><FiArrowLeft aria-hidden="true" /> Volver a Eventos</button>
        </nav>

        {loading ? (
          <div className="users-feedback" role="status"><span className="users-spinner" />Cargando usuarios...</div>
        ) : usuarios.length === 0 ? (
          <div className="users-feedback users-empty"><FiUsers aria-hidden="true" /><h2>No hay usuarios registrados</h2></div>
        ) : (
          <section className="users-grid" aria-label="Usuarios registrados">
            {usuarios.map((usuario) => (
              <article className="user-card" key={usuario.id}>
                <div className="user-card-heading">
                  <div className="user-avatar" aria-hidden="true"><FiUser /></div>
                  <div><h2>{usuario.nombre}</h2><a href={`mailto:${usuario.email}`}><FiMail aria-hidden="true" />{usuario.email}</a></div>
                </div>
                <div className="user-role-row">
                  <span className="role-label">Rol actual</span>
                  <span className={`role-badge role-${usuario.rol?.toLowerCase()}`}><FiShield aria-hidden="true" />{usuario.rol}</span>
                </div>
                <button className="change-role-button" onClick={() => handleCambiarRol(usuario)}><FiRefreshCw aria-hidden="true" /> Cambiar Rol</button>
              </article>
            ))}
          </section>
        )}
      </div>
    </main>
  );
}
