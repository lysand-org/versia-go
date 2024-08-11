import {HttpClient, HttpClientRequest, HttpClientResponse,} from "@effect/platform";
import * as Schema from "@effect/schema/Schema";
import {Effect, Schedule} from "effect";
import {OTLP_TRACE_PROPAGATION} from "../env.ts";
import {APIResponse, BASE_URL, FailedAPIResponse} from "./response.ts";

export const UserId = Schema.UUID.pipe(Schema.brand("UserId"));
export type UserId = Schema.Schema.Type<typeof UserId>;

export class UserInfo extends Schema.Class<UserInfo>("UserInfo")({
    id: UserId,
    username: Schema.String,
}) {
}

export const fetchUserInfo = HttpClientRequest.get(
    `${BASE_URL}/api/auth/authn/@me`,
).pipe(
    HttpClient.fetch,
    HttpClientResponse.schemaBodyJsonScoped(APIResponse(UserInfo)),
    Effect.filterOrFail(
        (res) => res.ok,
        (res) => (res as FailedAPIResponse).error,
    ),
    Effect.retry({times: 3, schedule: Schedule.exponential(250, 2)}),
    Effect.map((d) => d.data),
    Effect.either,
    Effect.withSpan("fetchUserInfo"),
    HttpClient.withTracerPropagation(OTLP_TRACE_PROPAGATION),
);

export class UserCreation extends Schema.Class<UserCreation>("UserCreation")({
    username: Schema.String,
}) {
}

const createUserBody = HttpClientRequest.schemaBody(UserCreation);
export const createUser = (data: UserCreation) =>
    HttpClientRequest.post(`${BASE_URL}/api/app/users/`).pipe(
        createUserBody(data),
        Effect.andThen(HttpClient.fetch),
        HttpClientResponse.schemaBodyJsonScoped(APIResponse(UserInfo)),
        Effect.filterOrFail(
            (res) => res.ok,
            (res) => (res as FailedAPIResponse).error,
        ),
        Effect.map((d) => d.data),
        Effect.either,
        Effect.withSpan("createUser", {attributes: {data: data}}),
        HttpClient.withTracerPropagation(OTLP_TRACE_PROPAGATION),
    );
