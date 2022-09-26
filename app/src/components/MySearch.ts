import { html, ReactiveControllerHost } from "lit";
import { customElement, property } from "lit/decorators.js";
import { initialState, StatusRenderer, Task } from '@lit-labs/task';
import { ExtElement } from "./Common.js"

export type Result = Array<{ name: string }>;
export type Kind = typeof kinds[number];

export const baseUrl = 'https://swapi.dev/api/';

export const kinds = [
    '',
    'people',
    'starships',
    'species',
    'planets',
    // Inserted to demo an error state.
    'error'
] as const;

export class NamesController {
    host: ReactiveControllerHost;
    value?: string[];
    readonly kinds = kinds;
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
                    const response = await fetch(`${baseUrl}${kind}`);
                    const data = await response.json();
                    return data.results as Result;
                } catch {
                    throw new Error(`Failed to fetch "${kind}"`);
                }
            }, () => [this.kind]
        );
    }

    set kind(value: Kind) {
        this._kind = value;
        this.host.requestUpdate();
    }
    get kind() { return this._kind; }

    render(renderFunctions: StatusRenderer<Result>) {
        return this.task.render(renderFunctions);
    }
}

@customElement('my-search')
export class MySearch extends ExtElement {
    private names = new NamesController(this);

    render() {
        return html`
        <div class="m-8">

        <select @change=${this._selectKind}>
            ${this.names.kinds.map((k) => html`<option value=${k}>${k}</option>`)}
        </select>

        ${this.names.render({
            complete: (result: Result) => html`
                <p>List of ${this.names.kind}</p>
        
                <ul>
                    ${result.map(i => html`<li>${i.name}</li>`)}
                </ul>
            `,
            initial: () => html`<p>Select a kind...</p>`,
            pending: () => html`<p>Loading ${this.names.kind}...</p>`,
            error: (e: any) => html`<p>${e}</p>`
        })}

        </div>
        `;
    }

    private _selectKind(e: Event) {
        this.names.kind = (e.target as HTMLSelectElement).value as Kind;
    }
}
