import { inject, Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class HttpService {
	private readonly http: HttpClient = inject(HttpClient);
	private readonly backendUrl = '/api';
	readonly wsUrl = 'http://localhost:5000/ws';

	private buildUrl(url: string) {
		return this.backendUrl.concat(url);
	}

	post<T, Y>(url: string, data: Y): Observable<T> {
		const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
		return this.http.post<T>(this.buildUrl(url), data, { headers });
	}
}
