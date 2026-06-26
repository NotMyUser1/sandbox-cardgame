# Fallstudie - Sandbox Cardgame
* Ist für Einträge aus dem Kapitel Ausblicke
## Titel
- Sandbox Cardgame - Spiele nach deinen Regeln

## Einleitung

## Problemdefinition
- Die meisten Seiten lassen nicht alle Regeln, oder Spielmöglichkeiten zu, deswegen können Hausregeln o.ä. nicht online gespielt werden
- Außerdem gibt es einige Spiele, die online nicht verfügbar sind
	- Beispiele (?)

## Zielsetzung
Leitfrage:
- Wie lässt sich ein Online-Kartenspiel designen, womit man die meisten Hausregeln abbilden kann

## Lösungsvorschlag
- Spielkonzept:
	- Beliebig viele Spieler
	- Großes und kleines Kartenset
	- Jeder kann zu jederzeit jede Karte spielen
	- Mit dem Stapel kann uneingeschränkt interagiert werden
	- *Regelwerk-Konfiguration
	- *Regelverstoß-Mahnung per Majority-Voting, oder als GM Entscheidung
	- *Custom Kartendecks hochladbar, für Spiele die nicht mit Standarddeck spielbar sind

	Prozess(Zur Erklärung der Konzepte):
	- Wie wir zu den Regeln gekommen sind?
		- Welche Kartenspiele wollen wir abbilden
		- Flussdiagramm wie der Spielablauf wäre
	
- Architektur:
	- Welches klassische Architektur-Prinzip wird genutzt?
	- Client-Server
	- Server hält den Spielstand
	- Kommunikation läuft parallel über Websockets
	- SPA im Frontend
	- Skalierbar, durch Spielids, die dann mit Containern im Backend unterschiedlichen Servern zuordnen und nach Bedürnfnis starten(horizontale Skalierung) Sharding(?))
	- *Handling vom Gamestate an unterschiedlichen Orten(RAM, Datenbank, verteilter Speicher) Vor-Nachteile
	
- Technologien:
	- Welche Technologien standen zur Auswahl?
	- Go im Backend
		- Gewählt wegen der hohen Performance
		- Auch als Lernprojekt
	- Angular im Frontend
		- Praktisch zur Entwicklung von SPAs
		- Vorkenntnisse aus dem Unternehmen
	- Websockets für die Kommunikation, wie und warum
	- Docker für Entwicklung und Deployment
	- Codebeispiele, um die Struktur zu zeigen

## Prototyp
- UML-Diagramm
- UI Mockup (wie?)
- *Mobile Interface, oder nur PC?

		
## Stakeholder-Analyse
- 1 User Story
	- Enthusiastischer Kartenspieler
	- Interessantes Profil vom Menschen
	- Menschen möchten dass nicht geschummelt wird

## Chancen und Risiken(?)
Chancen:
- Zentrales Hub für alle
- Menschen zusammenbringen
- Potenziell unendliche Erweiterbarkeit
- Lerneffekt für uns(?)

Risiken:
- Serverlimitierung(CPU/RAM)
- Finanzielles Risiko, da es noch kein Revenue-Modell gibt
- Für Unwissende eher ungeeignet


## Fazit
[ ] Wenn der Rest fertig ist

## Folienaufteilung
Jannes:
- Lösungsansatz
Yannik:
- Thema, Problemdefinition
- User Story

Bis zur nächster Woche etwas haben
