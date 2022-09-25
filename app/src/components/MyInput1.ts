import { customElement } from "lit/decorators.js";
import { ExtElement, counter, input1, input2 } from "./Common.js"

@customElement("my-input1")
export class MyInput1 extends ExtElement {
    static inputClasses = `
        p-2 m-8`

    render() {
        return html`
        <input
            class=${MyInput1.inputClasses}
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
