import { SignalWatcher } from './signal-watcher.js';
import { signal } from '@preact/signals-core';

import { LitElement, html } from "lit";
import { customElement, property } from "lit/decorators.js";

import './index.css';

const counter = signal(1);

const input1 = signal(1);
const input2 = signal(1);

class ExtElement extends SignalWatcher(LitElement) {
    createRenderRoot() {
        return this.shadowRoot;
    }
}

@customElement("my-element")
export class MyElement extends ExtElement {
    @property() who = "";
    @property() mul = 1;
    @property() counter = 0;

    render() {
        return html`
            <button type="button"
                class="m-2 inline-flex items-center rounded-md border
                border-transparent bg-gray-600 px-6 py-3 text-base font-medium
                text-white shadow-sm hover:bg-indigo-700 focus:outline-none
                focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                @click=${() => this.counter++}>
            ${this.who} ${this.counter * this.mul}
            </button>
        `;
    }
}

@customElement("my-input1")
export class MyInput1 extends ExtElement {
    render() {
        return html`
            <input
                class="p-2 m-4"
                @input=${this.resetCounter} placeholder="Enter a number">
        `;
    }

    resetCounter(event: Event) {
        const input = event.target as HTMLInputElement;
        input1.value = parseInt(input.value);
        if (isNaN(input1.value)) {
            input1.value = 0
        }
        counter.value = input1.value + input2.value
    }
}

@customElement("my-input2")
export class MyInput2 extends ExtElement {
    render() {
        return html`
            <input
                class="p-2 m-4"
                @input=${this.resetCounter} placeholder="Enter a number">
        `;
    }

    resetCounter(event: Event) {
        const input = event.target as HTMLInputElement;
        input2.value = parseInt(input.value);
        if (isNaN(input2.value)) {
            input2.value = 0
        }
        counter.value = input1.value + input2.value
    }
}