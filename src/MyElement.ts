import { html } from "lit";
import { customElement, property } from "lit/decorators.js";

import { ExtElement, counter } from "./Common.js"

@customElement("my-element")
export class MyElement extends ExtElement {
    @property() who = "Earthly";
    @property() mul = 1;

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
