import {createLazyFileRoute} from "@tanstack/react-router";
import {useCreateNote} from "../hooks/useCreateNote.ts";
import {NoteCreation} from "../api/notes.ts";
import {useState} from "react";

export const Route = createLazyFileRoute("/post")({
    component: Post,
});

function Post() {
    const post = useCreateNote();
    const [content, setContent] = useState("");

    return (
        <div className="p-2 flex flex-col gap-2">
      <textarea
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="w-full rounded-md border-2 border-gray-300 p-2"
          placeholder="What's on your mind?"
          rows={10}
      />

            <button
                onClick={() =>
                    post.mutate(
                        new NoteCreation({
                            content: content,
                            visibility: "public",
                            mentions: [
                                "https://lysand.i.devminer.xyz/users/0190d697-c83a-7376-8d15-0f77fd09e180",
                            ],
                        }),
                    )
                }
                disabled={post.status === "pending"}
            >
                Post
            </button>
        </div>
    );
}
