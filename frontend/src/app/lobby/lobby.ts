import { Component, inject, signal, WritableSignal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { firstValueFrom } from 'rxjs';
import { HttpService } from '../shared_services/http.service';

@Component({
	imports: [CommonModule, ReactiveFormsModule],
	templateUrl: './lobby.html',
	styleUrl: './lobby.scss',
})
export class Lobby {
	loading = signal(false);
	error: WritableSignal<string | null> = signal(null);
	private http: HttpService = inject(HttpService);
	private router = inject(Router);
	private fb = inject(FormBuilder);
	form = this.fb.group({
		name: ['', [Validators.required]],
		gameId: ['', [Validators.pattern(/^[a-zA-Z0-9]*$/)]], // optional but must be alphanumeric if present
	});

	async onJoin() {
		if (this.form.invalid) {
			this.form.markAllAsTouched();
			return;
		}

		this.loading.set(true);
		this.error.set(null);

		const name = this.form.value.name!.trim();
		let gameId = (this.form.value.gameId || '').trim();

		try {
			if (!gameId) {
				// create new game
				const createResp = await firstValueFrom(
					this.http.post<{ game_id: string }, {}>('/create', {}),
				);
				gameId = createResp.game_id;
			}

			// join the game
			const joinResp = await firstValueFrom(
				this.http.post<{ player_id: string }, { game_id: string; name: string }>('/join', {
					game_id: gameId,
					name,
				}),
			);

			await this.router.navigate(['game'], {
				queryParams: { game_id: gameId, player_id: joinResp.player_id },
			});
		} catch (err: any) {
			this.error.set(err?.error || 'Failed to join/create game');
		} finally {
			this.loading.set(false);
		}
	}
}
