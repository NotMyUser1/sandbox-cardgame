import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Gametable } from './gametable';

describe('Gametable', () => {
	let component: Gametable;
	let fixture: ComponentFixture<Gametable>;

	beforeEach(async () => {
		await TestBed.configureTestingModule({
			imports: [Gametable],
		}).compileComponents();

		fixture = TestBed.createComponent(Gametable);
		component = fixture.componentInstance;
		await fixture.whenStable();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});
