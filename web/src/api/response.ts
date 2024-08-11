import * as Schema from "@effect/schema/Schema";

export const BASE_URL = import.meta.env.VITE_BASE_URL ?? "INVALID";

export class APIError extends Schema.Class<APIError>("APIError")({
    status_code: Schema.Number,
    description: Schema.String,
    metadata: Schema.optional(
        Schema.Record({key: Schema.String, value: Schema.Any}),
    ),
}) {
    _tag = "APIError" as const;

    static toString(error: APIError) {
        return `${error.status_code}: ${error.description}${error.metadata ? JSON.stringify(error.metadata) : ""}`;
    }

    override toString() {
        return APIError.toString(this);
    }

    get message() {
        return this.toString();
    }
}

export const FailedAPIResponse = Schema.Struct({
    ok: Schema.Literal(false),
    error: APIError,
    data: Schema.Null,
});
export type FailedAPIResponse = Schema.Schema.Type<typeof FailedAPIResponse>;

export const APIResponse = <A, I>(data: Schema.Schema<A, I>) =>
    Schema.Union(
        Schema.Struct({
            ok: Schema.Literal(true),
            data: data,
            error: Schema.Null,
        }),
        FailedAPIResponse,
    );
