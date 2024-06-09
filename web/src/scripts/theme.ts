export const ThemeDark = "dark";
export const ThemeLight = "light";

export function setTheme(theme: string) {
	const root = document.querySelector(":root");
	if (root == null) {
		return;
	}

	localStorage.setItem("theme", theme);
	root.setAttribute("data-theme", theme);
}

function loadTheme() {
	const current = localStorage.getItem("theme");
	if (current != null) {
		setTheme(current);
		return;
	}

	if (window.matchMedia("(prefers-color-scheme: light)").matches) {
		setTheme(ThemeLight);
		return;
	}

	setTheme(ThemeDark);
}

loadTheme();
