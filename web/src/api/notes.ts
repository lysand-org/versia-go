import {HttpClient, HttpClientRequest, HttpClientResponse,} from "@effect/platform";
import * as Schema from "@effect/schema/Schema";
import {Effect} from "effect";
import {OTLP_TRACE_PROPAGATION} from "../env.ts";
import {APIResponse, BASE_URL, FailedAPIResponse} from "./response.ts";

export const NoteId = Schema.UUID.pipe(Schema.brand("UserId"));
export type NoteId = Schema.Schema.Type<typeof NoteId>;

export const NoteVisibility = Schema.Literal("public", "private", "direct");
export type NoteVisibility = Schema.Schema.Type<typeof NoteVisibility>;

export class Note extends Schema.Class<Note>("Note")({
    id: NoteId,
    visibility: NoteVisibility,
}) {
}

export class NoteCreation extends Schema.Class<NoteCreation>("NoteCreation")({
    content: Schema.String,
    visibility: Schema.optionalWith(NoteVisibility, {default: () => "public"}),
    mentions: Schema.optional(Schema.Array(Schema.String)),
}) {
}

const createNoteBody = HttpClientRequest.schemaBody(NoteCreation);
export const createNote = (data: NoteCreation) =>
    HttpClientRequest.post(`${BASE_URL}/api/app/notes/`).pipe(
        createNoteBody(data),
        Effect.andThen(HttpClient.fetch),
        HttpClientResponse.schemaBodyJsonScoped(APIResponse(Note)),
        Effect.filterOrFail(
            (res) => res.ok,
            (res) => (res as FailedAPIResponse).error,
        ),
        Effect.map((d) => d.data),
        Effect.either,
        Effect.withSpan("createNote", {attributes: {data: data}}),
        HttpClient.withTracerPropagation(OTLP_TRACE_PROPAGATION),
    );
