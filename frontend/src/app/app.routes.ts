import { Routes } from '@angular/router';

export const routes: Routes = [
	{
		path: '',
		pathMatch: 'full',
		loadComponent: () => import('./lobby/lobby').then((m) => m.Lobby),
	},
	{
		path: 'game',
		loadComponent: () => import('./gametable/gametable').then((m) => m.GameTable),
	},
	{
		path: '**',
		redirectTo: '',
	},
];
