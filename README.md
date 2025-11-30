# FindMeAJob

**FindMeAJob** est une plateforme intelligente d'automatisation de recherche d'emploi. Plus qu'un simple agrégateur, elle agit comme un assistant personnel de carrière.

## 🎯 Vision du projet

L'objectif est de révolutionner la candidature en automatisant le matching et la personnalisation des documents :

1.  **Smart Matching** : Analyse de votre profil (expériences, envies, paramètres) pour identifier les offres qui vous correspondent vraiment.
2.  **Génération Automatique** : Création dynamique de **CV** et de **Lettres de Motivation** sur-mesure pour chaque offre.
3.  **Pertinence Contextuelle** : L'IA met en lien vos tâches passées spécifiques avec les pré-requis du poste visé pour maximiser vos chances.

## 🚀 État du projet

Le projet est actuellement en phase de **développement**.

### Stack Technique
- **Frontend** : [Next.js](https://nextjs.org/) (React)
- **Backend** : [NestJS](https://nestjs.com/) (Node.js)
- **Scraper** : [Go](https://go.dev/) (Golang)
- **Base de données** : PostgreSQL
- **Cache/Queue** : Redis

## 🛠️ Installation et Lancement

Le projet est entièrement conteneurisé pour faciliter le développement et le déploiement.

### Pré-requis
- [Docker](https://www.docker.com/) et Docker Desktop installés sur votre machine.

### Démarrage rapide

Pour lancer l'ensemble de l'application (Frontend, Backend, Scraper, Base de données), exécutez simplement la commande suivante à la racine du projet :

```bash
docker-compose up --build
```

Cette commande va construire les images et démarrer tous les services nécessaires.

## 🌐 Accès aux services

Une fois les conteneurs démarrés, vous pouvez accéder aux différentes parties de l'application via votre navigateur :

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend** | [http://localhost:3000](http://localhost:3000) | Interface utilisateur principale |
| **Backend API** | [http://localhost:3001](http://localhost:3001) | API REST (Documentation Swagger si disponible) |
| **Scraper** | *Background Process* | Accessible via les logs Docker. Récupère les offres périodiquement. |

## 📂 Structure du projet

- `backend/` : Code source de l'API NestJS.
- `frontend/` : Code source de l'application Next.js.
- `scraper/` : Scripts Go pour le scraping des données.
