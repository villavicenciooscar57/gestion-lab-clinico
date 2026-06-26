# Sistema de Gestión para Laboratorio Clínico

Este proyecto es una aplicación web integral desarrollada para la automatización de procesos de laboratorio clínico, facilitando el registro de pacientes, la captura de resultados analíticos y la gestión de historiales de examenes con capacidad de impresión de reportes.

## 👥 Datos/Autor
- **Estudiante:** Villavicencio Yubi Oscar Santiago
- **Institución:** Universidad Internacional del Ecuador (UIDE)
- **Materia:** Programación Orientada a Objetos 1
- **Fecha:** Junio 2026

## 🎯 Objetivo del Programa
Desarrollar y desplegar un sistema de gestión analítica de datos clínicos bajo una arquitectura cliente-servidor robusta, garantizando la integridad de los resultados médicos, la trazabilidad del historial del paciente y la generación de reportes físicos inmediatos mediante tecnologías web modernas.

## 🚀 Funcionalidades Principales
1. **Registro de Pacientes:** Ingreso de datos demográficos con generación de IDs inmutables (UUID) para asegurar la unicidad de los registros.
2. **Catálogo de Exámenes:** Gestión centralizada de pruebas clínicas disponibles en el laboratorio.
3. **Módulo de Reporte:** Interfaz dinámica para la carga de valores clínicos, con validación de datos en tiempo real.
4. **Historial Clínico:** Consulta rápida por cédula para visualizar la evolución del paciente.
5. **Reporte e Impresión:** Sistema de renderizado dedicado para la entrega física de resultados médicos.

## 🛠️ Tecnologías Utilizadas
- **Backend:** Go (Golang) - Uso de `net/http` para servicios RESTful.
- **Frontend:** HTML5, CSS3, JavaScript (Fetch API para consumo de servicios).
- **Base de Datos:** Supabase (PostgreSQL Cloud).
- **Control de Versiones:** Git / GitHub.
