import './index.css'
import { LitElement, html, css } from "lit";
import { customElement, property } from "lit/decorators.js";
import { SignalWatcher } from './signal-watcher.js';
import { signal } from '@preact/signals-core';

const counter = signal(1);

class ExtElement extends LitElement {
    createRenderRoot() {
        return this;
    }
}

@customElement("my-element")
export class MyElement extends SignalWatcher(ExtElement) {
    @property() name = "Earthly";

    render() {
        return html`
            <button type="button"
                class="m-2 inline-flex items-center rounded-md border
                border-transparent bg-gray-600 px-6 py-3 text-base font-medium
                text-white shadow-sm hover:bg-indigo-700 focus:outline-none
                focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                @click=${() => counter.value++}>
            ${this.name} ${counter}
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
            </svg>
            </button>
        `;
    }
}

@customElement("my-input")
export class MyInput extends SignalWatcher(ExtElement) {
    render() {
        return html`
            <input
                class="p-2 m-4"
                @input=${this.resetCounter} placeholder="Enter a number">
        `;
    }

    resetCounter(event: Event) {
        const input = event.target as HTMLInputElement;
        counter.value = parseInt(input.value);
    }
}
