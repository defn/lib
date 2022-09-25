import { html } from "lit";
import { customElement } from "lit/decorators.js";
import { ExtElement, counter, input1, input2 } from "./Common.js"

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
