import {useMutation, useQueryClient} from "@tanstack/react-query";
import {Effect, Either} from "effect";
import {createUser, UserCreation} from "../api/auth.ts";
import {runtime} from "../api";

type T = Effect.Effect.Success<ReturnType<typeof createUser>>;

export const useCreateUser = () => {
    const qc = useQueryClient();

    return useMutation<
        Either.Either.Right<T>,
        Either.Either.Left<T>,
        UserCreation
    >({
        mutationKey: ["user", "@me"],
        mutationFn: async (m) => {
            const v = await createUser(m).pipe(runtime.runPromise);
            if (Either.isLeft(v)) throw v.left;

            return v.right;
        },
        onSuccess: (data) => {
            qc.setQueryData(["user", "@me"], data);
        },
    });
};
