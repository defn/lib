import { html } from "lit";
import { customElement, property } from "lit/decorators.js";

import { ExtElement, counter } from "./Common.js";

@customElement("my-element")
export class MyElement extends ExtElement {
    @property() mul: number = 3;
    @property() who: string = "Earthly";

    render() {
        return html`
            <button type="button"
                class="m-2 inline-flex items-center rounded-md border
                border-transparent bg-gray-600 px-6 py-3 text-base font-medium
                text-white shadow-sm hover:bg-indigo-700 focus:outline-none
                focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                @click=${() => counter.value++}>
        ${this.who} ${this.mul} x ${counter.value} = ${counter.value * this.mul}
</button>
    `;
    }
}
