export function newToast(level: string, content: string) {
	const toastsBox: HTMLDivElement | null = document.querySelector(".toasts-box");
	if (toastsBox == null) {
		throw new Error("toasts box is missing");
	}

	const toastTemplate: HTMLTemplateElement | null = document.querySelector("#toast");
	if (toastTemplate == null) {
		throw new Error("toast template is missing");
	}

	const toast = toastTemplate.content.cloneNode(true) as DocumentFragment;

	const toastContent = toast.querySelector(".toast-content");
	if (toastContent == null) {
		throw new Error("missing content in toast template");
	}

	toast.firstElementChild?.setAttribute("data-toast", level);
	toastContent.innerHTML = content;

	toastsBox.appendChild(toast);
}
