import { SignalWatcher } from './signal-watcher.js';
import { signal } from '@preact/signals-core';

import { LitElement, html, css, unsafeCSS } from "lit";

import componentStyles from './index.css'

const counter = signal(1);

const input1 = signal(1);
const input2 = signal(1);

class ExtElement extends SignalWatcher(LitElement) {
    static styles = [
        css`${unsafeCSS(componentStyles)}`
    ]
}

export class MyElement extends ExtElement {
    who: string;
    mul: number;

    static properties = {
        who: { type: String },
        mul: { type: Number },
    }

    constructor() {
        super();
        this.who = ""
        this.mul = 1;
    }

    render() {
        return html`
            <button type="button"
                class="m-2 inline-flex items-center rounded-md border
                border-transparent bg-gray-600 px-6 py-3 text-base font-medium
                text-white shadow-sm hover:bg-indigo-700 focus:outline-none
                focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                @click=${() => counter.value++}>
            ${this.who} ${counter.value * this.mul}
            </button>
        `;
    }
}

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

customElements.define('my-element', MyElement);
customElements.define('my-input1', MyInput1);
customElements.define('my-input2', MyInput2);
