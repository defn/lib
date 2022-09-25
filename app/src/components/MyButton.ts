import { html } from "lit";
import { customElement, property } from "lit/decorators.js";
import { ExtElement, counter, input1, input2 } from "./Common.js"

@customElement("my-button")
export class MyButton extends ExtElement {
    @property() who = ""
    @property() mul = 1

    static buttonClasses = `
        m-8 inline-flex items-center rounded-md border border-gray-300 bg-white
        px-4 py-2 font-medium text-gray-700 shadow-sm hover:bg-gray-50
        focus:outline-none focus:ring-2 focus:ring-indigo-500
        focus:ring-offset-2 sm:text-sm`

    render() {
        return html`
        <button type="button" class=${MyButton.buttonClasses}
            @click=${() => counter.value += 1}>
            ${this.who}: ${counter.value} * ${this.mul} = ${counter.value * this.mul}
            </button>
        `;
    }
}
