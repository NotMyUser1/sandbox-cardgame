import { expect, test } from '@playwright/test';

test('textarea clears on submit', async ({ page }) => {
	await page.goto('/');

	const textarea = page.locator('#input-text');
	await textarea.fill('Hello from Playwright');
	await page.getByRole('button', { name: 'Submit' }).click();

	await expect(textarea).toHaveValue('');
});
