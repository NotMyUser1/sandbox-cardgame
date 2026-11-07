import { Component, inject, signal } from '@angular/core';
import { HttpService } from '../shared_services/http.service';
import { ActivatedRoute } from '@angular/router';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
	imports: [ReactiveFormsModule],
	templateUrl: './gametable.html',
	styleUrl: './gametable.scss',
})
export class GameTable {
	message = signal<string>('');
	error = signal<string | null>(null);
	fb = inject(FormBuilder);
	form = this.fb.group({
		json: ['{"type":"play_card","index":1}', [Validators.required]],
	});
	gameId!: string;
	playerId!: string;
	private route = inject(ActivatedRoute);
	private http = inject(HttpService);
	private ws!: WebSocket;

	constructor() {
		// Get query parameters from the URL
		this.gameId = this.route.snapshot.queryParamMap.get('game_id') ?? '';
		this.playerId = this.route.snapshot.queryParamMap.get('player_id') ?? '';

		if (!this.gameId || !this.playerId) {
			this.error.set('Missing game_id or player_id');
			return;
		}

		this.openWebSocket();
	}

	sendMessage(): void {
		this.error.set('');
		const raw = this.form.value.json ?? '';
		try {
			// Validate JSON before sending
			const parsed = JSON.parse(raw);
			this.ws.send(JSON.stringify(parsed));
		} catch {
			this.error.set('Invalid JSON');
		}
	}

	ngOnDestroy(): void {
		if (this.ws && this.ws.readyState === WebSocket.OPEN) {
			this.ws.close();
		}
	}

	private openWebSocket(): void {
		const url = `${this.http.wsUrl}?game_id=${this.gameId}&player_id=${this.playerId}`;

		this.ws = new WebSocket(url);

		this.ws.onopen = () => {
			console.log('WebSocket connected');
		};

		this.ws.onmessage = (event: MessageEvent) => {
			// Assume each message is a JSON string; store it as formatted text
			try {
				const parsed = JSON.parse(event.data);
				const pretty = JSON.stringify(parsed, null, 2);
				this.message.set(pretty);
			} catch {
				// If not JSON, just store raw text
				this.message.set(event.data);
			}
		};

		this.ws.onerror = (ev) => {
			this.error.set('WebSocket error');
			console.error('WebSocket error', ev);
		};

		this.ws.onclose = () => {
			console.log('WebSocket closed');
		};
	}
}
