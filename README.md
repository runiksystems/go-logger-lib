# Go Logger Library
*RunikSystems - 2025*

> Un module Go professionnel et structuré pour la journalisation (logging) en utilisant le package standard log/slog. Conçu pour le Cloud Native (format JSON) et facilement configurable.

## Fonctionnalités

Format structuré : Utilise JSON par défaut pour l'intégration avec les systèmes de logs (Elasticsearch, Loki, etc.).

Contexte de Service : Ajoute un champ service permanent à toutes les entrées de log.

Niveau de Log Dynamique : Supporte Debug, Info, Warn, et Error.

## Utilisation

### 1. Installation

Ajoutez le module à votre projet Go (remplacez VotreUtilisateur par votre nom d'utilisateur GitHub) :

```bash
go get github.com/runiksystems/go-logger-lib@v1.0.0
```


### 2. Initialisation et Configuration

Appelez logger.Init() au début de votre fonction main :

```go
package main

import (
	"log/slog"
	"github.com/runiksystems/go-logger-lib"
)

func main() {
	// Configuration personnalisée : DEBUG activé, format JSON, nom du service
	cfg := logger.Config{
		EnableJSON:  true,
		LogLevel:    slog.LevelDebug,
		ServiceName: "mon-api-gateway",
		AddSource:   false,
	}

	logger.Init(cfg)
	
	// Utilisation simple via le logger par défaut (slog.Default())
	slog.Info("Démarrage de l'application", "port", 8080)
	slog.Debug("Tentative de connexion à la base de données...")
}
```

## Licence

Ce projet est sous licence GNU Affero General Public License Version 3 (AGPLv3).

Cette licence est une forme de Copyleft Fort (**Anti-SaaS**).

Libertés Garanties : L'usage commercial et la modification sont autorisés.

**Obligation : Si vous modifiez ce code et l'utilisez pour offrir un service à distance via un réseau (SaaS, API, etc.), vous êtes dans l'obligation de mettre à disposition le code source complet et modifié à tous les utilisateurs du service.**

Pour plus de détails, consultez le fichier LICENSE à la racine du dépôt, ou la page officielle de la licence :
https://www.gnu.org/licenses/agpl-3.0.html