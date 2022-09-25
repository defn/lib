import { LitElement, html } from "lit";
import { SignalWatcher } from './signal-watcher.js';
import { signal } from '@preact/signals-core';

import './index.css'

export const counter = signal(1);
export const input1 = signal(1);
export const input2 = signal(1);

export class ExtElement extends SignalWatcher(LitElement) {
    xcreateRenderRoot() {
        return this.shadowRoot;
    }
}
