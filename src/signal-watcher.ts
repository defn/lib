import type { ReactiveElement } from 'lit';
import { signal, effect } from '@preact/signals-core';

type ReactiveElementConstructor = new (...args: any[]) => ReactiveElement;

export function SignalWatcher<T extends ReactiveElementConstructor>(Base: T): T {
    return class SignalWatcher extends Base {
        private _disposeEffect?: () => void;

        performUpdate() {
            if (!this.isUpdatePending) {
                return;
            }
            this._disposeEffect?.();
            this._disposeEffect = effect(() => {
                this.isUpdatePending = true;
                super.performUpdate();
            });
        }
    };
}
