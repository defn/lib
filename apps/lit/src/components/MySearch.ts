import { html, ReactiveControllerHost } from "lit";
import { customElement, property } from "lit/decorators.js";
import { initialState, StatusRenderer, Task } from '@lit-labs/task';
import { ExtElement } from "./Common.js"

import { meh } from '@defn/utils'

export type Result = Array<{ name: string }>;
export type Kind = typeof kinds[number];

export const kinds = [
    '',
    'meh',
] as const;

export class NamesController {
    host: ReactiveControllerHost;
    value?: string[];

    private task!: Task;
    private _kind: Kind = '';

    constructor(host: ReactiveControllerHost) {
        this.host = host;

        this.task = new Task<[Kind], Result>(host,
            async ([kind]: [Kind]) => {
                if (!kind?.trim()) {
                    return initialState;
                }

                try {
                    const response = await fetch(`https://control-0.tiger-mamba.ts.net/api/${kind}`);
                    const data = await response.json();
                    return data.results as Result;
                } catch (err) {
                    throw new Error(`Failed to fetch ${kind}: ${err.message}`);
                }
            }, () => [this.kind]
        );
    }

    get kind() { return this._kind; }

    set kind(value: Kind) {
        this._kind = value;
        this.host.requestUpdate();
    }

    render(renderFunctions: StatusRenderer<Result>) {
        return this.task.render(renderFunctions);
    }
}

@customElement('my-search')
export class MySearch extends ExtElement {
    private names = new NamesController(this);

    static selectClasses = `
        inline-flex items-center rounded-md border border-gray-300 bg-white px-4 pr-8
        py-2 font-medium text-gray-700 shadow-sm hover:bg-gray-50
        focus:outline-none focus:ring-2 focus:ring-indigo-500
        focus:ring-offset-2 sm:text-sm`

    render() {
        return html`
        <div class="m-8">

        <select class=${MySearch.selectClasses}
            @change=${this._selectKind}>
            ${kinds.map((k) => html`<option value=${k}>${k}</option>`)}
        </select>

        <div class="m-8">
        ${this.names.render({
            initial: () => html`<p>Make a selection ${meh}</p>`,

            pending: () => html`<p>Fetching ${this.names.kind}...</p>`,

            error: (e: any) => html`<p>${e}</p>`,

            complete: (result: Result) => html`
                <ul>
                    ${result.map(i => html`<li>${i.name}</li>`)}
                </ul>
                `
        })}
        </div>

        </div>
        `;
    }

    private _selectKind(e: Event) {
        this.names.kind = (e.target as HTMLSelectElement).value as Kind;
    }
}
