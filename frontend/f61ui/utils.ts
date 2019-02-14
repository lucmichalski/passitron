import { defaultErrorHandler } from 'backenderrors';
import { datetimeRFC3339 } from 'f61ui/types';
import * as moment from 'moment';

export function unrecognizedValue(value: never): never {
	throw new Error(`Unrecognized value: ${value}`);
}

let uniqueDomIdCounter = 0;

export function uniqueDomId(): number {
	return ++uniqueDomIdCounter;
}

export function relativeDateFormat(dateIso: datetimeRFC3339): string {
	return moment(dateIso).fromNow();
}

// - cannot "await" unless "await" sits in a function that itself is "async"
// - even if you catch all exceptions in "async" func, tslint no-floating-promises still
//   complains if your *caller* calls this function whose promise will never reject
// - therefore this hack was made to please tslint
export function shouldAlwaysSucceed(prom: Promise<any>): void {
	prom.then(() => {
		/* noop */
	}, defaultErrorHandler);
}

export function focusRetainer(logicThatMessesUpFocus: () => void) {
	const activeElementBefore = document.activeElement;

	logicThatMessesUpFocus();

	if (activeElementBefore instanceof HTMLElement) {
		activeElementBefore.focus();
	}
}
