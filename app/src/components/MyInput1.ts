import { html } from "lit";
import { customElement } from "lit/decorators.js";
import { ExtElement, counter, input1, input2 } from "./Common.js"

@customElement("my-input1")
export class MyInput1 extends ExtElement {
    render() {
        return html`
            <input
                class="p-2 m-8"
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
