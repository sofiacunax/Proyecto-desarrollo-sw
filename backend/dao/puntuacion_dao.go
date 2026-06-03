package dao

import (
	"proyecto-desarrollo-sw/backend/db"
	"proyecto-desarrollo-sw/backend/dtos"
	"proyecto-desarrollo-sw/backend/models"
)

type PuntuacionDAO struct{}

func NewPuntuacionDAO() *PuntuacionDAO {
	return &PuntuacionDAO{}
}

func (dao *PuntuacionDAO) CrearPuntuacion(puntuacion models.Puntuacion) error {
	query := `
		INSERT INTO puntuaciones
		(usuario_id, evento_id, puntuacion)
		VALUES (?, ?, ?)
	`

	_, err := db.DB.Exec(
		query,
		puntuacion.UsuarioID,
		puntuacion.EventoID,
		puntuacion.Puntuacion,
	)

	return err
}

func (dao *PuntuacionDAO) ObtenerPuntuacionesPorEvento(eventoID int) ([]models.Puntuacion, error) {
	query := `
		SELECT id, usuario_id, evento_id, puntuacion
		FROM puntuaciones
		WHERE evento_id = ?
	`

	rows, err := db.DB.Query(query, eventoID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var puntuaciones []models.Puntuacion

	for rows.Next() {
		var puntuacion models.Puntuacion

		err := rows.Scan(
			&puntuacion.ID,
			&puntuacion.UsuarioID,
			&puntuacion.EventoID,
			&puntuacion.Puntuacion,
		)

		if err != nil {
			return nil, err
		}

		puntuaciones = append(puntuaciones, puntuacion)
	}

	return puntuaciones, nil
}

func (dao *PuntuacionDAO) ObtenerPromedioPorEvento(eventoID int) (float64, error) {
	query := `
		SELECT COALESCE(AVG(puntuacion), 0)
		FROM puntuaciones
		WHERE evento_id = ?
	`

	var promedio float64

	err := db.DB.QueryRow(query, eventoID).Scan(&promedio)
	if err != nil {
		return 0, err
	}

	return promedio, nil
}

func (dao *PuntuacionDAO) ObtenerRankingEventos() ([]dtos.RankingDTO, error) {
	query := `
		SELECT 
			e.id,
			e.titulo,
			COALESCE(AVG(p.puntuacion), 0) AS promedio
		FROM eventos e
		LEFT JOIN puntuaciones p ON e.id = p.evento_id
		GROUP BY e.id, e.titulo
		ORDER BY promedio DESC
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ranking []dtos.RankingDTO

	for rows.Next() {
		var item dtos.RankingDTO

		err := rows.Scan(
			&item.EventoID,
			&item.Titulo,
			&item.Promedio,
		)

		if err != nil {
			return nil, err
		}

		ranking = append(ranking, item)
	}

	return ranking, nil
}
