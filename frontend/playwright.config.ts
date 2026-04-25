import { defineConfig, devices } from '@playwright/test';

// noinspection JSUnusedGlobalSymbols
export default defineConfig({
	testDir: './e2e',
	fullyParallel: true,
	timeout: 30_000,
	use: {
		baseURL: 'http://127.0.0.1:4200',
		trace: 'on-first-retry',
	},
	webServer: {
		command: 'npm run start -- --host 127.0.0.1 --port 4200',
		url: 'http://127.0.0.1:4200',
		reuseExistingServer: !process.env.CI,
		timeout: 120_000,
	},
	projects: [
		{
			name: 'chromium',
			use: { ...devices['Desktop Chrome'] },
		},
	],
});
