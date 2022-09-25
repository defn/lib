import { LitElement, css, unsafeCSS, CSSResultGroup } from "lit";
import { SignalWatcher } from './signal-watcher.js';
import { signal } from '@preact/signals-core';

// @ts-ignore
import componentStyle from './index.css'

export const counter = signal(1);
export const input1 = signal(1);
export const input2 = signal(1);

export class ExtElement extends SignalWatcher(LitElement) {
    static styles = css`${unsafeCSS(componentStyle)}`
    xcreateRenderRoot() {
        return this.shadowRoot;
    }
}
