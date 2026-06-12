# Fallstudie - Sandbox Cardgame
## Titel
- Sandbox Cardgame - Spiele nach deinen Regeln

## Einleitung

## Problemdefinition
- Die meisten Seiten lassen nicht alle Regeln, oder Spielmöglichkeiten zu, deswegen können Hausregeln o.ä. nicht online gespielt werden
- Außerdem gibt es einige Spiele, die online nicht verfügbar sind
	- Beispiele (?)

## Zielsetzung
- Ein generelles Spiel, womit sich die meisten Kartenspiele abbilden lassen

## Lösungsvorschlag
- Spielkonzept:
	- Beliebig viele Spieler
	- Großes und kleines Kartenset
	- Jeder kann zu jederzeit jede Karte spielen
	- Mit dem Stapel kann uneingeschränkt interagiert werden
	
- Architektur:
	- Client-Server
	- Server hält den Spielstand
	- Kommunikation läuft parallel über Websockets
	- SPA im Frontend
	- Skalierbar, durch Spielids, die dann mit Containern im Backend unterschiedlichen Servern zuordnen und nach Bedürnfnis starten(horizontale Skalierung) Sharding(?))
	
- Technologien:
	- Go im Backend
		- Gewählt wegen der hohen Performance
		- Auch als Lernprojekt
	- Angular im Frontend
		- Praktisch zur Entwicklung von SPAs
		- Vorkenntnisse aus dem Unternehmen
	- Websockets für die Kommunikation, wie und warum
	- Docker für Entwicklung und Deployment
	- Codebeispiele, um die Struktur zu zeigen

## Konzepte
- Regelwerk-Konfiguration
- Mobile Interface, oder nur PC?
- Regelverstoß-Mahnung per Majority-Voting, oder als GM Entscheidung
- Handling vom Gamestate an unterschiedlichen Orten(RAM, Datenbank, verteilter Speicher) Vor-Nachteile
- *Custom Kartendecks hochladbar, für Spiele die nicht mit Standarddeck spielbar sind
		
## Analyse aus verschiedenen Perspektiven
- User Stories:
- Vielleicht später

## Chancen und Risiken(?)
Chancen:
- Zentrales Hub für alle
- Menschen zusammenbringen
- Potenziell unendliche Erweiterbarkeit
- Lerneffekt für uns(?)

Risiken:
- Serverlimitierung(CPU/RAM)
- Finanzielles Risiko, da es noch kein Revenue-Modell gibt

## Fazit
[ ] Wenn der Rest fertig ist
