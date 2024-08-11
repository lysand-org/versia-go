import {useMutation, useQueryClient} from "@tanstack/react-query";
import {Effect, Either} from "effect";
import {runtime} from "../api";
import {createNote, Note, NoteCreation} from "../api/notes.ts";

type T = Effect.Effect.Success<ReturnType<typeof createNote>>;

export const useCreateNote = () => {
    const qc = useQueryClient();

    return useMutation<
        Either.Either.Right<T>,
        Either.Either.Left<T>,
        NoteCreation
    >({
        mutationKey: ["note"],
        mutationFn: async (m) => {
            const v = await createNote(m).pipe(runtime.runPromise);
            if (Either.isLeft(v)) throw v.left;

            return v.right;
        },
        onSuccess: (data) => {
            const notes = qc.getQueryData(["notes", "@me"]) as Note[] | undefined;
            if (notes) {
                qc.setQueryData(["notes", "@me"], {
                    ...notes,
                    data: [...notes, data],
                });
            }
        },
    });
};
